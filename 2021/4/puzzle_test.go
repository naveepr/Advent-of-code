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

	_, err = f.WriteString("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")
	check(err)

	f, err = os.Open(f.Name())
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	t.Run("TestPuzzle_1", func(t *testing.T) {
		ans := SolvePart1(scanner)
		if ans != 150 {
			t.Errorf("TestPuzzle_1 returned %d expected %d", ans, 150)
		}
	})

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	t.Run("TestPuzzle_2", func(t *testing.T) {
		ans := SolvePart2(scanner)
		if ans != 900 {
			t.Errorf("TestPuzzle_2 returned %d expected %d", ans, 900)
		}
	})
}

func TestInput1(t *testing.T) {
	f, err := os.Open("input.txt")
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	if ans := SolvePart1(scanner); ans != 1714680 {
		t.Errorf("TestInput1 returned %d expected %d", ans, 1714680)

	}
}

func TestInput2(t *testing.T) {
	f, err := os.Open("input.txt")
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	if ans := SolvePart2(scanner); ans != 1963088820 {
		t.Errorf("TestInput2 returned %d expected %d", ans, 1963088820)

	}
}
