package main

import ("fmt"
        "math/rand"
        "time")

func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  fmt.Println("A number from 1 to 100", rand.Intn(100))
  fmt.Println("f", time.Tick(5))
}



