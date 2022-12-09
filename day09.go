package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type knot struct {
	x, y int
	name string
}

func (k *knot) print() {
	Printlnf("(%d, %d)\n", k.x, k.y)
}

func Day09_FirstAnswer(inputs []string) int {

	H, T := &knot{name: "H"}, &knot{name: "T"}
	visited := make(map[string]struct{})
	visited[fmt.Sprintf("%d:%d", T.x, T.y)] = struct{}{}

	for _, input := range inputs {
		sp := strings.Split(input, " ")
		moveDir := sp[0]
		steps, _ := strconv.Atoi(sp[1])

		for i := 1; i <= steps; i++ {
			switch moveDir {
			case "L":
				H.x--
				// Printlnf("H moved Left (%d, %d)", H.x, H.y)
			case "R":
				H.x++
				// Printlnf("H moved Right (%d, %d)", H.x, H.y)
			case "U":
				H.y--
				// Printlnf("H moved UP (%d, %d)", H.x, H.y)
			case "D":
				H.y++
				// Printlnf("H moved Down (%d, %d)", H.x, H.y)
			}
			day09_moveTail(H, T)
			visited[fmt.Sprintf("%d:%d", T.x, T.y)] = struct{}{}
		}
	}

	return len(visited)
}

func day09_moveTail(h, t *knot) {

	mody, modx := math.Abs(float64(h.y-t.y)), math.Abs(float64(h.x-t.x))
	if modx < 2 && mody < 2 {
		// already touching no need to move tail

		// fmt.Printf("T didn't move (%d, %d)\n", t.x, t.y)
		return
	}

	if h.y+2 == t.y {
		t.y--
		t.x = h.x
	} else if h.y-2 == t.y {
		t.y++
		t.x = h.x
	}

	if h.x-2 == t.x {
		t.x++
		t.y = h.y
	} else if h.x+2 == t.x {
		t.x--
		t.y = h.y
	}
	Printlnf("%s moved to (%d, %d)", t.name, t.x, t.y)

}

// moveTail version 2 support moving diagonal
func day09_moveTail2(h, t *knot) {

	mody, modx := math.Abs(float64(h.y-t.y)), math.Abs(float64(h.x-t.x))
	if modx < 2 && mody < 2 {
		// already touching no need to move tail

		// fmt.Printf("T didn't move (%d, %d)\n", t.x, t.y)
		return
	}

	if modx == 2 && mody == 2 {
		if h.y+2 == t.y && h.x+2 == t.x {
			t.y--
			t.x--
		} else if h.y+2 == t.y && h.x-2 == t.x {
			t.y--
			t.x++
		} else if h.y-2 == t.y && h.x+2 == t.x {
			t.y++
			t.x--
		} else if h.y-2 == t.y && h.x-2 == t.x {
			t.y++
			t.x++
		} else {
			fmt.Printf("Unexpected case: %s (%d, %d), %s (%d, %d)", h.name, h.x, h.y, t.name, t.x, t.y)
			return
		}
	} else {
		if h.y+2 == t.y {
			t.y--
			t.x = h.x
		} else if h.y-2 == t.y {
			t.y++
			t.x = h.x
		}

		if h.x-2 == t.x {
			t.x++
			t.y = h.y
		} else if h.x+2 == t.x {
			t.x--
			t.y = h.y

		}
	}
	Printlnf("%s moved to (%d, %d)", t.name, t.x, t.y)

}

func Day09_SecondAnswer(inputs []string) int {
	knots := [10]*knot{}
	for i := range knots {
		name := fmt.Sprintf("%d", i)
		if i == 0 {
			name = "H"
		}
		knots[i] = &knot{name: name}
	}

	visited := make(map[string]struct{})
	visited[fmt.Sprintf("%d:%d", knots[9].x, knots[9].y)] = struct{}{}

	for _, input := range inputs {
		Printlnf("======================\n%s", input)
		sp := strings.Split(input, " ")
		moveDir := sp[0]
		steps, _ := strconv.Atoi(sp[1])

		for i := 1; i <= steps; i++ {
			switch moveDir {
			case "L":
				knots[0].x--
				Printlnf("%s moved Left (%d, %d)\n", knots[0].name, knots[0].x, knots[0].y)
			case "R":
				knots[0].x++
				Printlnf("%s moved Right (%d, %d)\n", knots[0].name, knots[0].x, knots[0].y)
			case "U":
				knots[0].y--
				Printlnf("%s moved UP (%d, %d)\n", knots[0].name, knots[0].x, knots[0].y)
			case "D":
				knots[0].y++
				Printlnf("%s moved Down (%d, %d)\n", knots[0].name, knots[0].x, knots[0].y)
			}

			for j := 1; j <= 9; j++ {
				day09_moveTail2(knots[j-1], knots[j])
			}

			visited[fmt.Sprintf("%d:%d", knots[9].x, knots[9].y)] = struct{}{}

			for _, k := range knots {
				Printf("%s:", k.name)
				k.print()
			}
		}

	}

	return len(visited)
}
