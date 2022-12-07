package main

func Day06_FirstAnswer(inputs []string) []int {
	var res []int
	for _, input := range inputs {
		res = append(res, day06_consecutiveRunes(input, 4))
	}
	return res
}

func Day06_SecondAnswer(inputs []string) []int {
	var res []int
	for _, input := range inputs {
		res = append(res, day06_consecutiveRunes(input, 14))
	}
	return res
}

func day06_consecutiveRunes(input string, n int) int {
	var arr []rune
	for i, r := range input {
		// fmt.Printf("Got: %c\n", r)
		if arr == nil {
			arr = append(arr, r)
			continue
		}

		for j, p := range arr {
			if p == r {
				arr = arr[j+1:]
				break
			}
		}

		arr = append(arr, r)
		// fmt.Printf("Arr: %+v\n", arr)
		if len(arr) == n {
			return i + 1
		}
	}
	return -1

}
