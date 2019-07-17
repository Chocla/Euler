from itertools import combinations
from time import time

#return list of lists containing primes of length i in list[i]
def seive(n):
    A = [True]*(n+1)
    i = 2
    while i*i < n:
        if A[i]:
            j = 2*i
            while j <= n:
                A[j] = False
                j += i
        i += 1
    primes = []
    l = 1
    pTmp = []
    for i in range(2,n+1):
        if A[i]:
            if len(str(i)) > l:            
                primes.append(pTmp)
                pTmp = []
                l += 1
            else:
                pTmp.append(str(i))
    primes.append(pTmp)
    return primes
def split(p,d):
    a,b = "",""
    for i in range(len(p)):
        if i in d:
            b += p[i]
        else: 
            a += p[i]
    return a,b

def reconfigureNum(f,d,primes):
    for i in range(0,10):
        newStr = [""]*(len(f)+len(d))
        j = 0
        for k in range(len(newStr)):
            if k in d:
                newStr[k] = str(i)
            else:
                newStr[k] = f[j]
                j += 1
        newStr = "".join(newStr)
        if newStr in primes:
            return newStr
    return ""

def findAnswer(target,seiveMax,pLen):
    primes = seive(seiveMax)
    length = pLen
    primes = primes[length-1:][0]

    for k in range(0,length - 1):
        digitCombs = list(combinations(range(length-1),k ))

        for d in digitCombs:
            freq = {}
            for p in primes:
                a,b = split(p,d)
                if len(set(b)) == 1:
                    if not a in freq:
                        freq[a] = 1
                    else:
                        freq[a] = freq[a] + 1

            for f in freq:
                if freq[f] == target:
                    return reconfigureNum(f,d,primes)
    return -1

t0 = time()
ans = findAnswer(8,1000000,6)
print("Answer:",ans,"\nTime:",time() - t0)
