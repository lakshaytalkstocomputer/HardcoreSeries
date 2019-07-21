def sqrt(x):
    def average(a, b):
        return (a+b)/2
    
    def improve(guess, a):
        return average(guess, a/guess)

    def absol(c):
        if (c >= 0):
            return c
        elif c < 0:
            return -c
    
    def use(guess):
        # print("guess: " + str(guess))
        if(absol(guess*guess - x)) <= 0.0001:
            # print("final guess: " + str(guess))
            return guess
        else:
            return use(improve(guess, x))       
        
    return use(1)
    
print(sqrt(9))
    
    