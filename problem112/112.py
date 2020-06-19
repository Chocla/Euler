def isBouncy(n):
    if n < 100:
        return False
    s = str(n)
    dec = inc = False
    for i in range(1,len(s)):
        if s[i] < s[i-1]:
            dec = True
        if s[i] > s[i-1]:
            inc = True
        if inc and dec:
            return True
    return False

def findAnswer():
    c = 0
    i = 0
    while True:
        if isBouncy(i):
            c += 1

        if c != 0 and c/i == 99/100:
            return i
        i += 1

print(findAnswer())