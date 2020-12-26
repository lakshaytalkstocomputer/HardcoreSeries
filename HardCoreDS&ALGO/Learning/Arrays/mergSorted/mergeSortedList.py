import time
class Solution:
    def solve(self, lst0, lst1):
        len_lst_0 = len(lst0)
        len_lst_1 = len(lst1)
        min_list = lst0 if len_lst_0 < len_lst_1 else lst1
        min_list_len = len(min_list)

        new_lst = list()

        lst_0_index = 0
        lst_1_index = 0

        while(lst_0_index < len_lst_0 and lst_1_index < len_lst_1 ):
            if lst0[lst_0_index] < lst1[lst_1_index]:
                new_lst.append(lst0[lst_0_index])
                lst_0_index += 1
            else:
                new_lst.append(lst1[lst_1_index])
                lst_1_index += 1

        if lst_0_index < len_lst_0 :
            for i in range(lst_0_index, len_lst_0):
                new_lst.append(lst0[i])

        if lst_1_index < len_lst_1 :
            for i in range(lst_1_index, len_lst_1):
                new_lst.append(lst1[i])

        return new_lst
if __name__ == '__main__':
    sol = Solution()

    # starting time
    start = time.time()

    num1 = [1]
    num2 = [1]
    output = sol.solve(num1, num2)

    # end time
    end = time.time()

    print("Output     : {} \nTime Taken : {}".format(output, end - start))