import numpy as np
from scipy.sparse.csgraph import floyd_warshall
from time import time
f = open("p082_matrix.txt","r")
# f = open("test","r")

s = [[int(y) for y in x.split(",")] for x in f.read().split("\n")]

#Slow implementation of Floyd_Warshall shortest path algorithm, used scipy instead for time saving
def floyd(dist,n):
    for k in np.arange(len(dist)):
        for i in np.arange(len(dist)):
            for j in np.arange(len(dist)):
                if dist[i][j] > dist[i][k] + dist[k][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    return dist

def createWeights(s):
    n = len(s)
    n2 = n**2
    dist = np.full((n2,n2),np.inf)
    for i in np.arange(n2):
        if i % n != n-1:
            dist[i][i] = 0 
        else:
            dist[i][i] = s[i // n][i % n]
    for q in range(len(s)):
        for r in range(len(s)):
            if q < n-1: 
                dist[n*q+r][n*(q+1) + r] = s[q][r]
            if q > 0:
                dist[n*q + r][n*(q-1) + r] = s[q][r]
            if r < n-1:
                dist[n*q + r][n*q + r + 1] = s[q][r]
            # if r > 1:
            #     dist[5*q + r][5*q + r - 1] = s[q][r]
    return dist

def findMinPath(dist,s):
    n = len(s)
    minPath = np.inf
    for i in range(0,n**2,n):
        for j in range(n-1,n**2,n):
            dist[i][j] += s[j // n][j % n]
            if dist[i][j] < minPath:
                minPath = dist[i][j]
    return minPath

dist = createWeights(s)
print("Weights Created")
t0 = time()
dist = floyd_warshall(dist)
print("Paths Found: ",time() - t0)
print(findMinPath(dist,s))





