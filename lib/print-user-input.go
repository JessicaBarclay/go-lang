package main

import "fmt"


func main() {
  // a program which takes a number entered by the user
  // and returns that number doubled
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}
