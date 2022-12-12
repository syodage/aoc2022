package day11

import (
	"fmt"
	"syodage/aoc2022/utils"
)

// import "syodage/aoc2022/utils"

var (
	// LocalInput = []string{
	// 	"Monkey 0:",
	// 	"  Starting items: 79, 98",
	// 	"  Operation: new = old * 19",
	// 	"  Test: divisible by 23",
	// 	"    If true: throw to monkey 2",
	// 	"    If false: throw to monkey 3",
	// 	"",
	// 	"Monkey 1:",
	// 	"  Starting items: 54, 65, 75, 74",
	// 	"  Operation: new = old + 6",
	// 	"  Test: divisible by 19",
	// 	"    If true: throw to monkey 2",
	// 	"    If false: throw to monkey 0",
	// 	"",
	// 	"Monkey 2:",
	// 	"  Starting items: 79, 60, 97",
	// 	"  Operation: new = old * old",
	// 	"  Test: divisible by 13",
	// 	"    If true: throw to monkey 1",
	// 	"    If false: throw to monkey 3",
	// 	"",
	// 	"Monkey 3:",
	// 	"  Starting items: 74",
	// 	"  Operation: new = old + 3",
	// 	"  Test: divisible by 17",
	// 	"    If true: throw to monkey 0",
	// 	"    If false: throw to monkey 1",
	// }

	throwToFunc = func(b, c, d int64) func(item) int64 {
		return func(i item) int64 {
			if i.divisibleBy(b) {
				return c
			}
			return d
		}
	}

	LocalInput = []*monkey{
		{
			id:    0,
			items: []item{&contItem{w: 79}, &contItem{w: 98}},
			// operation:     func(a int64) int64 { return a * 19 },
			operation: func(i item) { i.mul(19) },
			throwTo:   throwToFunc(23, 2, 3),
		},
		{
			id:    1,
			items: []item{&contItem{w: 54}, &contItem{w: 65}, &contItem{w: 75}, &contItem{w: 74}},
			// operation: func(a int64) int64 { return a + 6 },
			operation: func(i item) { i.add(6) },
			throwTo:   throwToFunc(19, 2, 0),
		},
		{
			id:        2,
			items:     []item{&contItem{w: 79}, &contItem{w: 60}, &contItem{w: 97}},
			operation: func(i item) { i.squre() },
			throwTo:   throwToFunc(13, 1, 3),
		},
		{
			id:    3,
			items: []item{&contItem{w: 74}},
			// operation: func(a int64) int64 { return a + 3 },
			operation: func(i item) { i.add(3) },
			throwTo:   throwToFunc(17, 0, 1),
		},
	}

	ManagedLocalInput = []*monkey{
		{
			id:    0,
			items: []item{newMLocal(79), newMLocal(98)},
			// operation:     func(a int64) int64 { return a * 19 },
			operation: func(i item) { i.mul(19) },
			throwTo:   throwToFunc(23, 2, 3),
		},
		{
			id:    1,
			items: []item{newMLocal(54), newMLocal(65), newMLocal(75), newMLocal(74)},
			// operation: func(a int64) int64 { return a + 6 },
			operation: func(i item) { i.add(6) },
			throwTo:   throwToFunc(19, 2, 0),
		},
		{
			id:        2,
			items:     []item{newMLocal(79), newMLocal(60), newMLocal(97)},
			operation: func(i item) { i.squre() },
			throwTo:   throwToFunc(13, 1, 3),
		},
		{
			id:    3,
			items: []item{newMLocal(74)},
			// operation: func(a int64) int64 { return a + 3 },
			operation: func(i item) { i.add(3) },
			throwTo:   throwToFunc(17, 0, 1),
		},
	}
	// div  23, 19, 13, 17
	FinalInput = []*monkey{
		{
			id:    0,
			items: []item{&contItem{w: 64}, &contItem{w: 89}, &contItem{w: 65}, &contItem{w: 95}},
			// operation: func(a int64) int64 { return a * 7 },
			operation: func(i item) { i.mul(7) },
			throwTo:   throwToFunc(3, 4, 1),
		},
		{
			id:    1,
			items: []item{&contItem{w: 76}, &contItem{w: 66}, &contItem{w: 74}, &contItem{w: 87}, &contItem{w: 70}, &contItem{w: 56}, &contItem{w: 51}, &contItem{w: 66}},
			// operation: func(a int64) int64 { return a + 5 },
			operation: func(i item) { i.add(5) },
			throwTo:   throwToFunc(13, 7, 3),
		},
		{
			id:    2,
			items: []item{&contItem{w: 91}, &contItem{w: 60}, &contItem{w: 63}},
			// operation: func(a int64) int64 { return a * a },
			operation: func(i item) { i.squre() },
			throwTo:   throwToFunc(2, 6, 5),
		},
		{
			id:    3,
			items: []item{&contItem{w: 92}, &contItem{w: 61}, &contItem{w: 79}, &contItem{w: 97}, &contItem{w: 79}},
			// operation: func(a int64) int64 { return a + 6 },
			operation: func(i item) { i.add(6) },
			throwTo:   throwToFunc(11, 2, 6),
		},
		{
			id:    4,
			items: []item{&contItem{w: 93}, &contItem{w: 54}},
			// operation: func(a int64) int64 { return a * 11 },
			operation: func(i item) { i.mul(11) },
			throwTo:   throwToFunc(5, 1, 7),
		},
		{
			id:    5,
			items: []item{&contItem{w: 60}, &contItem{w: 79}, &contItem{w: 92}, &contItem{w: 69}, &contItem{w: 88}, &contItem{w: 82}, &contItem{w: 70}},
			// operation: func(a int64) int64 { return a + 8 },
			operation: func(i item) { i.add(8) },
			throwTo:   throwToFunc(17, 4, 0),
		},
		{
			id:    6,
			items: []item{&contItem{w: 64}, &contItem{w: 57}, &contItem{w: 73}, &contItem{w: 89}, &contItem{w: 55}, &contItem{w: 53}},
			// operation: func(a int64) int64 { return a + 1 },
			operation: func(i item) { i.add(1) },
			throwTo:   throwToFunc(19, 0, 5),
		},
		{
			id:    7,
			items: []item{&contItem{w: 62}},
			// operation: func(a int64) int64 { return a + 4 },
			operation: func(i item) { i.add(4) },
			throwTo:   throwToFunc(7, 3, 2),
		},
	}

	ManagedFinalInput = []*monkey{
		{
			id:    0,
			items: []item{newM(64), newM(89), newM(65), newM(95)},
			// operation: func(a int64) int64 { return a * 7 },
			operation: func(i item) { i.mul(7) },
			throwTo:   throwToFunc(3, 4, 1),
		},
		{
			id:    1,
			items: []item{newM(76), newM(66), newM(74), newM(87), newM(70), newM(56), newM(51), newM(66)},
			// operation: func(a int64) int64 { return a + 5 },
			operation: func(i item) { i.add(5) },
			throwTo:   throwToFunc(13, 7, 3),
		},
		{
			id:    2,
			items: []item{newM(91), newM(60), newM(63)},
			// operation: func(a int64) int64 { return a * a },
			operation: func(i item) { i.squre() },
			throwTo:   throwToFunc(2, 6, 5),
		},
		{
			id:    3,
			items: []item{newM(92), newM(61), newM(79), newM(97), newM(79)},
			// operation: func(a int64) int64 { return a + 6 },
			operation: func(i item) { i.add(6) },
			throwTo:   throwToFunc(11, 2, 6),
		},
		{
			id:    4,
			items: []item{newM(93), newM(54)},
			// operation: func(a int64) int64 { return a * 11 },
			operation: func(i item) { i.mul(11) },
			throwTo:   throwToFunc(5, 1, 7),
		},
		{
			id:    5,
			items: []item{newM(60), newM(79), newM(92), newM(69), newM(88), newM(82), newM(70)},
			// operation: func(a int64) int64 { return a + 8 },
			operation: func(i item) { i.add(8) },
			throwTo:   throwToFunc(17, 4, 0),
		},
		{
			id:    6,
			items: []item{newM(64), newM(57), newM(73), newM(89), newM(55), newM(53)},
			// operation: func(a int64) int64 { return a + 1 },
			operation: func(i item) { i.add(1) },
			throwTo:   throwToFunc(19, 0, 5),
		},
		{
			id:    7,
			items: []item{newM(62)},
			// operation: func(a int64) int64 { return a + 4 },
			operation: func(i item) { i.add(4) },
			throwTo:   throwToFunc(7, 3, 2),
		},
	}

	// div 3, 13, 2, 11, 5, 17, 19, 7

	// Operation , how worry level change after monkey inspect the item
	// If monkey's inspection didn'd damage the item, Int( x / 3 )

	// FinalInput = utils.ReadLines("inputs/day00_input.txt")
)

