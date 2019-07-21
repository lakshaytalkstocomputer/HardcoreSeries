def arith(x, y):
    if x == 0:
        return y
    else:
        return arith(x-1, y+1)

print(arith(3, 4))