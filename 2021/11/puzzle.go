package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) [10][10]int {
	var str []string
	var matrix [10][10]int

	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		str = strings.Split(s, "")
		for j, val := range str {
			num, err := strconv.Atoi(val)
			check(err)
			matrix[i][j] = num
		}
	}

	return matrix
}

func DFSUtil(matrix *[10][10]int, row, col int, visited *[10][10]bool, blink *int) {

	if row < 0 || col < 0 || row >= 10 || col >= 10 || visited[row][col] {
		return
	}

	matrix[row][col] = (matrix[row][col] + 1) % 10

	if matrix[row][col] == 0 {
		(*blink)++
		visited[row][col] = true

		DFSUtil(matrix, row-1, col-1, visited, blink)
		DFSUtil(matrix, row-1, col, visited, blink)
		DFSUtil(matrix, row-1, col+1, visited, blink)
		DFSUtil(matrix, row, col-1, visited, blink)
		DFSUtil(matrix, row, col+1, visited, blink)
		DFSUtil(matrix, row+1, col-1, visited, blink)
		DFSUtil(matrix, row+1, col, visited, blink)
		DFSUtil(matrix, row+1, col+1, visited, blink)
	}

}

func Solve(matrixptr *[10][10]int, steps int, part1 bool) (int, int) {
	matrix := matrixptr
	var visited [10][10]bool
	var blink, maxblink, step int

	for step = 0; step < steps; step++ {
		maxblink = blink
		for i, matrixRow := range matrix {
			for j := range matrixRow {
				DFSUtil(matrixptr, i, j, &visited, &blink)
			}

		}
		if !part1 && (blink-maxblink) == 100 {
			break
		}
		visited = [10][10]bool{}
	}

	fmt.Println(matrix)
	return blink, step + 1
}

func main() {
	arg := os.Args
	steps := math.MaxInt
	part1 := true

	if arg[2] == "part2" {
		part1 = false
	}

	if part1 {
		steps, _ = strconv.Atoi(arg[3])
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	matrix := parseInput(scanner)

	ans1, ans2 := Solve(&matrix, steps, part1)

	fmt.Println("Answer is ", ans1, ans2)
}
