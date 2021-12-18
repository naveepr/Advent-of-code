package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// An Item is something we manage in a priority queue.
type Item struct {
	x    int // The priority of the item in the queue.
	y    int
	risk int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].risk < pq[j].risk
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) [][]int {
	var matrix [][]int
	for scanner.Scan() {
		s := scanner.Text()
		str := strings.Split(s, "")
		row := []int{}
		for _, s := range str {
			val, err := strconv.Atoi(s)
			check(err)
			row = append(row, val)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

type pair struct {
	x, y int
}

func getRC(matrix [][]int, row, col, rowsz, colsz int) int {
	r, c := (row % rowsz), (col % colsz)

	fmt.Println(rowsz, colsz, row, col, r, c)
	x := matrix[r][c] + row/rowsz + col/colsz
	return ((x-1)%9 + 1)
}

func Solve(matrix [][]int, part1 bool) int {
	var res, tempr, tempc int
	var visited [][]bool
	rowsz, colsz := len(matrix), len(matrix[0])

	tempr, tempc = rowsz, colsz

	if !part1 {
		rowsz *= 5
		colsz *= 5
	}
	cost := make([][]int, rowsz)

	res = math.MaxInt
	visited = make([][]bool, rowsz)
	for i := range visited {
		visited[i] = make([]bool, colsz)
		cost[i] = make([]int, colsz)
	}

	for i := 0; i < rowsz; i++ {
		for j := 0; j < colsz; j++ {
			cost[i][j] = math.MaxInt
		}
	}

	pq := make(PriorityQueue, 1)
	pq[0] = Item{0, 0, 0}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(Item)
		if visited[item.x][item.y] {
			continue
		}

		c := item.risk

		if c < cost[item.x][item.y] {
			cost[item.x][item.y] = c
		}
		if item.x == rowsz-1 && item.y == colsz-1 {
			res = c
			break
		}

		points := [4]pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, p := range points {
			newx := item.x + p.x
			newy := item.y + p.y
			if newx < 0 || newy < 0 || newx == rowsz || newy == colsz || visited[newx][newy] {
				continue
			}
			tempItem := Item{
				x:    newx,
				y:    newy,
				risk: (getRC(matrix, newx, newy, tempr, tempc) + c),
			}
			heap.Push(&pq, tempItem)
		}

		visited[item.x][item.y] = true
	}

	return res
}

func main() {
	arg := os.Args
	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	matrix := parseInput(scanner)
	ans1 := Solve(matrix, true)
	ans2 := Solve(matrix, false)

	fmt.Println("Answer is ", ans1, ans2)
}
