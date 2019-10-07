from time import time

TARGET = 2_000_000

def findAllInteriorRectangles(n,m):
    c = 0
    for a in range(1,n+1):
        for b in range(1,m+1):
            c += (n-a + 1)*(m-b+1)
    return c

def findGoodRecangle(N,M):
    bestDistance = TARGET
    bestArea = 0
    for n in range(1,N):
        for m in range(1,M):
            candidate = findAllInteriorRectangles(n,m)
            if candidate > TARGET - 100:
                if abs(candidate - TARGET) < bestDistance:
                    bestDistance = abs(candidate - TARGET)
                    bestArea = n*m
                break
    return bestArea
t0 = time()
ans = findGoodRecangle(200,200)
print("Answer:",ans,"Time:",time()-t0)