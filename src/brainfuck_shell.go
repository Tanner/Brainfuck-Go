package main

import (
  "fmt"
  "brainfuck"
)

func main() {
  err := brainfuck.Run("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")

  fmt.Println("\n", err)
}