import time
#Calculate how many pythagorean triples there are for each "length" n = a + b + c
def generateFrequencyList(maxVal):
    def gcd(a,b):
        while b != 0:
            t = b
            b = a % b
            a = t
        return a
    freq = {}
    m = 0
    c = 2
    while c < maxVal:
        for n in range(1,m):
            #generates primitive pythagorean triples provided a,c are integers and gcd(a,b) = 1
            a = (m*m - n*n)/2
            b = m*n
            c = (m*m + n*n)/2

            if a == int(a) and c == int(c) and gcd(a,b) == 1:
                k = 1
                l = int(a+b+c)
                #Find the triples that are multiples of l
                while k*l < maxVal:
                    freq[k*l] = freq.get(k*l, 0) + 1
                    k += 1
             
            if c > maxVal:
                break
        m += 1    

    return freq
def findAnswer(maxVal):
    freqs = generateFrequencyList(maxVal)
    count = 0
    for f in freqs:
        if freqs[f] == 1:
            count += 1
    return count


t0 = time.time()    
ans = findAnswer(1500001)
print("Answer",ans," Time: ",time.time() - t0)