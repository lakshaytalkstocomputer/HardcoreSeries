from typing import List
import time

class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        """
        :type nums: List[int]
        "rtype    : int
        """
        max_cons_one = 0
        cons_one = 0

        for num in nums:
            if num:
                cons_one = cons_one + 1
            else:
                if cons_one > max_cons_one:
                    max_cons_one = cons_one
                cons_one = 0
        else:
            if cons_one > max_cons_one:
                max_cons_one = cons_one

        return max_cons_one


if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.findMaxConsecutiveOnes([1, 0, 1, 1, 0, 1, 1, 1, 1])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end-start))

'''
    The Above approach is faster than maxConsecutiveOne-2 as 
         we are only checking when cons_one is greater than max_cons_one
         only when we get any 0 in array and always when we finish
     
'''