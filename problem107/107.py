import numpy as np

#Something like Dijkstra's Algorithm
def findMinTree(matrix):
    visited = [0]
    unvisited = range(1,len(matrix))
    minpath = 0
    while unvisited != []:
        minval = 10000
        mincol = 0
        for row in visited:
            for i in unvisited:
                if matrix[row][i] != 0 and minval > matrix[row][i]:
                    minval = matrix[row][i]
                    mincol = i
                    minrow = row
        print minval,mincol,visited,unvisited
        matrix[minrow][mincol] = 0
        visited.append(mincol)
        unvisited.remove(mincol)
        minpath += minval
    return minpath
#turn text file into weight matrix
def parseInput(path):
    matrix = []
    lines = open(path,'r').read().splitlines()
    for l in lines:
        tmp = l.split(',')
        for i in range(len(tmp)):
            if tmp[i] != '-':
                tmp[i] = int(tmp[i])
            else:
                tmp[i] = 0
        matrix.append(tmp)
    return matrix
#sum upper triangle values
def totalweight(matrix):
    tot = 0
    for i in range(len(matrix)):
        for j in range(i,len(matrix)):
            tot += matrix[i][j]
    return tot
def solve107(path):
    matrix = parseInput(path)
    total = totalweight(matrix)
    minweight = findMinTree(matrix)
    print "Answer: ", (total - minweight)

solve107('example.txt')
solve107('p107_network.txt')