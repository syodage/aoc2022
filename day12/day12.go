package day12

import (
	"fmt"
	"syodage/aoc2022/utils"
)

// import "syodage/aoc2022/utils"

var (
	LocalInput = []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	FinalInput = utils.ReadLines("inputs/day12_input.txt")
)

type cell struct {
	x, y    int
	steps   int
	visited bool
}

func FirstPart(inputs []string) int {
	path := [][]*cell{}
	grid := toGrid(inputs)

	leftB, rightB, upB, downB := -1, len(grid[0]), -1, len(grid)
	utils.Printlnf("leftB: %d, rightB: %d, upB: %d, downB: %d", leftB, rightB, upB, downB)
	var sx, sy, ex, ey int
	max := len(grid)*len(grid[0]) + 1
	fmt.Printf("Max value is: %d\n", max)
	for i, g := range grid {
		var p []*cell
		for j, r := range g {
			if r == 'S' {
				sx = j
				sy = i
				utils.Printlnf("Start (%d, %d)", sx, sy)
				p = append(p, &cell{x: j, y: i, steps: 0, visited: true})
				// replaced starting point
				grid[sy][sx] = 'a'
				continue
			}

			if r == 'E' {
				ex = j
				ey = i
				utils.Printlnf("End (%d, %d)", ex, ey)
				grid[ey][ex] = 'z'
			}
			p = append(p, &cell{x: j, y: i, steps: max})
		}
		path = append(path, p)
	}

	// printGrid(grid)

	isValid := func(x, y int) bool {
		return leftB < x && x < rightB && upB < y && y < downB
	}

	bfsList := []*cell{path[sy][sx]}

	process := func(cx, cy, nx, ny int) {
		cc := path[cy][cx]
		nc := path[ny][nx]

		ch := grid[cy][cx]
		nh := grid[ny][nx]

		if nc.visited {
			return
		}

		// if to much higher than current ignore
		diff := nh - ch
		if 1 < diff {
			return
		}

		nc.visited = true
		nc.steps = cc.steps + 1
		// utils.Printlnf("Adding cell (x: %d, y: %d, steps: %d)", nx, ny, nc.steps)
		bfsList = append(bfsList, nc)
	}

	for len(bfsList) > 0 {
		// printPathCord(path)
		currCell := bfsList[0]
		bfsList = bfsList[1:]

		cx, cy := currCell.x, currCell.y
		// utils.Printlnf("Processing  cell (x: %d, y: %d, steps: %d, rune; %c)", cx, cy, currCell.steps, grid[cy][cx])

		if cx == ex && cy == ey {
			return currCell.steps
		}

		// left
		nx, ny := cx-1, cy
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}
		// right
		nx, ny = cx+1, cy
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}

		// up
		nx, ny = cx, cy-1
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}

		// down
		nx, ny = cx, cy+1
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}
	}
	return -1
}

func SecondPart(inputs []string) int {
	path := [][]*cell{}
	grid := toGrid(inputs)

	var bfsList []*cell

	leftB, rightB, upB, downB := -1, len(grid[0]), -1, len(grid)
	utils.Printlnf("leftB: %d, rightB: %d, upB: %d, downB: %d", leftB, rightB, upB, downB)
	var ex, ey int
	max := len(grid)*len(grid[0]) + 1
	fmt.Printf("Max value is: %d\n", max)
	for i, g := range grid {
		var p []*cell
		for j, r := range g {

			c := &cell{x: j, y: i, steps: max}
			x, y := j, i
			if r == 'a' {
				utils.Printlnf("Start (%d, %d)", x, y)
				c.steps = 0
				c.visited = true
				bfsList = append(bfsList, c)
			}

			if r == 'S' {
				utils.Printlnf("Start (%d, %d)", x, y)
				c.steps = 0
				c.visited = true
				// replaced starting point
				grid[y][x] = 'a'
				bfsList = append(bfsList, c)
			}

			if r == 'E' {
				ex = j
				ey = i
				utils.Printlnf("End (%d, %d)", ex, ey)
				grid[ey][ex] = 'z'
			}
			p = append(p, c)
		}
		path = append(path, p)
	}

	// printGrid(grid)

	isValid := func(x, y int) bool {
		return leftB < x && x < rightB && upB < y && y < downB
	}

	process := func(cx, cy, nx, ny int) {
		cc := path[cy][cx]
		nc := path[ny][nx]

		ch := grid[cy][cx]
		nh := grid[ny][nx]

		if nc.visited {
			return
		}

		// if to much higher than current ignore
		diff := nh - ch
		if 1 < diff {
			return
		}

		nc.visited = true
		nc.steps = cc.steps + 1
		// utils.Printlnf("Adding cell (x: %d, y: %d, steps: %d)", nx, ny, nc.steps)
		bfsList = append(bfsList, nc)
	}

	dd := 0
	for len(bfsList) > 0 {
		dd--
		if dd > 0 {
			printPathCord(path)
		}

		currCell := bfsList[0]
		bfsList = bfsList[1:]

		cx, cy := currCell.x, currCell.y
		// utils.Printlnf("Processing  cell (x: %d, y: %d, steps: %d, rune; %c)", cx, cy, currCell.steps, grid[cy][cx])

		if cx == ex && cy == ey {
			return currCell.steps
		}

		// left
		nx, ny := cx-1, cy
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}
		// right
		nx, ny = cx+1, cy
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}

		// up
		nx, ny = cx, cy-1
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}

		// down
		nx, ny = cx, cy+1
		if isValid(nx, ny) {
			process(cx, cy, nx, ny)
		}
	}

	// return path[ey][ex].steps
	return -1
}

func printPath(paths [][]*cell) {
	for _, row := range paths {
		for _, c := range row {
			fmt.Printf("-%d", c.steps)
		}
		fmt.Println()
	}

	fmt.Println("==================================")

}

func printPathCord(paths [][]*cell) {
	for _, row := range paths {
		for _, c := range row {
			fmt.Printf("-(%d, %d, %d)", c.x, c.y, c.steps)
		}
		fmt.Println()
	}

	fmt.Println("==================================")

}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("-%c", c)
		}
		fmt.Println()
	}

	fmt.Println("==================================")

}

func toGrid(inputs []string) [][]rune {
	var res [][]rune
	for _, input := range inputs {
		var row []rune
		for _, r := range input {
			row = append(row, r)

		}
		res = append(res, row)
	}
	return res
}
