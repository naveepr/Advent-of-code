#!/usr/bin/python

def solve(str):
    wrong_str = ['ab','cd','pq','xy']
    vowel = ['a','e','i','o','u']
    count = 0
    repeat = False
    if any(x in str for x in wrong_str):
        return False
    for i in range(len(str)):
        if str[i] in vowel:
            count = count +1 
        if i>0 and str[i] == str[i-1]:
            repeat = True

    return repeat and (count>=3) 
 
if __name__ == "__main__":
    print('enter string') 
    str = input()
    ret = solve(str)
    if ret is True:
        print('nice string')
    else:
        print('not nice string')

