package main

import "fmt"
import "math"

const s string = "string constant"

func main() {

  // constants
    fmt.Println(s)

    const x = 789
    const y = 123456 / x
    const z = 72e5

    fmt.Println(y)
    fmt.Println(int64(z))
    fmt.Println(math.Sin(x))

  // basic for loop with single condition
    i := 1
    for i <=3 {
        fmt.Println(i)
        i = i + 1
    }

  // classic initial - condition - after
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

  // for without condition, loop until break or return
    for {
        fmt.Println("loop without condition")
        break
    }

  // continue to the next iteration of the loop

    for n := 0; n <= 5; n++ {
        if n % 2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
