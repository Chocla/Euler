def squareDigits(N):
    s = int(0)
    while N > 0:
        s += (N%10)**2
        N /= 10
    return s

def generateSquares(size):
    squares = [0]*(size+1)
    for i in range(size+1):
        squares[i] = squareDigits(i)
    return squares

def parseSquares():
    for i in range(1,len(squares)):
        if squares[i] == 1 or squares[i] == 89:
            continue
       
        tmpList = []
        j = i
        while True:
            tmpList.append(j)
            j = squares[j]
            if j > len(squares):
                print "Incoming Error:", i,j
            if j == 1 or j == 89:
                break
        for t in tmpList:
            squares[t] = j    
    return 0

squares =  generateSquares(10000000)
parseSquares()
count = 0
for s in squares:
    if s == 89:
        count+= 1
print count