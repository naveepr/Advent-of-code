package main

import (
	"bufio"
	"os"
	"testing"
)

func TestPuzzle(t *testing.T) {
	f, err := os.CreateTemp("./", "sample")
	check(err)

	defer os.Remove(f.Name())

	_, err = f.WriteString("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	check(err)

	f, err = os.Open(f.Name())
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	t.Run("TestPuzzle_1", func(t *testing.T) {
		ans := SolvePart1(scanner)
		if ans != 7 {
			t.Errorf("TestPuzzle_1 returned %d expected %d", ans, 8)
		}
	})

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	t.Run("TestPuzzle_2", func(t *testing.T) {
		ans := SolvePart2(scanner)
		if ans != 5 {
			t.Errorf("TestPuzzle_2 returned %d expected %d", ans, 5)
		}
	})
}

func TestInput1(t *testing.T) {
	f, err := os.Open("input.txt")
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	if ans := SolvePart1(scanner); ans != 1167 {
		t.Errorf("TestInput1 returned %d expected %d", ans, 1167)

	}
}

func TestInput2(t *testing.T) {
	f, err := os.Open("input.txt")
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	if ans := SolvePart2(scanner); ans != 1130 {
		t.Errorf("TestInput2 returned %d expected %d", ans, 1130)

	}
}
