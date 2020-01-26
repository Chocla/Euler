from collections import defaultdict
from time import time

def createFactorCombinations2(n):
    res = []
    def helper(tmp,n,start):
        if n <= 1 and len(tmp) > 1:
            res.append(tmp[:])
        for i in range(start,n+1):
            if n % i == 0:
                tmp.append(i)
                helper(tmp,n // i, i)
                tmp.pop(len(tmp) - 1)

    helper([],n,2)
    return res

def findValidKs(N):
    fc = createFactorCombinations2(N)
    ks = []
    for f in fc:
        if [] not in f:

            ks.append(N + len(f) - sum(f))
    return ks

#uBound - 1 elements inserted
def findAnswer(uBound):
    i = 4
    count = 0
    minKs = defaultdict(int)
    while count != uBound - 1:
        ks = findValidKs(i)
        for k in ks:
            if minKs[k] == 0 and k <= uBound and k > 1:
                minKs[k] = i
                count += 1
        i += 1
    return sum(set(minKs.values()))
t0 = time()
print("Answer:",findAnswer(12000),"\nTime:",time() - t0)
