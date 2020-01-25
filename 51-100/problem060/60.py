from collections import defaultdict
import networkx as nx
from time import time
uBound = 100000000
LEN = 9
def seive(N):
    A = [True]*(N+1)
    i = 2
    while i*i < N:
        if A[i]:
            j = 2*i
            while j <= N:
                A[j] = False
                j += i
        i += 1
    pDict = defaultdict(bool)
    pList = [] 
    for i in range(2,N+1):
        if A[i]:
            pDict[i] = True
            pList.append(i)
    return pDict,pList

def isPrime(n):
    return pDict[n]
def checkTwo(p1,p2):
    c1 = int(str(p1) + str(p2))
    c2 = int(str(p2) + str(p1))
    return isPrime(c1) and isPrime(c2)

t0 = time()
pDict,pList = seive(uBound)
print("Done seiving",len(pList), "primes in ",time()-t0)

G = nx.Graph()

for i,p1 in enumerate(pList):
    for j in range(i+1,len(pList)):
        if len(str(p1)) + len(str(pList[j])) >= LEN:
            break
        if checkTwo(p1,pList[j]):
            G.add_edge(p1,pList[j])

print("Graph constructed with",len(G.edges()), "edges in ", time() - t0)
for x in nx.find_cliques(G):
    if len(x) == 5:
        print(x, sum(x))
print(time()-t0)