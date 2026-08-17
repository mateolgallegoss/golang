package main

import (
  "fmt"
  "math/rand"
)

func main(){
  fmt.Println("------------------------ EXERCISE 1 ------------------------")
  s := make([]int, 100)
  for i := range s{
    s[i] = rand.Intn(100)
  }
  fmt.Printf("Slice original: %d \n", s)

  fmt.Println("------------------------ EXERCISE 2 ------------------------")
  for _, v := range s{
    if v %2 == 0{
      fmt.Println("Two!")
    }
    if v %3 == 0{
      fmt.Println("Three!")
    }
    if v %2 == 0 && v %3 == 0{
      fmt.Println("Six!")
    }
    fmt.Println("Never mind")
  }

  fmt.Println("------------------------ EXERCISE 3 ------------------------")
  var total int
  //Loop can be modernize with the rang built-in keyword, but this was made in purpose for practice
  //The variable total was scope to a iteration of the for range, no to all interations (control structure for range)
  for i := 0; i < 10; i++{
    total := total + i
    fmt.Println(total)
  } 
}
