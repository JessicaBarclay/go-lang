package main

import "fmt"

func main(){

  // simple string
  fmt.Println("hello world\n")
  fmt.Println("go" + "-lang\n")
  fmt.Println("1+1=", 1+1, "\n")
  fmt.Println("7.5 / 1.45=", 7.5 / 1.45, "\n")

  // operators
  fmt.Println(true || false, "\n")
  fmt.Println(false && true, "\n")
  fmt.Println(!false, "\n")

  // variables
  var str string = "hello"
  fmt.Println(str, "\n")

  var x, y, z int = 1, 2, 3
  fmt.Println(x + y + z, "\n")

  f := "foo"
  fmt.Println(f)

  // define multiple variables
  var (
    a = 1
    b = 2
    c = 3
  )

  fmt.Println(a, b, c)
}
