package main

import (
	"fmt"
	"syodage/aoc2022/day13"
	"syodage/aoc2022/utils"
)

func main() {
	// fmt.Println("Hello AoC++")
	// fmt.Printf("Day01_01: Max Calary sum: %d\n", day01.FirstAnswer())
	// fmt.Printf("Day01_02: Sum of top 3 max Calary sums: %d\n", day01.SecondAnswer())

	// fmt.Printf("Day02_01: Total Score: %d\n", day02.FirstPart(day02.FinalInput))
	// fmt.Printf("Day02_02: Total Score: %d\n", day02.SecondPart(day02.FinalInput))

	// fmt.Printf("Day03_01: Total Score: %d\n", day03.FirstPart(day03.FinalInput))
	// fmt.Printf("Day03_02: Total Score: %d\n", day03.SecondPart(day03.FinalInput))

	// fmt.Printf("Day04_01: Total Score: %d\n", day04.FirstPart(day04.FinalInput))
	// fmt.Printf("Day04_02: Total Score: %d\n", day04.SecondPart(day04.FinalInput))

	// fmt.Printf("Day05_01: Total Score: %s\n", day05.FirstAnswer(day05.FinalInput, day05.FinalCrates))
	// fmt.Printf("Day05_02: Total Score: %s\n", day05.SecondAnswer(day05.FinalInput, day05.FinalCrates))

	// fmt.Printf("Day06_01: Data need to processed: %v\n", day06.FirstPart(day06.FinalInput))
	// fmt.Printf("Day06_01: Data need to processed: %v\n", day06.SecondPart(day06.FinalInput))

	// fmt.Printf("Day07_01: Data need to processed: %v\n", day07.FirstPart(day07.FinalInput, 100000))
	// fmt.Printf("Day07_02: delete dir with size: %d\n", day07.SecondPart(day07.FinalInput))

	// - / (dir)
	// - a (dir)
	//   - e (dir)
	//     - i (file, size=584)
	//   - f (file, size=29116)
	//   - g (file, size=2557)
	//   - h.lst (file, size=62596)
	// - b.txt (file, size=14848514)
	// - c.dat (file, size=8504156)
	// - d (dir)
	//   - j (file, size=4060174)
	//   - d.log (file, size=8033020)
	//   - d.ext (file, size=5626152)
	//   - k (file, size=7214296)

	// fmt.Printf("\nDay08_01: visible trees: %d\n", day08.FirstPart(day08.FinalIput))
	// fmt.Printf("\nDay08_02: heighest scenic score: %d\n", day08.SecondPart(day08.FinalIput))

	// fmt.Printf("Day09_01: visited cells: %d\n", day09.FirstAnswer(day09.FinalInput))
	// fmt.Printf("Day09_02: visited cells: %d\n", day09.SecondAnswer(day09.FinalInput))

	// fmt.Printf("Day10_01: register value: %d\n", day10.FirstAnswer(day10.BigInput))
	// fmt.Println("Day10_02: CRT Screen:")
	// day10.SecondAnswer(day10.FinalInput)

	// fmt.Printf("Day11_01: monkey business: %d\n", day11.FirstPart(day11.LocalInput, 20))
	// fmt.Printf("Day11_01: monkey business: %d\n", day11.FirstPart(day11.FinalInput, 20))
	// fmt.Printf("Day11_02: monkey business with managed wl: %d\n", day11.SecondPart(day11.ManagedLocalInput, 10000))
	// fmt.Printf("Day11_02: monkey business with managed wl: %d\n", day11.SecondPart(day11.ManagedFinalInput, 10000))

	// Day 12
	// fmt.Printf("Day12_01: answer: %d\n", day12.FirstPart(day12.LocalInput))
	// fmt.Printf("Day12_01: answer: %d\n", day12.FirstPart(day12.FinalInput))

	// fmt.Printf("Day12_02: answer: %d\n", day12.SecondPart(day12.LocalInput))
	// fmt.Printf("Day12_02: answer: %d\n", day12.SecondPart(day12.FinalInput))

	utils.Debug = false
	// fmt.Printf("Day13_01: answer: %d\n", day13.FirstPart(day13.FinalInput))
	fmt.Printf("Day13_02: decoded key: %d\n", day13.SecondPart(day13.FinalInput))
}
