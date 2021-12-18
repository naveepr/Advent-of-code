package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) map[string][]string {
	var str []string
	matrix := make(map[string][]string)

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		str = strings.Split(s, "-")

		matrix[str[0]] = append(matrix[str[0]], str[1])
		matrix[str[1]] = append(matrix[str[1]], str[0])
	}

	return matrix
}

func DFSUtil(matrix map[string][]string, visited map[string]int, start string, res *int, part1 bool, smallcave string) {
	tempStart := start

	if start == "end" {
		if part1 || (visited[smallcave] == 0) {
			*res++
		}
		return
	}

	if start == "start" || (start != "" && (start == strings.ToLower(tempStart)) && (visited[start] <= 0)) {
		return
	}

	if start == "" {
		start = "start"
	}
	visited[start]--

	for _, val := range matrix[start] {
		DFSUtil(matrix, visited, val, res, part1, smallcave)
	}

	visited[start]++

}

func Solve(matrix map[string][]string, part1 bool) int {
	visited := make(map[string]int)
	var res1 int

	for k := range matrix {
		visited[k] = 1
	}
	DFSUtil(matrix, visited, "", &res1, part1, "")
	if !part1 {
		for k := range matrix {
			visited[k] = 1
		}
		for k := range matrix {
			tempk := k
			if k != "start" && k != "end" && k == strings.ToLower(tempk) {
				visited[k] = 2
				DFSUtil(matrix, visited, "", &res1, part1, k)
				visited[k] = 1
			}
		}
	}
	return res1
}

func main() {
	arg := os.Args
	part1 := true

	if len(arg) == 3 && arg[2] == "part2" {
		part1 = false
	}
	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := parseInput(scanner)

	for k, matrixRow := range input {
		fmt.Println(k, "--->", matrixRow)
		// for _, val := range matrixRow {
		// 	fmt.Printf("val is %v tye %T", val, val)
		// }
	}
	ans := Solve(input, part1)

	fmt.Println("Answer is ", ans)
}
