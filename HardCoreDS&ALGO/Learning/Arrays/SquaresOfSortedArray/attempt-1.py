import time

class Solution(object):
    def sortedSquares(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        squared_nums = [num**2 for num in nums]
        return sorted(squared_nums)



if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.sortedSquares([-7,-3,2,3,11])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end - start))