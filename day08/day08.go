package day08

import (
	"fmt"
	"strconv"
	"syodage/aoc2022/utils"
)

var (
	FinalIput  = utils.ReadLines("inputs/day08_input.txt")
	LocalInput = []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
)

type tree struct {
	h            int
	leftVisible  bool
	rightVisible bool
	upVisible    bool
	downVisible  bool
}

func (t *tree) print() {
	fmt.Printf("-%t-", t.leftVisible)
}

func FirstPart(inputs []string) int {
	grid := buildGrid(inputs)

	// left to right
	var currMax int
	for _, tr := range grid {
		for j, tc := range tr {
			if j == 0 {
				tc.leftVisible = true
				currMax = tc.h
				continue
			}

			if currMax < tc.h {
				tc.leftVisible = true
				currMax = tc.h
			}
		}
	}

	// // right to left
	currMax = 0
	for _, tr := range grid {
		x := len(tr)
		for j := range tr {
			tc := tr[x-(j+1)]
			if j == 0 {
				tc.rightVisible = true
				currMax = tc.h
				continue
			}

			if currMax < tc.h {
				tc.rightVisible = true
				currMax = tc.h
			}
		}
	}

	// up
	currMax = 0
	y := len(grid)
	for i, tr := range grid {
		for j := range tr {
			tc := grid[j][i]
			if j == 0 {
				tc.upVisible = true
				currMax = tc.h
				continue
			}

			if currMax < tc.h {
				tc.upVisible = true
				currMax = tc.h
			}
		}
	}

	// down
	currMax = 0
	y = len(grid)
	for i, tr := range grid {
		for j := range tr {
			tc := grid[y-j-1][i]
			if j == 0 {
				tc.downVisible = true
				currMax = tc.h
				continue
			}

			if currMax < tc.h {
				tc.downVisible = true
				currMax = tc.h
			}
		}
	}

	var visibleCount int
	for _, tr := range grid {
		// fmt.Println()
		for _, tc := range tr {
			// tc.print()
			if tc.leftVisible || tc.rightVisible || tc.upVisible || tc.downVisible {
				visibleCount++
			}
		}
	}
	return visibleCount
}

func SecondPart(inputs []string) int {

	grid := buildGrid(inputs)

	var maxScenicScore int

	for i, tr := range grid {
		for j, tc := range tr {

			var ls, rs, us, ds int

			// left
			for x := j - 1; x >= 0; x-- {
				c := grid[i][x]
				if c.h < tc.h {
					ls++
				} else {
					ls++
					break
				}
			}

			//right
			for x := j + 1; x < len(tr); x++ {
				c := grid[i][x]
				if c.h < tc.h {
					rs++
				} else {
					rs++
					break
				}
			}

			//  up
			for y := i - 1; y >= 0; y-- {
				c := grid[y][j]
				if c.h < tc.h {
					us++
				} else {
					us++
					break
				}
			}

			//  down
			for y := i + 1; y < len(grid); y++ {
				c := grid[y][j]
				if c.h < tc.h {
					ds++
				} else {
					ds++
					break
				}
			}

			score := ls * rs * us * ds
			fmt.Printf("-%d-", score)
			if maxScenicScore < score {
				maxScenicScore = score
			}
		}
		fmt.Println()
	}
	return maxScenicScore
}

func buildGrid(inputs []string) [][]*tree {
	var grid [][]*tree
	for _, r := range inputs {
		var row []*tree
		for _, c := range r {
			a, _ := strconv.Atoi(string(c))
			row = append(row, &tree{h: a})
		}
		grid = append(grid, row)
	}
	return grid
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
