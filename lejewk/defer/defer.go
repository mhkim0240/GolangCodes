package main

import "fmt"

func main() {
	fmt.Println("Main func start")                 // 1
	defer func() { fmt.Println("First defer") }()  // 4
	defer func() { fmt.Println("Second defer") }() // 3
	fmt.Println("Main func end")                   // 2
}
