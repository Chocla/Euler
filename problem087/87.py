from time import time
from math import sqrt
def findAnswer(primes,n):
    seen = {}
    for a in primes:
        if a**4 > n:
            break
        for b in primes:
            if a**4 + b**3 > n:
                break
            for c in primes:
                if a**4 + b**3 + c**2 > n:
                    break
                tmp = int(a**4 + b**3 + c**2)
                seen[tmp] = 1
    return(len(seen))

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
    for i in range(2,n+1):
        if A[i]:
            primes.append(i)
    return primes    


t0 = time()
BOUND = 50_000_000
#Biggest Prime we need to check is if it is the square term and 2 is the other two terms
PRIME_BOUND = int(sqrt(BOUND - 2**4 - 2**3)) + 1 
primes = seive(PRIME_BOUND)
ans = findAnswer(primes,BOUND)
print("Answer:",ans,"\nTime:",time() - t0)
