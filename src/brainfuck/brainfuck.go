// Package brainfuck impements a simple library from the language Brainfuck.
// http://en.wikipedia.org/wiki/Brainfuck
package brainfuck

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

const TAPE_SIZE = 30000

// Run the specified brainfuck code using the given input string for inputs and returns, if successful.
func Run(code string, output io.Writer, reader io.Reader) error {
	if Validate(code) == false {
		return errors.New("Code was not valid brainfuck.")
	}

	tape := make([]byte, TAPE_SIZE, TAPE_SIZE)

	index := 0

	loop := 0

	i := 0

	for i < len(code) {
		v := code[i]

		i++

		switch v {
		case '>':
			index++
		case '<':
			index--
		case '+':
			tape[index]++
		case '-':
			tape[index]--
		case '.':
			fmt.Fprintf(output, "%c", tape[index])
		case ',':
			var input byte

			n, err := fmt.Fscanf(reader, "%d", &input)

			if err == nil && n >= 1 {
				tape[index] = input
			} else {
				return err
			}
		case '[':
			if tape[index] == 0 {
				loop = 1

				for loop > 0 {
					i++

					value := code[i]

					if value == '[' {
						loop++
					} else if value == ']' {
						loop--
					}
				}

				i++
			}
		case ']':
			loop = 1
			i--

			for loop > 0 {
				i--

				value := code[i]

				if value == '[' {
					loop--
				} else if value == ']' {
					loop++
				}
			}
		}
	}

	return nil
}

// Validate the given code string and returns a boolean value corresponding to the validity of the code.
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
