#         # Two-pointer
#         answer = [0] * len(A)
#         l, r = 0, len(A) - 1
#         while l <= r:
#             left, right = abs(A[l]), abs(A[r])
#             if left > right:
#                 answer[r - l] = left * left
#                 l += 1
#             else:
#                 answer[r - l] = right * right
#                 r -= 1
#         return answer

import time

class Solution(object):
    def sortedSquares(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        answer = [0] * len(A)
        l, r = 0, len(A) - 1
        while l <= r:
            left, right = abs(A[l]), abs(A[r])
            if left > right:
                answer[r - l] = left * left
                l += 1
            else:
                answer[r - l] = right * right
                r -= 1
        return answer



if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    output = sol.sortedSquares([-7,-3,2,3,11])

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end - start))