type item interface {
	wl() int64
	add(int64)
	mul(int64)
	squre()
	divisibleBy(int64) bool
}

type contItem struct {
	w int64
}

func (c *contItem) wl() int64 {
	return c.w
}

func (c *contItem) add(a int64) {
	c.w += a
}

func (c *contItem) mul(a int64) {
	c.w *= a
}

func (c *contItem) squre() {
	c.mul(c.wl())
}

func (c *contItem) divisibleBy(a int64) bool {
	return c.w%a == 0
}

func newMLocal(level int64) *managedItem {
	mm := make(map[int64]int64)
	for _, k := range []int64{23, 19, 13, 17} {
		mm[k] = level % k
	}

	return &managedItem{mm}
}

func newM(level int64) *managedItem {
	mm := make(map[int64]int64)
	for _, k := range []int64{3, 13, 2, 11, 5, 17, 19, 7} {
		mm[k] = level % k
	}

	return &managedItem{mm}
}

type managedItem struct {
	modMap map[int64]int64
}

func (m *managedItem) wl() int64 {
	panic("Doesn't support this")
	// return 0
}

func (m *managedItem) mul(a int64) {
	for k, v := range m.modMap {
		// a (y + v) => ay + av // y is divistible by k
		m.modMap[k] = (v * a) % k
	}
}

