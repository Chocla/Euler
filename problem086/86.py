from math import sqrt
def findNextCM(M):
    c = M
    count = 0
    for a in range(1,M+1):
        for b in range(a,M+1):
            l1 = sqrt(c*c + (a+b)*(a+b))
            if int(l1) == l1:
                count += 1
                # print(a,b,c,int(l1))
    return count

def findAnswer(target):
    CM = 0
    M = 0
    while CM < target:
        M += 1
        CM += findNextCM(M)
        if M % 100 == 0:
            print(M,CM)
    return M,CM
print(findAnswer(1000000))
#Unoptimized, takes about 20 minutes to run