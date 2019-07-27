def fibonacci(n):
    if n < 2:
        return n
    else:
        return fibonacci(n-1) + fibonacci(n-2)


def fibonacci2(n, i):
    print("->"*i + " " + str(n))
    if n < 2:
        return n
    else:
        return fibonacci2(n-1, i+1) + fibonacci2(n-2, i+1)


print([fibonacci(i) for i in range(10)])
print(fibonacci2(5, 0))
