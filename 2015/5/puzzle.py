#!/usr/bin/python

def solve(str, part1):
    wrong_str = ['ab','cd','pq','xy']
    vowel = ['a','e','i','o','u']
    count = 0
    repeat = False
    recur = False
    if part1:
        if any(x in str for x in wrong_str):
            return False
        for i in range(len(str)):
            if str[i] in vowel:
                count = count +1 
            if i>0 and str[i] == str[i-1]:
                repeat = True

        return repeat and (count>=3) 
    else:
        for i in range(len(str)-2):
            if not repeat and -1 != str.find(str[i:i+2],i+2):
                repeat = True
            
            if not recur and str[i] == str[i+2]: 
                recur = True

            if repeat and recur:
                return True

        return False
        
if __name__ == "__main__":
    print("Enter the input file")
    fileName = input()
    print("Enter part 1 or 2")
    part1 = input()
    count=0
    with open(fileName) as f:
        for line in f:
            ret = solve(line, (int(part1)==1))
            if ret is True:
                count = count+1 
    
    print('count is {:d}'.format(count))
