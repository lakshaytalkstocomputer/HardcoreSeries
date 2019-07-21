# Take two numbers as string
# Perform Grade school Multiplication on it


# Returns 2 Strings from USER in the form (x, y)

def take_input():
    return (input(), input())


def multiply(i, x):

    ## first reverse the elements
    z = x[::-1]

    ans = str()

    ## take every element of z
    ## convert it into int()
    ## multiply it with int() of i
    ## take the carry by dividing the  result with 10
    ## and taking the int() of that
    ## add that carry to the result of next multipllication
    ## store the modulus 10 of result to string
    i = int(i)
    carry = 0
    for j in z:
        res = i * int(j) + carry
        #print(res)
        carry = int(res/10)
        ans += str(res % 10)

    if carry:
        ans += str(carry)

    ##print(ans[::-1])
    return ans[::-1]


def print_answer():

    # take input
    x, y = take_input()

    # testing lines
    print("length of x :" + str(len(x)))
    print("length of y :" + str(len(y)))

    # for every digit in y multiply it with x
    # store the result in result
    ## first reverse the elements of y
    y = y[::-1]

    # testing
    print("length of y :" + str(len(y)))
    print("y: " + y)

    ## multiplying the digits
    ## and storing the answer in -> result

    result = []
    count = 0
    for i in y:
        result.append(multiply(i, x)+ ("0"*count))
        count = count + 1

    print(result)


if (__name__ == '__main__'):
    print_answer()

