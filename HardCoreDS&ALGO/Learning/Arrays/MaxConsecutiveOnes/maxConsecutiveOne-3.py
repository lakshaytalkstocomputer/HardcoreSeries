from typing import List
import time


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        """
        :type nums: List[int]
        "rtype    : int
        """
        num_one_list = []
        one_found = 0
        for num in nums:
            if num:
                one_found = one_found + 1
            else:
                num_one_list.append(one_found)
                one_found = 0

        num_one_list.append(one_found)
        return max(num_one_list)


if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.findMaxConsecutiveOnes([1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end-start))

'''
    The Above approach is little fast but the performance decreases with every zero  
'''