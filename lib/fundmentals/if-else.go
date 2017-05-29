package main

import "fmt"

func main() {

  // basic if statement - parentheses optional, braces required
  if 7 + 3 == 5 * 2 {
    fmt.Println("if condition true")
  } else {
    fmt.Println("why am I even here...")
  }

  // if statement without an else
  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  // A statement can precede conditionals;
  // Any variables declared in this statement are available in all branches.
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}
