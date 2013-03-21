package brainfuck

import (
	"bytes"
	"strings"
	"testing"
)

func TestValidation(t *testing.T) {
	if !Validate("><+-.,[]") {
		t.Error("Expected valid tokens to pass validation")
	}

	if Validate("><+-.,[]a!$") {
		t.Error("Expected invalid tokens to fail validation")
	}

	if Validate("><+-.,[") || Validate("><+-.,[[]") {
		t.Error("Expected unmatched '[' token to fail validation")
	}

	if Validate("><+-.,]") || Validate("><+-.,[]]") {
		t.Error("Expected unmatched ']' token to fail validation")
	}
}

func TestRunning(t *testing.T) {
	helper := func(code string, in string, correctOutput string) {
		output := new(bytes.Buffer)
		input := strings.NewReader(in)

		err := Run(code, output, input, 10000)

		if err == nil {
			if output.String() != correctOutput {
				t.Errorf("'%s' want '%s'", output, correctOutput)
			}
		} else {
			t.Error(err)
		}
	}

	// Test "Hello World!"
	helper("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.", "", "Hello World!")

	// Test incrementation
	helper("++++++++++++++++++++++++++++++++++++++++++++++++.", "", "0")

	// Test subtraction
	helper("+++++++++++++++++++++++++++++++++++++++++++++++++-.", "", "0")

	// Test looping
	helper("+++++[>++++++++++<-]>++.", "", "4")

	// Test input
	helper(",.", "57", "9")

	// Test negative index on tape
	output := new(bytes.Buffer)
	input := strings.NewReader("")

	err := Run("<-", output, input, 10)

	if err == nil {
		t.Error("Expected an error about the index for the tape being out of bounds")
	}

	// Test exceeding cycles
	output.Reset()
	input = strings.NewReader("")

	err = Run("-[-]", output, input, 10)

	if err == nil {
		t.Error("Expected an error about max cycles being exceeded")
	}
}
