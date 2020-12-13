import time

class Solution(object):
    def findNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        num_of_even_digit = 0
        for num in nums:
            num_len = 0
            while num > 0:
                num_len += 1
                num = num // 10

            if num_len % 2 == 0:
                num_of_even_digit += 1

        return num_of_even_digit

if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.findNumbers([12,345,2,6,7896])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end-start))