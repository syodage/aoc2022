package main

import (
	"fmt"
	"strconv"
	"strings"
)


func Day05_FirstAnswer(inputs []string, crates map[int][]rune) string {
    fmt.Printf("Map: %+v\n", crates)
  // input is in format of "move x crates from y to z"
  for _, input := range inputs {
    words := strings.Split(input, " ")
    x, _ := strconv.Atoi(words[1])
    y, _ := strconv.Atoi(words[3])
    z, _ := strconv.Atoi(words[5])

    // fmt.Printf("Move %d from %d to %d\n", x, y, z)

    for i := 0; i < x ; i++ {
      lenY := len(crates[y])
      crates[z] = append(crates[z], crates[y][lenY-1])
      crates[y] = crates[y][:lenY-1]
    }

    // fmt.Printf("After Move: %+v\n", crates)
  } 

  var res string
  for j := 1; j <= len(crates); j++ {
    v := crates[j]
    res += string(v[len(v)-1])
  }
  return res 
}

func Day05_SecondAnswer(inputs []string, crates map[int][]rune) string {
    fmt.Printf("Map: %+v\n", crates)
  for _, input := range inputs {
    words := strings.Split(input, " ")
    x, _ := strconv.Atoi(words[1])
    y, _ := strconv.Atoi(words[3])
    z, _ := strconv.Atoi(words[5])

    // fmt.Printf("Move %d from %d to %d\n", x, y, z)

  // input is in format of "move x crates from y to z"
      lenY := len(crates[y])
      crates[z] = append(crates[z], crates[y][lenY-x:]...)
      crates[y] = crates[y][:lenY-x]

    // fmt.Printf("After Move: %+v\n", crates)
  } 

  var res string
  for j := 1; j <= len(crates); j++ {
    v := crates[j]
    res += string(v[len(v)-1])
  }
  return res 
}

