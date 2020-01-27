#!/usr/bin/python
import hashlib 

def solve(secretKey, part1):
    i=0
    while True:
        word = secretKey + str(i)
        result = hashlib.md5(word.encode())
        out = result.hexdigest()
        if part1 and out[0:5] == '00000':
            return i
        elif not part1 and out[0:6] == '000000':
            return i
        i += 1

if __name__ == "__main__":
    print('Enter the secret key')
    secretKey = input()
    ans = str(solve(secretKey, False))
    print('Answer is {0:s}'.format(ans))
