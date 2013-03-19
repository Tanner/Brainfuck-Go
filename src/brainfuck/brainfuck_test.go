package brainfuck

import "testing"

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