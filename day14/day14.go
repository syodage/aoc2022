package day14

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

import "syodage/aoc2022/utils"

var (
	LocalInput = []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}

	FinalInput = utils.ReadLines("inputs/day14_input.txt")
)

type point struct {
	x, y     int
	isRock   bool
	isFilled bool
}

type path struct {
	points []*point
}

func FirstPart(inputs []string) int {
	// s := &point{500, 0, false}
	pathList := parse(inputs)

	left, right, up, down := findDim(pathList)
	// give some room
	left--
	right++
	down++
	fmt.Printf("left: %d, right: %d, up: %d, down: %d\n", left, right, up, down)

	cave := make(map[int]map[int]*point)
	for i := 0; i <= down; i++ {
		raw := make(map[int]*point)
		for j := left; j <= right; j++ {
			raw[j] = &point{x: j, y: i}
		}
		cave[i] = raw
	}
	// printCave(cave, true)
	// fmt.Println("")
	addRock(cave, pathList)
	// printCave(cave, false)

	notAbyss := true
	var unitOfSands int
	currX, currY := 500, 0
	for notAbyss {
		// fmt.Printf("Curr: (%d, %d)\n", currX, currY)
		if currX <= left || currX >= right || currY >= down {
			notAbyss = false
			continue
		}

		// can go straight down
		var p *point
		// fmt.Printf("Prcessing: (%d, %d)\n", currX, currY)
		p = cave[currY+1][currX]
		if !p.isFilled {
			currY = currY + 1
			continue
		}

		p = cave[currY+1][currX-1]
		if !p.isFilled {
			currX = currX - 1
			currY = currY + 1
			continue
		}

		p = cave[currY+1][currX+1]
		if !p.isFilled {
			currX = currX + 1
			currY = currY + 1
			continue
		}

		c := cave[currY][currX]
		c.isFilled = true
		unitOfSands++
		currX = 500
		currY = 0
	}

	// printCave(cave, false)
	return unitOfSands
}

func SecondPart(inputs []string, margin int) int {
	pathList := parse(inputs)

	left, right, up, down := findDim(pathList)
	left -= margin
	right += margin
	down += 2
	fmt.Printf("left: %d, right: %d, up: %d, down: %d\n", left, right, up, down)

	// add the floor line
	pathList = append(pathList, &path{points: []*point{{x: left, y: down}, {x: right, y: down}}})

	cave := make(map[int]map[int]*point)
	for i := 0; i <= down; i++ {
		raw := make(map[int]*point)
		for j := left; j <= right; j++ {
			raw[j] = &point{x: j, y: i}
		}
		cave[i] = raw
	}
	// printCave(cave, true)
	// fmt.Println("")
	addRock(cave, pathList)
	// printCave(cave, false)

	var unitOfSands int
	currX, currY := 500, 0
	s := cave[0][500]
	for !s.isFilled {
		// can go straight down
		var p *point
		// fmt.Printf("Prcessing: (%d, %d)\n", currX, currY)
		p = cave[currY+1][currX]
		if !p.isFilled {
			currY = currY + 1
			continue
		}

		p = cave[currY+1][currX-1]
		if !p.isFilled {
			currX = currX - 1
			currY = currY + 1
			continue
		}

		p = cave[currY+1][currX+1]
		if !p.isFilled {
			currX = currX + 1
			currY = currY + 1
			continue
		}

		c := cave[currY][currX]
		c.isFilled = true
		unitOfSands++
		currX = 500
		currY = 0
	}

	// printCave(cave, false)
	return unitOfSands
}

func addRock(cave map[int]map[int]*point, pathList []*path) {

	for _, path := range pathList {
		if len(path.points) == 0 {
			continue
		}

		if len(path.points) == 1 {
			p := path.points[0]
			cave[p.y][p.x].isRock = true
			cave[p.y][p.x].isFilled = true
		}

		for i := 1; i < len(path.points); i++ {
			a := path.points[i-1]
			b := path.points[i]
			// fmt.Printf("from: (%d, %d) -> (%d, %d)\n", a.x, a.y, b.x, b.y)

			if a.x == b.x {
				if a.y > b.y {
					c := a
					a = b
					b = c
				}
				for j := a.y; j <= b.y; j++ {
					cave[j][a.x].isRock = true
					cave[j][a.x].isFilled = true

				}
			}

			if a.y == b.y {
				if a.x > b.x {
					c := a
					a = b
					b = c
				}

				// fmt.Printf("from: (%d, %d) -> (%d, %d)\n", a.x, a.y, b.x, b.y)
				for j := a.x; j <= b.x; j++ {
					cave[a.y][j].isRock = true
					cave[a.y][j].isFilled = true
					// fmt.Printf("Filling X axis: (%d, %d)\n", j, a.y)
				}
			}
		}
	}
}

func findDim(pl []*path) (int, int, int, int) {
	var left, right, up, down int

	for _, path := range pl {
		for _, p := range path.points {
			if left == 0 {
				left = p.x
				right = p.x
				up = p.y
				down = p.y
			}

			if left > p.x {
				left = p.x
			}

			if right < p.x {
				right = p.x
			}

			if up > p.y {
				up = p.y
			}

			if down < p.y {
				down = p.y
			}
		}
	}

	return left, right, up, down
}

func parse(inputs []string) []*path {
	var res []*path

	for _, input := range inputs {
		strP := strings.Split(input, "->")

		var points []*point
		for _, sp := range strP {
			sp = strings.Trim(sp, " ")
			cord := strings.Split(sp, ",")
			x, _ := strconv.Atoi(cord[0])
			y, _ := strconv.Atoi(cord[1])
			points = append(points, &point{x, y, false, false})
		}

		res = append(res, &path{points: points})
	}

	return res
}

func printCave(cave map[int]map[int]*point, verbos bool) {
	var keys []int
	for k := range cave {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d :", k)
		raw := cave[k]

		var innerKeys []int
		for k := range raw {
			innerKeys = append(innerKeys, k)
		}
		sort.Ints(innerKeys)
		for _, ik := range innerKeys {
			p := raw[ik]
			var f rune
			if p.isRock {
				f = '#'
			} else if p.isFilled {
				f = 'o'
			} else if k == 0 && ik == 500 {
				f = '+'
			} else {
				f = '.'
			}

			if verbos {
				fmt.Printf(" (%d, %d, %c) ", p.x, p.y, f)
			} else {

				fmt.Printf(" %c ", f)
			}
		}
		fmt.Println("")
	}
}
