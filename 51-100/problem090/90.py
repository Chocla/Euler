from itertools import combinations

squares = ["01", "04", "09", "16", "25", "36", "49", "64", "81"]

def findAnswer():
    combs = list(combinations(range(0,10),6))
    count = 0
    for i in combs:
        for j in combs:
            if valid(i,j):
                count += 1
    #this double counts the number of valid combinations.
    return count // 2

def valid(A,B):
    A = augment(A)
    B = augment(B)
    for s in squares:
        # if any of the square numbers cannot be made by taking the first digit from A and second digit from B,
        # or vice versa, then it isn't a valid configuration
        if not (((int(s[0]) in A) and (int(s[1]) in B )) or ((int(s[1]) in A) and (int(s[0]) in B))):
            return False
    return True

# 6 is equivalent to 9, so if either is in the set add the other one.
def augment(A):
    A = set(A)
    if 6 in A or 9 in A:
        A.add(9)
        A.add(6)
    return A

print(findAnswer())