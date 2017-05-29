package main

import "fmt"
import "time"


func main() {

  // Switch statements express conditionals across many branches.
  // Basic switch
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  // optional default statement
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
      fmt.Println("It's a weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's past noon")
  }
// A type switch compares types instead of values.
  / You can use this to discover the the type of an interface value.
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Println("Don't know what type \n", t)
      }
    }
  whatAmI(true)
  whatAmI(7)
  whatAmI("unknown")
}
