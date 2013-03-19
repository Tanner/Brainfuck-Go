package main

import (
  "fmt"
  "brainfuck"
)

func main() {
  valid := brainfuck.Validate("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")

  fmt.Println(valid)
}