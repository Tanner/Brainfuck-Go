package brainfuck

import (
  "bytes"
)

func Run(code string) error {
  // array := [30000]byte

  return nil
}

func Validate(code string) bool {
  loop_count := 0

  characters := []byte{'>', '<', '+', '-', '.', ',', '[', ']'}

  for _, v := range code {
    if bytes.IndexRune(characters, v) == -1 {
      return false
    }

    if v == '[' {
      loop_count++
    } else if v == ']' {
      loop_count--
    }
  }

  if loop_count != 0 {
    return false
  }

  return true
}