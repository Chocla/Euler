from copy import copy
from itertools import permutations 

def generateData():
    def findFourDigitPolygonNums(func):
        n  = 1
        vals = []
        while True:
            tmp = func(n)
            n += 1
            if tmp< 1000:
                continue
            if tmp > 9999:
                break
            vals.append(int(tmp))
            
        return vals
    functions = [
        lambda x: int(x*(x+1)/2),
        lambda x: int(x*x),
        lambda x: int(x*(3*x-1)/2),
        lambda x: int(x*(2*x-1)),
        lambda x: int(x*(5*x-3)/2),
        lambda x: int(x*(3*x-2)),
    ]
    data = []
    for f in functions:
        data.append(findFourDigitPolygonNums(f))
    return data


def findChains(data,perm):
    newCandidates = [[j] for j in data[0]]
    for i in perm:
        oldCandidates = copy(newCandidates)
        newCandidates = []
        for j in range(0,len(oldCandidates)):
            tail = oldCandidates[j][-1] % 100
            if tail < 10:
                continue
            for val in data[i]:
                if int(val / 100) == tail:
                    newChain = copy(oldCandidates[j])
                    newChain.append(val)
                    newCandidates.append(newChain)
    return newCandidates


data = generateData()
perms = list(permutations(range(1,6)))
for p in perms:
    candidates = findChains(data, p)
    for t in candidates:
        if len(t) > 4 and int(t[0] / 100) == t[-1] % 100:
            print("Answer:", t, "Sum:", sum(t), "Permutation:", p)
            exit()
