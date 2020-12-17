import time


class Solution(object):
    def checkIfExist(self, nums):
        """
        :type nums: List[int]
        :rtype: Boolean
        """
        for num in nums:
            for item in nums:
                if num * 2 == item:
                    return True

if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    nums = [-2, 0, 10, -19, 4, 6, -8]
    output = sol.checkIfExist(nums)

    # end time
    end = time.time()

    print("Output     : {} {} \nTime Taken : {}".format(output, nums, end - start))
