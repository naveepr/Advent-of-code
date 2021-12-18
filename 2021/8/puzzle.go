package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func parseInput(scanner *bufio.Scanner) ([][]string, [][]string) {
	var input [][]string
	var num [][]string

	for scanner.Scan() {
		s := scanner.Text()
		str := strings.Split(s, "|")
		a := strings.TrimSpace(str[0])
		tempinput := strings.Split(a, " ")
		b := strings.TrimSpace(str[1])
		tempnum := strings.Split(b, " ")

		input = append(input, tempinput)
		num = append(num, tempnum)
	}

	return num, input
}

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

// subset returns true if the first array is completely
// contained in the second array. There must be at least
// the same number of duplicate values in second as there
// are in first.
func subset(first, second string) bool {
	set := make(map[rune]int)
	for _, value := range second {
		set[value]++
	}

	for _, value := range first {
		if _, found := set[value]; !found {
			return false
		}
	}

	return true
}

func diff(first, second string) string {
	var str string
	set := make(map[rune]int)
	for _, value := range first {
		set[value]++
	}

	for _, value := range second {
		if _, found := set[value]; !found {
			str += string(value)
		}
	}

	return str
}

/*
 * 0 -> 6
 * 1 -> 2
 * 2 -> 5
 * 3 -> 5
 * 4 -> 4
 * 5 -> 5
 * 6 -> 6
 * 7 -> 3
 * 8 -> 7
 * 9 -> 6
 */
func decodeNonUniqueDigit(fiveDigitWiring []string, sixDigitWiring []string, numMapPtr *map[int]string) {
	numMap := *numMapPtr
	var threepos, ninepos, sixpos int
	var leftbot string

	for i, num := range fiveDigitWiring {
		if subset(numMap[1], num) {
			numMap[3] = num
			threepos = i
		}
	}

	fiveDigitWiring = append(fiveDigitWiring[:threepos], fiveDigitWiring[threepos+1:]...)
	for i, num := range sixDigitWiring {
		if subset(numMap[4], num) {
			numMap[9] = num
			ninepos = i
			continue
		}

		if !subset(numMap[1], num) {
			numMap[6] = num
			sixpos = i
		}
	}

	if sixpos > ninepos {
		sixDigitWiring = append(sixDigitWiring[:sixpos], sixDigitWiring[sixpos+1:]...)
		sixDigitWiring = append(sixDigitWiring[:ninepos], sixDigitWiring[ninepos+1:]...)
	} else {
		sixDigitWiring = append(sixDigitWiring[:ninepos], sixDigitWiring[ninepos+1:]...)
		sixDigitWiring = append(sixDigitWiring[:sixpos], sixDigitWiring[sixpos+1:]...)
	}

	numMap[0] = sixDigitWiring[0]
	leftbot = diff(numMap[9], numMap[8])

	numMap[5] = diff(leftbot, numMap[6])
	for _, num := range fiveDigitWiring {
		if num != numMap[5] {
			numMap[2] = num
		}
	}
}

func Solve(numptr *[][]string, inputptr *[][]string, part1 bool) int {
	numbers, inputs := *numptr, *inputptr
	numMap := make(map[int]string, 10)
	var res int = 0
	var sixDigitWiring, fiveDigitWiring []string

	for i, ins := range inputs {
		for k := range numMap {
			delete(numMap, k)
		}
		fiveDigitWiring, sixDigitWiring = []string{}, []string{}
		for _, input := range ins {
			l := len(input)
			switch l {
			case 2:
				numMap[1] = sortString(input)
			case 3:
				numMap[7] = sortString(input)
			case 4:
				numMap[4] = sortString(input)
			case 7:
				numMap[8] = sortString(input)
			case 5:
				fiveDigitWiring = append(fiveDigitWiring, sortString(input))
			case 6:
				sixDigitWiring = append(sixDigitWiring, sortString(input))
			}
		}

		if !part1 {
			decodeNonUniqueDigit(fiveDigitWiring, sixDigitWiring, &numMap)
		}
		nums := numbers[i]
		tempres := 0
		for _, num := range nums {
			num = sortString(num)
			if part1 {
				if num == numMap[1] || num == numMap[4] || num == numMap[8] || num == numMap[7] {
					res++
				}
			} else {
				for k, v := range numMap {
					if v == num {
						tempres *= 10
						tempres += k
					}
				}
			}
		}

		res += tempres
	}

	return res
}

func main() {
	arg := os.Args
	part1 := true
	if len(arg) == 3 && arg[2] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	num, input := parseInput(scanner)

	ans := Solve(&num, &input, part1)

	fmt.Println("Answer is ", ans)
}
