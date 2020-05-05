#!/usr/bin/python3
import re


def parse(line, point):
    point = re.findall('[\d]+', line)
    print(point)
    if "on" in line:
        return 1
    elif "off" in line:
        return 0
    else:
        return 2


def solve(f, matrix):
    point = []
    while 1:
        line = f.readline()
        if not line:
            break
        point.clear()
        val = parse(line, point)
        if val == 0:
            for row in range(point[0], point[2] + 1):
                for col in range(point[1], point[3] + 1):
                    matrix[row][col] = 0
        elif val == 1:
            for row in range(point[0], point[2] + 1):
                for col in range(point[1], point[3] + 1):
                    matrix[row][col] = 1
        elif val == 2:
            for row in range(point[0], point[2] + 1):
                for col in range(point[1], point[3] + 1):
                    matrix[row][col] = ~matrix[row][col]


if __name__ == "__main__":
    print("Enter fileName")
    fileName = input()
    matrix = [[0 for i in range(1000)] for i in range(1000)]

    with open(fileName) as f:
        solve(f, matrix)

    f.close()
