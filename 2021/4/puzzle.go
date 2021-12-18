package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) ([][][]string, []string) {
	var matrixList [][][]string
	var randomList []string
	var line string

	if scanner.Scan() {
		line = scanner.Text()
		randomList = strings.Split(line, ",")
	}

	for scanner.Scan() {
		matrix := make([][]string, 5)
		for i := range matrix {
			scanner.Scan()
			line = scanner.Text()
			space := regexp.MustCompile(`\s+`)
			line = space.ReplaceAllString(line, " ")
			line = strings.TrimSpace(line)
			// fmt.Println(line)
			matrix[i] = strings.Split(line, " ")
		}

		matrixList = append(matrixList, matrix)
	}

	// fmt.Println(matrixList)
	return matrixList, randomList
}

func searchMatrix(Matrixptr *[][]string, searchNum string) (int, int, bool) {
	var Matrix [][]string = *Matrixptr
	var rowIndex, colIndex int
	var row []string
	var num string

	for rowIndex, row = range Matrix {
		// fmt.Println("Naveen2")
		for colIndex, num = range row {
			// fmt.Println(colIndex, num)
			if num == searchNum {
				row[colIndex] = "X"
				return rowIndex, colIndex, true
			}
		}
	}

	// fmt.Println("Naveen")
	return rowIndex, colIndex, false
}

func isBingo(Matrixptr *[][]string, rowIndex int, colIndex int) bool {
	var Matrix [][]string = *Matrixptr
	var i int

	for _, row := range Matrix {
		for i = 0; i < len(row); i++ {
			if row[i] != "X" {
				break
			}
		}

		if i == len(row) {
			return true
		}
	}

	for i = 0; i < len(Matrix); i++ {
		// fmt.Println(i, colIndex)
		if Matrix[i][colIndex] != "X" {
			break
		}
	}

	if i == len(Matrix) {
		return true
	}

	return false
}

func getSum(Matrixptr *[][]string) int {
	var Matrix [][]string = *Matrixptr
	sum := 0

	for _, i := range Matrix {
		for _, s := range i {
			if s != "X" {
				num, err := strconv.Atoi(s)
				check(err)
				sum += num
			}
		}
	}

	return sum
}

func Solve(MatrixListptr *[][][]string, randomNum []string) (int, int) {
	min_sum, max_sum, min_moves, max_moves, num_moves := 0, 0, len(randomNum), 0, 0
	var matrixList [][][]string = *MatrixListptr

	for _, i := range matrixList {
		num_moves = 0
		for _, val := range randomNum {
			num_moves++
			rowIndex, colIndex, found := searchMatrix(&i, val)
			if found && isBingo(&i, rowIndex, colIndex) {
				num, err := strconv.Atoi(val)
				check(err)
				if num_moves < min_moves {
					min_moves = num_moves
					min_sum = num * getSum(&i)
				}
				if num_moves > max_moves {
					max_moves = num_moves
					max_sum = num * getSum(&i)
				}
				break

			}
		}
	}
	return min_sum, max_sum
}

func main() {
	arg := os.Args

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	matrixList, randomList := parseInput(scanner)

	ans1, ans2 := Solve(&matrixList, randomList)

	fmt.Println("Answer is ", ans1, ans2)
}
