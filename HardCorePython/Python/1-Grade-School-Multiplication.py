x = int(input())
y = int(input())

result = 0

count = 0
while(y>0):
    a = y%10
    b = pow(10, count)
    result = result + (x * a)*b
    count =count + 1
    y = int(y/10)

print(result)