package day10

import (
	"fmt"
	"strconv"
	"strings"
	"syodage/aoc2022/utils"
)

var (
	SmallInput = []string{
		"noop",
		"addx 3",
		"addx -5",
	}

	BigInput = []string{
		"addx 15",
		"addx -11",
		"addx 6",
		"addx -3",
		"addx 5",
		"addx -1",
		"addx -8",
		"addx 13",
		"addx 4",
		"noop",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx -35",
		"addx 1",
		"addx 24",
		"addx -19",
		"addx 1",
		"addx 16",
		"addx -11",
		"noop",
		"noop",
		"addx 21",
		"addx -15",
		"noop",
		"noop",
		"addx -3",
		"addx 9",
		"addx 1",
		"addx -3",
		"addx 8",
		"addx 1",
		"addx 5",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx -36",
		"noop",
		"addx 1",
		"addx 7",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"addx 6",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx 7",
		"addx 1",
		"noop",
		"addx -13",
		"addx 13",
		"addx 7",
		"noop",
		"addx 1",
		"addx -33",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"noop",
		"noop",
		"noop",
		"addx 8",
		"noop",
		"addx -1",
		"addx 2",
		"addx 1",
		"noop",
		"addx 17",
		"addx -9",
		"addx 1",
		"addx 1",
		"addx -3",
		"addx 11",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx -13",
		"addx -19",
		"addx 1",
		"addx 3",
		"addx 26",
		"addx -30",
		"addx 12",
		"addx -1",
		"addx 3",
		"addx 1",
		"noop",
		"noop",
		"noop",
		"addx -9",
		"addx 18",
		"addx 1",
		"addx 2",
		"noop",
		"noop",
		"addx 9",
		"noop",
		"noop",
		"noop",
		"addx -1",
		"addx 2",
		"addx -37",
		"addx 1",
		"addx 3",
		"noop",
		"addx 15",
		"addx -21",
		"addx 22",
		"addx -6",
		"addx 1",
		"noop",
		"addx 2",
		"addx 1",
		"noop",
		"addx -10",
		"noop",
		"noop",
		"addx 20",
		"addx 1",
		"addx 2",
		"addx 2",
		"addx -6",
		"addx -11",
		"noop",
		"noop",
		"noop",
	}

	FinalInput = utils.ReadLines("inputs/day10_input.txt")
)

type instruction struct {
	cycles int
	v      int
}

func FirstAnswer(inputs []string) int {
	var in []*instruction
	for _, input := range inputs {
		if input == "noop" {
			in = append(in, &instruction{cycles: 1})
		} else {
			v, _ := strconv.Atoi(strings.Split(input, " ")[1])
			in = append(in, &instruction{cycles: 2, v: v})
		}
	}

	var steps int
	var ss int
	X := 1
	for {
		if in == nil || len(in) == 0 {
			break
		}

		steps++
		// utils.Printlnf("Step: %d", steps)
		if multiply(steps) {
			utils.Printlnf("Multiply %d * %d", X, steps)
			ss += X * steps
		}
		if in[0].cycles == 1 {
			// utils.Printlnf("Adding %d, X is: %d", in[0].v, X)
			X += in[0].v
			in = in[1:]
		} else {
			in[0].cycles--
		}
	}
	return ss
}

func multiply(a int) bool {
	if a < 20 {
		return false
	}

	return a == 20 || (a-20)%40 == 0
}

func SecondAnswer(inputs []string) int {

	var in []*instruction
	for _, input := range inputs {
		if input == "noop" {
			in = append(in, &instruction{cycles: 1})
		} else {
			v, _ := strconv.Atoi(strings.Split(input, " ")[1])
			in = append(in, &instruction{cycles: 2, v: v})
		}
	}

	pos := 1
	sprite := 1
	// fmt.Print("#")
	for {
		if in == nil || len(in) == 0 {
			break
		}
		if sprite <= pos && pos <= sprite+2 {
			fmt.Print("#")
			utils.Printlnf("Sprite %d, pos is: %d", sprite, pos)
		} else {
			fmt.Print(".")
			utils.Printlnf("Sprite %d, pos is: %d", sprite, pos)
		}

		if in[0].cycles == 1 {
			sprite += in[0].v
			in = in[1:]
			utils.Printlnf("Sprite %d, pos is: %d", sprite, pos)
		} else {
			in[0].cycles--
		}

		if (pos % 40) == 0 {
			fmt.Println()
			pos = 0
		}
		pos++
	}
	return 0
}
