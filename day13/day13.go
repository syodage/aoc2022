package day13

import (
	"fmt"
	"strconv"
	"syodage/aoc2022/utils"
)

// import "syodage/aoc2022/utils"

var (
	SmallInput = []string{
		"[[1],[2,3,4]]",
		"[[1],4]",
	}
	LocalInput = []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}

	FinalInput = utils.ReadLines("inputs/day13_input.txt")
)

type PacketType int

const (
	List PacketType = iota
	Integer
	Root
)

type packet struct {
	_type  PacketType
	value  int
	values []*packet
}

func (p *packet) print() {
	switch p._type {
	case Root:
		if len(p.values) > 0 {
			fmt.Print("<")
			for i, c := range p.values {
				c.print()
				if i != len(p.values)-1 {
					fmt.Print(",")
				}
			}
			fmt.Print(">")
		} else {
			fmt.Printf("%d", p.value)
		}

	case Integer:
		fmt.Printf("%d", p.value)
	case List:
		fmt.Print("[")
		for i, c := range p.values {
			c.print()
			if i != len(p.values)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
	}
}

func (lhs *packet) compare(rhs *packet) int {
	if utils.Debug {
		utils.Printlnf("Compare:")
		lhs.print()
		utils.Printlnf("\n")
		rhs.print()
		utils.Printlnf("====")
	}
	// if len(lhs) > len(rhs)
	if rhs == nil {
		return 1
	}

	// convert types
	if lhs._type == Integer && rhs._type == List {
		converted := &packet{
			_type: List,
			values: []*packet{
				lhs,
			},
		}

		return converted.compare(rhs)
	} else if lhs._type == List && rhs._type == Integer {
		converted := &packet{
			_type: List,
			values: []*packet{
				rhs,
			},
		}
		return lhs.compare(converted)
	}

	if lhs._type == Integer && rhs._type == Integer {
		utils.Printlnf("lhs val: %d, rhs val: %d", lhs.value, rhs.value)
		if lhs.value < rhs.value {
			return 1
		} else if lhs.value == rhs.value {
			return 0
		} else {
			return -1
		}
	}

	if lhs._type == List && rhs._type == List ||
		lhs._type == Root && rhs._type == Root {

		for i, node := range lhs.values {
			if i >= len(rhs.values) {
				return -1
			}

			ch := node.compare(rhs.values[i])
			if ch != 0 {
				return ch
			}
		}

		if len(lhs.values) < len(rhs.values) {
			return 1
		}

		return 0
	}
	return 0
}

func FirstPart(inputs []string) int {
	var firstPacket, secondPacket string
	var totalScore int
	var pair int
	for i, row := range inputs {
		ll := (i + 1) % 3
		switch ll {
		case 1:
			firstPacket = row
		case 2:
			secondPacket = row
		default:
			pair++
			utils.Printlnf("1: %s, 2: %s", firstPacket, secondPacket)
			if compare(firstPacket, secondPacket) >= 0 {
				fmt.Printf("Pair: %d is in right order\n", pair)
				totalScore += pair
			}
		}
	}
	pair++
	utils.Printlnf("1: %s, 2: %s", firstPacket, secondPacket)
	if compare(firstPacket, secondPacket) >= 0 {
		fmt.Printf("Pair: %d is in right order\n", pair)
		totalScore += pair
	}
	return totalScore
}

func compare(first, second string) int {
	utils.Printlnf("Parsing %s", first)
	r1 := parse(first)
	utils.Printlnf("Parsing %s", second)
	r2 := parse(second)
	utils.Printf("\nComparing following:")
	if utils.Debug {
		r1.print()
	}
	utils.Printlnf("\n<===>")
	if utils.Debug {
		r2.print()
	}
	utils.Printlnf("\n==============")

	return r1.compare(r2)
}

// if a < b return -1
// if a == b return 0
// if a > b return 1
func parse(input string) *packet {
	r := &packet{
		_type: Root,
	}

	parseInner(input, r)
	return r
}

func parseInner(input string, p *packet) string {
	if input == "" {
		panic("oh noh")
	}

	utils.Printlnf("Parsing: %s", input)

	r := input[0]
	suffix := input[1:]

	if r == '[' {
		curr := &packet{
			_type: List,
		}
		p.values = append(p.values, curr)

		utils.Printlnf(">>>")
		suffix = parseInner(suffix, curr)
		for suffix[0] != ']' {
			suffix = parseInner(suffix, curr)
		}
		utils.Printlnf("<<<")
		suffix = suffix[1:]
	} else if r == ',' {
		// nothing
	} else if r == ']' {
		suffix = input
	} else {
		// integer
		val, l := readInt(input)
		curr := &packet{
			_type: Integer,
			value: val,
		}
		p.values = append(p.values, curr)
		suffix = input[l:]
	}
	return suffix
}

func readInt(str string) (int, int) {
	var res int
	for i, r := range str {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return res, i
		}

		res = res*10 + d
	}
	return res, len(str)
}
func SecondPart(inputs []string) int {
	var signals []*packet

	div1 := parse("[[2]]")
	div2 := parse("[[6]]")

	for _, signal := range inputs {
		if signal == "" {
			continue
		}

		sig := parse(signal)
		signals = append(signals, sig)
	}
	signals = append(signals, div1)
	signals = append(signals, div2)

	sorted := sort(signals)

	var indeces []int

	for i, p := range sorted {
		p.print()
		fmt.Println("")
		if p.compare(div1) == 0 || p.compare(div2) == 0 {
			fmt.Println("found: ")
			p.print()
			fmt.Printf(", in index: %d\n", i+1)
			indeces = append(indeces, i+1)
		}
	}

	return indeces[0] * indeces[1]
}

func sort(pa []*packet) []*packet {
	var sorted []*packet

	for _, a := range pa {

		// printHelper("Before Adding:", a, sorted)

		if len(sorted) == 0 {
			sorted = append(sorted, a)
			continue
		}

		pos := binarySearch(sorted, a, 0, len(sorted)-1)
		// fmt.Printf("Pos: %d\n", pos)
		temp := append([]*packet{}, sorted[pos:]...)
		sorted = append(sorted[0:pos], a)
		sorted = append(sorted, temp...)
		// printHelper("After Adding:", a, sorted)

	}
	return sorted
}

func printHelper(format string, p *packet, sorted []*packet) {
	fmt.Println("")
	fmt.Print(format)
	p.print()

	fmt.Println("Sorted list: ")
	for _, b := range sorted {
		b.print()
		fmt.Println("")
	}
	fmt.Println("====")
}

func binarySearch(sorted []*packet, p *packet, low, high int) int {
	if high <= low {
		s := p.compare(sorted[low])
		if s == 1 {
			return low
		}
		return low + 1
	}

	mid := (low + high) / 2
	compMid := p.compare(sorted[mid])
	if compMid == 0 {
		return mid + 1
	}

	if compMid < 0 {
		return binarySearch(sorted, p, mid+1, high)
	}

	return binarySearch(sorted, p, low, mid-1)
}
