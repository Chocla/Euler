from itertools import combinations,permutations

operations = [
    lambda x,y: x + y,
    lambda x,y: x - y,
    lambda x,y: x * y,
    lambda x,y: x / y
]
parens = [
    lambda f1,f2,f3,a,b,c,d: f2(f1(a, b), f3(c, d)),
    lambda f1,f2,f3,a,b,c,d: f3(f2(f1(a, b), c), d),
    lambda f1,f2,f3,a,b,c,d: f3(f1(a, f2(b, c)), d),
    lambda f1,f2,f3,a,b,c,d: f1(a, f2(b, f3(c, d))),
    lambda f1,f2,f3,a,b,c,d: f1(a, f3(f1(b, c),d )),
]
def computeAllOperations(combo):
    results = []
    ps = list(permutations(combo))
    for (a,b,c,d) in ps:
        for f1 in operations:
            for f2 in operations:
                for f3 in operations:
                    for paren in parens:
                        try:
                            tmp = paren(f1,f2,f3,a,b,c,d)
                        except:
                            continue #divide by zero error won't give a valid result, continue
                        if int(tmp) == tmp and tmp > 0 and tmp not in results:
                            results.append(int(tmp))
    return sorted(results)
def score(combo):
    results = computeAllOperations(combo)
    for i in range(1,len(results)):
        if results[i] - results[i-1] != 1:
            return i
def findAnswer():
    combos = list(combinations(range(1,10),4))
    highScore = 0
    highCombo = (0,0,0,0)
    for combo in combos:
        candidateScore = score(combo)
        if candidateScore > highScore:
            highScore = candidateScore
            highCombo = combo
    return(highCombo,highScore)
print(findAnswer())