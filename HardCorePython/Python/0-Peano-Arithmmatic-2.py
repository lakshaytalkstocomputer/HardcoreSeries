def sum(x , y):
    if(x == 0):
        return y
    else:
        print("sum({0},{1})".format(x, y))
        return (1+ sum(x-1, y ))

if __name__ == "__main__":
    print(sum(3,4))