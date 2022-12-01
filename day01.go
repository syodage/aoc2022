package main

import (
	"log"
	"strconv"
)

var inputs = ReadLines("inputs/day01_input.txt")

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
func Day01_FirstAnswer() int {
  var maxCal, curCal int

  for _, input := range inputs {
    if input == "" {
      if maxCal < curCal {
        maxCal = curCal 
      }
      curCal = 0
      continue
    }

    calaries, err := strconv.Atoi(input)
    if err != nil {
      log.Fatalf("Error converting str %q, to int, error: %v", input, err)
    }
    curCal += calaries
    // fmt.Println(scanner.Text())
  }

  if maxCal < curCal {
    maxCal = curCal
  }
 
  return maxCal 
}


// Find the top three Elves carrying the most Calories.
// How many Calories are those Elves carrying in total?
func Day01_SecondAnswer() int {
  var maxCal1, maxCal2, maxCal3, curCal int

  for _, input := range inputs {
    if input == "" {
      if maxCal1 < curCal {
        maxCal3 = maxCal2
        maxCal2 = maxCal1
        maxCal1 = curCal
      } else if maxCal2 < curCal {
        maxCal3 = maxCal2
        maxCal2 = curCal 
      } else if maxCal3 < curCal {
        maxCal3 = curCal 
      }
      curCal = 0
      continue
    }

    calaries, err := strconv.Atoi(input)
    if err != nil {
      log.Fatalf("Error converting str %q, to int, error: %v", input, err)
    }
    curCal += calaries
    // fmt.Println(scanner.Text())
  }

  if maxCal3 < curCal {
    maxCal3 = curCal
  }
 
  top3Total := maxCal1 + maxCal2 + maxCal3
  // fmt.Printf("Top 3 Elf carrying Calories: %d\n", top3Cal)
  return top3Total 
}
