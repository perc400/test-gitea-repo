package main

import "fmt"

func main() {
  fmt.Println("---")
  g := 1
  fmt.Println(getTest(g))
  fmt.Println("main")
}

func getTest(id int) string {
  return "test"
}
