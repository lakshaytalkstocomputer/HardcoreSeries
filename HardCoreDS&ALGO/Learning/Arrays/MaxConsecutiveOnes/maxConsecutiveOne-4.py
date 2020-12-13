from typing import List
import time


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        """
        :type nums: List[int]
        "rtype    : int
        """
        cons_one_list = [0] * len(nums)
        cons_one_list_index = 0

        for num in nums:
            if num:
                cons_one_list[cons_one_list_index] += 1
            else:
                cons_one_list_index += 1

        return max(cons_one_list)

if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.findMaxConsecutiveOnes([1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end-start))
