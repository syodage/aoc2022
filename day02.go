package main

import (
  "fmt"
	"log"
	"strings"
)


func Day02_FirstAnswer(inputs []string) int {
  // A - Rock, B - Paper, C - Scissors
  // openant, <your-reason>
  // X - Rock, Y - Paper and Z - Scissors
  // Scores: 1 - Rock, 2 - Paper and 3 - Scissors
  // 0 - locss, 3 - draw, and 6 - win

  // rock < paper < scissors < rock

  draw := map[string]string{
    "A": "X",
    "B": "Y",
    "C": "Z",
  }

  win := map[string]string {
    "A": "Y",
    "B": "Z",
    "C": "X",
  }

  points := map[string]int {
    "X": 1,
    "Y": 2,
    "Z": 3,
  }


  fmt.Println("Starting")
  var totalScore int
  var score int
  for i, input := range inputs {
    values := strings.Split(input, " ")
    if len(values) != 2 {
      log.Fatalf("Invalid input(line: %d) length of values[%q] got: %d, want: 2",i, input, len(values))
    }

    op, me := values[0], values[1]
    // fmt.Printf("op: %q, you: %q\n", op, me)

    score = points[me] 

    if win[op] == me {
      score += 6
    } else if draw[op] == me {
      score += 3
    } 
    
    // fmt.Printf("score: %d\n", score)

    totalScore += score
    // fmt.Printf("totalScore: %d\n\n", totalScore)
  }

  return totalScore 
}




func Day02_SecondAnswer(inputs []string) int {
  // X - lose, Y - Draw and Z - Win

  lose := map[string]string{
    "A": "C",
    "B": "A",
    "C": "B",
  }

  win := map[string]string {
    "A": "B",
    "B": "C",
    "C": "A",
  }

  points := map[string]int {
    "A": 1,
    "B": 2,
    "C": 3,
  }

  fmt.Println("Starting")
  var totalScore int
  for i, input := range inputs {
    values := strings.Split(input, " ")
    if len(values) != 2 {
      log.Fatalf("Invalid input(line: %d) length of values[%q] got: %d, want: 2",i, input, len(values))
    }

    op, res := values[0], values[1]

    if res == "X" {
      totalScore += points[lose[op]] + 0
    } else if res == "Z" {
      totalScore += points[win[op]] + 6
    } else {
      totalScore += points[op] + 3
    }
  }

  return totalScore 
} 
