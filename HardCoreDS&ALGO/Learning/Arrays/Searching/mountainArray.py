import time
class Solution(object):
    def validMountainArray(self, arr):
        """
        :type arr: List[int]
        :rtype: bool
        """
        if len(arr) < 3:
            return False

        max_index = 0
        max_num = arr[0]

        for index, num in enumerate(arr):
            if num > max_num:
                max_index = index
                max_num = num

        is_strictly_inc = True
        for i in range(max_index):
            if arr[i] >= arr[i + 1]:
                is_strictly_inc = False

        is_strictly_dec = True
        for i in range(max_index, len(arr) - 1):
            if arr[i] <= arr[i + 1]:
                is_strictly_dec = False

        return is_strictly_inc and is_strictly_dec and max_index != len(arr) - 1 and max_index != 0

if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    nums = [-2, 0, 10, -19, 4, 6, -8]
    output = sol.checkIfExist(nums)

    # end time
    end = time.time()

    print("Output     : {} {} \nTime Taken : {}".format(output, nums, end - start))
