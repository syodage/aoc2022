package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello AoC++")
	// fmt.Printf("Day01_01: Max Calary sum: %d\n", Day01_FirstAnswer())
	// fmt.Printf("Day01_02: Sum of top 3 max Calary sums: %d\n", Day01_SecondAnswer())

	// day02Input := ReadLines("inputs/day02_input.txt")
	// fmt.Printf("Day02_01: Total Score: %d\n", Day02_FirstAnswer(day02Input))
	// day02Input := []string{"A Y", "B X", "C Z"}
	// fmt.Printf("Day02_02: Total Score: %d\n", Day02_SecondAnswer(day02Input))

	// day03Input := ReadLines("inputs/day03_input.txt")
	// day03Input := []string{
	//   "vJrwpWtwJgWrhcsFMMfFFhFp",
	//   "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	//   "PmmdzqPrVvPwwTWBwg",
	//   "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	//   "ttgJtRGJQctTZtZT",
	//   "CrZsJsPPZsGzwwsLwLmpwMDw",
	// }

	// fmt.Printf("Day03_01: Total Score: %d\n", Day03_FirstAnswer(day03Input))
	// fmt.Printf("Day03_02: Total Score: %d\n", Day03_SecondAnswer(day03Input))

	// day04Input := ReadLines("inputs/day04_input.txt")
	//   day04Input := []string {
	// "2-4,6-8",
	// "2-3,4-5",
	// "5-7,7-9",
	// "2-8,3-7",
	// "6-6,4-6",
	// "2-6,4-8",
	//   }
	// fmt.Printf("Day04_01: Total Score: %d\n", Day04_FirstAnswer(day04Input))
	// fmt.Printf("Day04_02: Total Score: %d\n", Day04_SecondAnswer(day04Input))

	// fmt.Println("Start")
	// for _, s := range ReadHead("inputs/day05_input.txt", 10) {
	//   fmt.Printf("%s\n", s)
	// }
	// fmt.Println("End")

	//   day05Input := []string {
	// "move 1 from 2 to 1",
	// "move 3 from 1 to 3",
	// "move 2 from 2 to 1",
	// "move 1 from 1 to 2",
	//   }
	//    day05_Crates := map[int][]rune {
	//      1: {'Z', 'N'},
	//      2: {'M', 'C', 'D'},
	//      3: {'P'},
	//    }

	// [D]                     [N] [F]
	// [H] [F]             [L] [J] [H]
	// [R] [H]             [F] [V] [G] [H]
	// [Z] [Q]         [Z] [W] [L] [J] [B]
	// [S] [W] [H]     [B] [H] [D] [C] [M]
	// [P] [R] [S] [G] [J] [J] [W] [Z] [V]
	// [W] [B] [V] [F] [G] [T] [T] [T] [P]
	// [Q] [V] [C] [H] [P] [Q] [Z] [D] [W]
	//  1   2   3   4   5   6   7   8   9
	// day05_Crates := map[int][]rune {
	//   1: {'Q', 'W', 'P', 'S', 'Z', 'R', 'H', 'D'},
	//   2: {'V', 'B', 'R', 'W', 'Q', 'H', 'F'},
	//   3: {'C', 'V', 'S', 'H'},
	//   4: {'H', 'F', 'G'},
	//   5: {'P', 'G', 'J', 'B', 'Z'},
	//   6: {'Q', 'T', 'J', 'H', 'W', 'F', 'L'},
	//   7: {'Z', 'T', 'W', 'D', 'L', 'V', 'J', 'N'},
	//   8: {'D', 'T', 'Z', 'C', 'J', 'G', 'H', 'F'},
	//   9: {'W', 'P', 'V', 'M', 'B', 'H'},
	// }
	// day05Input := ReadLinesWithIgnoreHead("inputs/day05_input.txt", 10)
	// fmt.Printf("line 1: %s\n", day05Input[0])
	// fmt.Printf("Day05_01: Total Score: %s\n", Day05_FirstAnswer(day05Input, day05_Crates))
	// fmt.Printf("Day05_02: Total Score: %s\n", Day05_SecondAnswer(day05Input, day05_Crates))

	// day06Input := ReadLines("inputs/day06_input.txt")
	// day06Input := []string{
	// 	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",    // 7
	// 	"bvwbjplbgvbhsrlpgdmjqwftvncz",      // 5
	// 	"nppdvjthqldpwncqszvftbrmjlhg",      // 6
	// 	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", // 10
	// 	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",  // 11
	// }
	// fmt.Printf("Day06_01: Data need to processed: %v\n", Day06_FirstAnswer(day06Input))
	// fmt.Printf("Day06_01: Data need to processed: %v\n", Day06_SecondAnswer(day06Input))

	// day07Input := ReadLines("inputs/day07_input.txt")
	// day07Input := []string{
	// 	"$ cd /",
	// 	"$ ls",
	// 	"dir a",
	// 	"14848514 b.txt",
	// 	"8504156 c.dat",
	// 	"dir d",
	// 	"$ cd a",
	// 	"$ ls",
	// 	"dir e",
	// 	"29116 f",
	// 	"2557 g",
	// 	"62596 h.lst",
	// 	"$ cd e",
	// 	"$ ls",
	// 	"584 i",
	// 	"$ cd ..",
	// 	"$ cd ..",
	// 	"$ cd d",
	// 	"$ ls",
	// 	"4060174 j",
	// 	"8033020 d.log",
	// 	"5626152 d.ext",
	// 	"7214296 k",
	// }
	// fmt.Printf("Day07_01: Data need to processed: %v\n", Day07_FirstAnswer(day07Input, 100000))
	// fmt.Printf("Day07_02: delete dir with size: %d\n", Day07_SecondAnswer(day07Input))

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

	day08Input := ReadLines("inputs/day08_input.txt")
	// day08Input := []string{
	// 	"30373",
	// 	"25512",
	// 	"65332",
	// 	"33549",
	// 	"35390",
	// }

	// fmt.Printf("\nDay08_01: visible trees: %d\n", Day08_FirstAnswer(day08Input))
	fmt.Printf("\nDay08_02: heighest scenic score: %d\n", Day08_SecondAnswer(day08Input))

}
