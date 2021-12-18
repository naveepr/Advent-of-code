package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SolvePart1(numbers *[]string) int {
	var onesByte, zeroesByte [12]int
	var numBits, gammaRate, epsilonRate, i int
	var bit rune

	for _, num := range *numbers {
		for i, bit = range num {
			if string(bit) == "1" {
				onesByte[i]++
			} else {
				zeroesByte[i]++
			}
		}
		if numBits == 0 {
			numBits = i + 1
		}
	}

	for i := 0; i < numBits; i++ {
		if onesByte[i] > zeroesByte[i] {
			gammaRate |= (1 << (numBits - 1 - i))
		} else {
			epsilonRate |= (1 << (numBits - 1 - i))
		}
	}

	return gammaRate * epsilonRate
}

func getCarbonOxygenRate(numbers *[]string, carbon bool) int64 {
	var origList []string
	var onesList, zeroesList, lookupList []int
	var oneBit, zeroBit, i, numBytes int

	origList = (*numbers)
	numBytes = len(origList[0])
	for index, _ := range origList {
		lookupList = append(lookupList, index)
	}

	for i = 0; ; i = ((i + 1) % numBytes) {
		oneBit, zeroBit = 0, 0
		onesList, zeroesList = nil, nil
		if len(lookupList) == 1 {
			break
		}
		for _, index := range lookupList {
			if string(origList[index][i]) == "1" {
				onesList = append(onesList, index)
				oneBit++
			} else {
				zeroBit++
				zeroesList = append(zeroesList, index)
			}

			if !carbon {
				if oneBit >= zeroBit {
					lookupList = onesList
				} else {
					lookupList = zeroesList
				}
			} else {
				if oneBit < zeroBit {
					lookupList = onesList
				} else {
					lookupList = zeroesList
				}
			}
		}
	}

	res, err := strconv.ParseInt(origList[lookupList[0]], 2, 64)
	check(err)

	return res
}

func SolvePart2(numbers *[]string) int {
	oxygenRate := getCarbonOxygenRate(numbers, false)
	carbonRate := getCarbonOxygenRate(numbers, true)

	lifeSupportRate := oxygenRate * carbonRate
	return int(lifeSupportRate)
}

func main() {
	part1 := true
	var ans int
	var numbers []string
	arg := os.Args
	if len(arg) == 3 && arg[2] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num := scanner.Text()
		numbers = append(numbers, num)
	}
	if part1 {
		ans = SolvePart1(&numbers)
	} else {
		ans = SolvePart2(&numbers)
	}

	fmt.Println("Answer is ", ans)
}
