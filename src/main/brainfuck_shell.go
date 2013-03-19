package main

import (
  "fmt"
  "brainfuck"
  "flag"
  "io/ioutil"
)

func main() {
  var path = flag.String("filename", "brainfuck.bf", "Name of file to run")
  var validate = flag.Bool("validate", false, "Validates instead of running")

  flag.Parse();

  content, err := ioutil.ReadFile(*path)

  if err != nil {
    fmt.Println(err)

    return
  }

  code := string(content)

  if (*validate) {
    if brainfuck.Validate(code) {
      fmt.Println("Valid code.")
    } else {
      fmt.Println("Non-valid code.")
    }
  } else {
    brainfuck.Run(code)
  }
}