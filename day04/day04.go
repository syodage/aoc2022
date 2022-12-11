package day04

import (
	"strconv"
	"strings"
	"syodage/aoc2022/utils"
)

var (
	LocalInput = []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	FinalInput = utils.ReadLines("inputs/day04_input.txt")
)

func FirstPart(inputs []string) int {

	var totalOverlap int
	for _, input := range inputs {
		works := strings.Split(input, ",")
		first := strings.Split(works[0], "-")
		second := strings.Split(works[1], "-")

		start1, _ := strconv.Atoi(first[0])
		end1, _ := strconv.Atoi(first[1])

		start2, _ := strconv.Atoi(second[0])
		end2, _ := strconv.Atoi(second[1])

		// does second include first
		if start2 <= start1 && end1 <= end2 {
			totalOverlap++
			continue
		}

		// does first include second
		if start1 <= start2 && end2 <= end1 {
			totalOverlap++
		}
	}
	return totalOverlap
}

func SecondPart(inputs []string) int {
	var totalOverlap int
	for _, input := range inputs {
		works := strings.Split(input, ",")
		first := strings.Split(works[0], "-")
		second := strings.Split(works[1], "-")

		start1, _ := strconv.Atoi(first[0])
		end1, _ := strconv.Atoi(first[1])

		start2, _ := strconv.Atoi(second[0])
		end2, _ := strconv.Atoi(second[1])

		// first pair fully overlap with second
		if start2 <= start1 && end1 <= end2 {
			totalOverlap++
			continue
		}

		// second pair fully overalp with first
		if start1 <= start2 && end2 <= end1 {
			totalOverlap++
			continue
		}

		// last part of first pair overlap with first part of second pair

		if (start1 < start2 && end1 < start2) || (end2 < start1 && end2 < end1) {
			continue
		}

		totalOverlap++

	}
	return totalOverlap
}
