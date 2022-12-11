package day03

import "syodage/aoc2022/utils"

var (
	LocalInput = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	FinalInput = utils.ReadLines("inputs/day03_input.txt")
)

func FirstPart(inputs []string) int {

	var totalValue int
	for _, input := range inputs {
		half := len(input) / 2

		m := make(map[byte]int)
		for i := 0; i < half; i++ {
			m[input[i]] = 1
			// fmt.Printf("First hlaf: Char: %c, byte: %d\n", input[i], input[i])
		}

		for j := half; j < len(input); j++ {
			_, ok := m[input[j]]
			if ok {
				var points = int(input[j]) - int('a') + 1
				if points < 0 {
					points = int(input[j]) - int('A') + 27
				}
				// fmt.Printf("Char: %c, byte: %d\n", input[j], points)
				totalValue += points
				break
			}
		}
	}
	return totalValue
}

func SecondPart(inputs []string) int {

	var badges = make(map[rune]int)
	for i := range [26]int{} {
		badges[rune('A'+i)] = 27 + i
		badges[rune('a'+i)] = 1 + i
	}

	var totalScore int
	groups := len(inputs) / 3
	for group := 0; group < groups; group++ {
		m1 := make(map[rune]int)
		m2 := make(map[rune]int)
		m3 := make(map[rune]int)
		for _, c := range inputs[group*3] {
			m1[c] = 1
		}
		for _, c := range inputs[group*3+1] {
			m2[c] = 1
		}
		for _, c := range inputs[group*3+2] {
			m3[c] = 1
		}

		for r, p := range badges {
			_, m1Ok := m1[r]
			_, m2Ok := m2[r]
			_, m3Ok := m3[r]

			if m1Ok && m2Ok && m3Ok {
				// fmt.Printf("Badge: %c, points: %d\n", r, p)
				totalScore += p
			}
		}
	}

	return totalScore
}
