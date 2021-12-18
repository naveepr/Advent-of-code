package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Stack struct {
	valList []string
	top     int
}

func (st *Stack) push(c string) {
	st.valList = append(st.valList, c)
	st.top += 1
}

func (st *Stack) pop() {
	if st.top == 0 {
		st.valList = []string{}
	} else {
		st.valList = st.valList[:st.top]
	}
	st.top -= 1
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) [][]string {
	var input [][]string

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		str := strings.Split(s, "")
		input = append(input, str)
	}

	return input
}

func isOpenBrace(c string) bool {

	return c == "(" || c == "[" || c == "{" || c == "<"
}

func isCloseBrace(c string) bool {

	return c == ")" || c == "]" || c == "}" || c == ">"
}

func Solve(inputptr *[][]string) (int, int) {
	var res1, res2 int
	input, corrupt := *inputptr, false
	var autoScore []int

	braceMap := map[string]string{")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}
	braceVal := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	autoVal := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	st := &Stack{
		top: -1,
	}

	for _, line := range input {
		st.valList = []string{}
		st.top = -1
		corrupt = false
		for _, c := range line {
			if isOpenBrace(c) {
				st.push(c)
			} else {
				if st.top == -1 {
					res1 += braceVal[c]
					corrupt = true
					break
				}
				if braceMap[c] == st.valList[st.top] {
					st.pop()
				} else {
					corrupt = true
					res1 += braceVal[c]
					break
				}
			}
		}

		res2 = 0
		if !corrupt {
			for i := st.top; i >= 0; i-- {
				val := st.valList[i]
				res2 *= 5
				res2 += autoVal[val]
			}
			autoScore = append(autoScore, res2)
		}
	}

	sort.Ints(autoScore)
	res2 = autoScore[len(autoScore)/2]
	return res1, res2
}

func main() {
	arg := os.Args

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := parseInput(scanner)

	ans1, ans2 := Solve(&input)

	fmt.Println("Answer is ", ans1, ans2)
}
