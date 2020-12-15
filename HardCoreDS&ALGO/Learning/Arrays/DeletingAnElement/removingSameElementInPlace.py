import time


class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        # 0 1 2 2 3 0 4 2
        # -> 0 1 - - 3 - 4 -   -> count   | totalLen = 8
        # ->
        totalLen, lastIndex = len(nums), len(nums) - 1
        valCount = 0

        for index, num in enumerate(nums):
            while nums[lastIndex] == val and index < lastIndex:
                lastIndex -= 1
                valCount += 1

            if nums[index] == val and index <= lastIndex:
                nums[index], nums[lastIndex] = nums[lastIndex], val
                lastIndex -= 1
                valCount += 1

        return totalLen - valCount


if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    nums = [3,3]
    output = sol.removeElement(nums, 3)

    # end time
    end = time.time()

    print("Output     : {} {} \nTime Taken : {}".format(output, nums, end - start))
