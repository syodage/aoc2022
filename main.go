package main

import (
  "fmt"
)

func main() {
  // fmt.Println("Hello AoC++")
  // fmt.Printf("Day01_01: Max Calary sum: %d\n", Day01_FirstAnswer())
  // fmt.Printf("Day01_02: Sum of top 3 max Calary sums: %d\n", Day01_SecondAnswer())


  day02Input := ReadLines("inputs/day02_input.txt")
  // fmt.Printf("Day02_01: Total Score: %d\n", Day02_FirstAnswer(day02Input))
  // day02Input := []string{"A Y", "B X", "C Z"} 
  fmt.Printf("Day02_02: Total Score: %d\n", Day02_SecondAnswer(day02Input))

}

