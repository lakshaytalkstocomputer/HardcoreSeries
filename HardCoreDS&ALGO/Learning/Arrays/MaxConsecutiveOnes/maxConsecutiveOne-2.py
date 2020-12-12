from typing import List
import time


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        """
        :type nums: List[int]
        "rtype    : int
        """
        max_cons_one_found = 0
        cons_one = 0

        for num in nums:
            if num:
                cons_one = cons_one + 1
            else:
                cons_one = 0

            if cons_one > max_cons_one_found:
                max_cons_one_found = cons_one

        return max_cons_one_found


if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.findMaxConsecutiveOnes([1, 0, 1, 1, 0, 1, 1, 1, 1])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end-start))

'''
    The Above approach is slow as for every number 
      we are checking the whether the cons_one is greater then max_con_one_found
'''
