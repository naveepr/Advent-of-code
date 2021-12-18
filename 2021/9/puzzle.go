package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) ([][]int, int, int) {
	var input []int
	var str []string
	var matrix [][]int
	var rowsz, colsz int

	for scanner.Scan() {
		rowsz++
		input = []int{}
		str = strings.Split(scanner.Text(), "")
		for _, val := range str {
			num, err := strconv.Atoi(val)
			check(err)
			input = append(input, num)
		}
		if colsz == 0 {
			colsz = len(str)
		}
		matrix = append(matrix, input)
	}

	return matrix, rowsz, colsz
}

func DFSUtil(matrixptr *[][]int, rowsz, colsz, row, col, low int, sumptr *int, first bool) {
	matrix := *matrixptr
	if row < 0 || col < 0 || row == rowsz || col == colsz || matrix[row][col] >= 9 || (!first && matrix[row][col] <= low) {
		return
	}

	*sumptr += 1

	low = matrix[row][col]
	matrix[row][col] += 100

	DFSUtil(&matrix, rowsz, colsz, row-1, col, low, sumptr, false)
	DFSUtil(&matrix, rowsz, colsz, row+1, col, low, sumptr, false)
	DFSUtil(&matrix, rowsz, colsz, row, col-1, low, sumptr, false)
	DFSUtil(&matrix, rowsz, colsz, row, col+1, low, sumptr, false)

	// matrix[row][col] -= 100
}

func Solve(matrixptr *[][]int, rowsz int, colsz int) (int, int) {
	matrix := *matrixptr
	var res, sum int
	list := []int{1, 1, 1}
	prodsum := 1

	for i, matrixRow := range matrix {
		for j, matrixCol := range matrixRow {
			if i > 0 && matrix[i-1][j] <= matrixCol {
				continue
			}
			if i < rowsz-1 && matrix[i+1][j] <= matrixCol {
				continue
			}
			if j > 0 && matrix[i][j-1] <= matrixCol {
				continue
			}
			if j < colsz-1 && matrix[i][j+1] <= matrixCol {
				continue
			}
			res += (matrixCol + 1)
			sum = 0
			DFSUtil(matrixptr, rowsz, colsz, i, j, matrixCol, &sum, true)
			if sum > list[0] {
				list[0] = sum
				sort.Ints(list)
			}
		}
	}

	for _, val := range list {
		prodsum *= val
	}
	return res, prodsum
}

func main() {
	arg := os.Args

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input, rowsz, colsz := parseInput(scanner)

	for _, matrixRow := range input {
		fmt.Println(matrixRow)
	}
	ans1, ans2 := Solve(&input, rowsz, colsz)

	fmt.Println("Answer is ", ans1, ans2)
}
