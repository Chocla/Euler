import numpy as np

def countUpToLengthN(n):

    counts = np.ones(10,dtype=np.uint64)
    counts[0] = 0
    total = 0
    for i in range(2,n+1):
        for j in range(0,10):
            counts[j] = sum(counts[j:])
        total += 2*sum(counts[1:]) + counts[0] - 9
    return total + 9

print(countUpToLengthN(100))