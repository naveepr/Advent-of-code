#!/usr/bin/python

def parse(file):
    return file.read().strip()

def solve(content, part1):
    dic = {}
    x1,y1 = 500,500
    x2,y2 = 500,500
    count=0
    for c in content:
        if (part1 or count%2==0):
            x,y=x1,y1
        else:
            x,y=x2,y2

        if c == '<':
            x-= 1
        elif c == '>':
            x+= 1
        elif c == '^':
            y+= 1
        else:
            y-= 1
        dic[(x, y)]= dic.get((x,y),0) + 1

        if (part1 or count%2==0):
            x1,y1 = x,y
        else:
            x2,y2 = x,y

        count +=1

    return len(dic)

if __name__ == "__main__":
    print("Enter the input file")
    fileName = input()
    content = []

    with open(fileName) as f:
        content = parse(f)
        part1=solve(content, True)
        part2=solve(content, False)
        print('part1 -- {:d}, part2 -- {:d} '.format(part1, part2))
    f.close()
