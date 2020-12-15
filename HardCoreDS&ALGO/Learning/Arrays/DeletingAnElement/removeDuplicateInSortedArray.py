import time


class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        # total_len = len(nums)
        # final_len = 0
        #
        # for index, num in enumerate(nums):
        #     next_index = index + 1
        #     while next_index < total_len and nums[next_index] <= num:
        #         next_index += 1
        #
        #     if next_index > (total_len - 1):
        #         final_len = index
        #         break
        #
        #     if next_index - index > 1:
        #         nums[index+1], nums[next_index] = nums[next_index], nums[index+1]
        #
        # return final_len + 1
        if len(nums) == 0:
            return 0
        n = len(nums)
        curr = 0
        curr_val = nums[0]
        i = 1
        for i in range(1, n):
            if nums[i] != curr_val:
                curr += 1
                nums[curr] = nums[i]
                curr_val = nums[i]
        return curr + 1

if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    nums = [0,0,1,1,2,2,2,3,3,4,4,4]
    output = sol.removeDuplicates(nums)

    # end time
    end = time.time()

    print("Output     : {} {} \nTime Taken : {}".format(output, nums, end - start))
