package main

import (
	"fmt"
)


func _Function() {
  myMessage() // call the function
  familyName("Ashraful")
   fmt.Println(mathFunction(1, 2))
}

func myMessage() {
  fmt.Println("function call work!")
}

func familyName(fname string) {
  fmt.Println("Hello", fname, "pathan")
}

func mathFunction(x int, y int) int {
  return x + y
}