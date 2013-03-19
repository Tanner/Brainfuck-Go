package brainfuck

import (
  "bytes"
  "errors"
)

func Run(code string) error {
  if Validate(code) == false {
    return errors.New("Code was not valid brainfuck.")
  }

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