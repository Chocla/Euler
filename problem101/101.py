import numpy as np 
from scipy import interpolate
def u(n):
    return 1 - n + n**2 - n**3 + n**4 -n**5 +n**6-n**7+n**8 -n**9 +n**10
def ex(n):
    return n**3
def findAnswer(u,deg):
    x = np.arange(1,deg+1,dtype=np.int64)
    y = u(x)
    fitsum = 1
    for i in range(2,len(x)+1):
        fitsum += np.round(interpolate.barycentric_interpolate(x[:i],y[:i],i+1))
  
    return fitsum
print "Example Sum:",int(findAnswer(ex,3))
print "Answer:",int(findAnswer(u,10))
