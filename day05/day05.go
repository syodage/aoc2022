package day05

import (
	"fmt"
	"strconv"
	"strings"
	"syodage/aoc2022/utils"
)

var (
	LocalInput = []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	LocalCrates = map[int][]rune{
		1: {'Z', 'N'},
		2: {'M', 'C', 'D'},
		3: {'P'},
	}

	// [D]                     [N] [F]
	// [H] [F]             [L] [J] [H]
	// [R] [H]             [F] [V] [G] [H]
	// [Z] [Q]         [Z] [W] [L] [J] [B]
	// [S] [W] [H]     [B] [H] [D] [C] [M]
	// [P] [R] [S] [G] [J] [J] [W] [Z] [V]
	// [W] [B] [V] [F] [G] [T] [T] [T] [P]
	// [Q] [V] [C] [H] [P] [Q] [Z] [D] [W]
	//  1   2   3   4   5   6   7   8   9
	FinalCrates = map[int][]rune{
		1: {'Q', 'W', 'P', 'S', 'Z', 'R', 'H', 'D'},
		2: {'V', 'B', 'R', 'W', 'Q', 'H', 'F'},
		3: {'C', 'V', 'S', 'H'},
		4: {'H', 'F', 'G'},
		5: {'P', 'G', 'J', 'B', 'Z'},
		6: {'Q', 'T', 'J', 'H', 'W', 'F', 'L'},
		7: {'Z', 'T', 'W', 'D', 'L', 'V', 'J', 'N'},
		8: {'D', 'T', 'Z', 'C', 'J', 'G', 'H', 'F'},
		9: {'W', 'P', 'V', 'M', 'B', 'H'},
	}
	FinalInput = utils.ReadLinesWithIgnoreHead("inputs/day05_input.txt", 10)
)

func FirstAnswer(inputs []string, crates map[int][]rune) string {
	fmt.Printf("Map: %+v\n", crates)
	// input is in format of "move x crates from y to z"
	for _, input := range inputs {
		words := strings.Split(input, " ")
		x, _ := strconv.Atoi(words[1])
		y, _ := strconv.Atoi(words[3])
		z, _ := strconv.Atoi(words[5])

		// fmt.Printf("Move %d from %d to %d\n", x, y, z)

		for i := 0; i < x; i++ {
			lenY := len(crates[y])
			crates[z] = append(crates[z], crates[y][lenY-1])
			crates[y] = crates[y][:lenY-1]
		}

		// fmt.Printf("After Move: %+v\n", crates)
	}

	var res string
	for j := 1; j <= len(crates); j++ {
		v := crates[j]
		res += string(v[len(v)-1])
	}
	return res
}

func SecondAnswer(inputs []string, crates map[int][]rune) string {
	fmt.Printf("Map: %+v\n", crates)
	for _, input := range inputs {
		words := strings.Split(input, " ")
		x, _ := strconv.Atoi(words[1])
		y, _ := strconv.Atoi(words[3])
		z, _ := strconv.Atoi(words[5])

		// fmt.Printf("Move %d from %d to %d\n", x, y, z)

		// input is in format of "move x crates from y to z"
		lenY := len(crates[y])
		crates[z] = append(crates[z], crates[y][lenY-x:]...)
		crates[y] = crates[y][:lenY-x]

		// fmt.Printf("After Move: %+v\n", crates)
	}

	var res string
	for j := 1; j <= len(crates); j++ {
		v := crates[j]
		res += string(v[len(v)-1])
	}
	return res
}