func (m *managedItem) squre() {
	for k, v := range m.modMap {
		// x(y + v) => (y + v)(y + v) => y(y + 2v) + v*v // y is divisble by k
		m.modMap[k] = (v * v) % k
	}
}

func (m *managedItem) add(a int64) {
	for k, v := range m.modMap {
		m.modMap[k] = (v + a) % k
	}
}

func (m *managedItem) divisibleBy(a int64) bool {
	v, ok := m.modMap[a]
	if !ok {
		panic(fmt.Sprintf("Invalid key: %d\n", a))
	}
	return v == 0
}

type monkey struct {
	id            int
	items         []item
	operation     func(item)
	throwTo       func(item) int64
	instpected    int64
	divideByThree bool
}

func (m *monkey) print() {
	fmt.Printf("Monkey %d: %+v, inspected: %d\n", m.id, m.items, m.instpected)
}

func FirstPart(monkeys []*monkey, rounds int) int64 {
	for _, m := range monkeys {
		m.divideByThree = true
	}
	return monkeyBusiness(monkeys, rounds)
}

func SecondPart(monkeys []*monkey, rounds int) int64 {
	return monkeyBusiness(monkeys, rounds)
}

func monkeyBusiness(monkeys []*monkey, rounds int) int64 {

	for r := 1; r <= rounds; r++ {
		for _, m := range monkeys {
			if len(m.items) == 0 {
				continue
			}

			for _, item := range m.items {
				m.instpected++
				m.operation(item)
				if m.divideByThree {
					switch ci := item.(type) {

					case *contItem:
						ci.w = int64(ci.w / 3)
						utils.Printlnf("Monkey %d, Worry level: %d", m.id, item.wl())
					default:
						panic(fmt.Sprintf("Unspported operation, devided by 3 with type: %T\n", ci))
					}
				}
				nm := m.throwTo(item)
				monkeys[nm].items = append(monkeys[nm].items, item)
			}

			m.items = []item{}

		}

		utils.Printlnf("After roound %d:", r)
		if utils.Debug {
			for _, m := range monkeys {
				m.print()
			}

		}
	}

	var a, b int64
	for _, m := range monkeys {
		m.print()
		if a < m.instpected {
			b = a
			a = m.instpected
		} else if b < m.instpected {
			b = m.instpected
		}
	}

	fmt.Printf("Most active counts, a: %d, b: %d\n", a, b)
	return a * b
}
