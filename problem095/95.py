
def removeNonChainValues():
    for i in range(len(sVals)):
        if sVals[i] == 1:
            continue
        tmpChain = []
        j = i
        while True:
            
            tmpChain.append(j)
            j = sVals[j]
            #Real chain found, add to list of chains 
            if j in tmpChain:
                #only keep the repeated elements
                realChain = tmpChain[tmpChain.index(j):]
                chains.append(realChain)
                break
            #Chain ended
            if j == 1 or j > len(sVals):

                break
        #set all seen elements to 1 for efficiency
        for element in tmpChain:
            if element > len(sVals):
                continue
            sVals[element] = 1
        i += 1
def findAnswer(chains):
    maxChain = []
    maxLength = 0
    for chain in chains:
        if len(chain) > maxLength:
            maxChain = chain
            maxLength = len(chain)
    return min(maxChain)
def parse():
    f = open("SigmaValues.txt","r") #calculated in main.go
    tmp = f.read().split(",")
    sVals = []
    for i in range(len(tmp)):
        sVals.append(int(tmp[i]))
    return sVals
sVals = parse()
chains = []
removeNonChainValues()
print "Answer:", findAnswer(chains)