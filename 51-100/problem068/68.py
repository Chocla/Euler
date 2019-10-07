from itertools import permutations


def setString(N,p):
    v = [str(x) for x in p]
    s = ""
    for i in range(N):
        s += v[N+i] + v[i] + v[(1+i) % N]
    return s
def isMagicNGon(N,p):
    for i in range(N+1,2*N):
        if p[N] > p[i]:
            return False
    s = p[N] + p[0] + p[1]
    for i in range(N-1):
        if s != p[N + 1 + i] + p[1 + i] + p[(2 + i) % N]:
            return False
    return True

def findAllMagicNGons(N):
    perms = list(permutations(range(1,2*N + 1)))
    sets = []
    for p in perms:
        if isMagicNGon(N,p):
            sets.append(setString(N,p))
    return sets

candidates = sorted(findAllMagicNGons(5))
print("Answer: ",candidates[-1])
