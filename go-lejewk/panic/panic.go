package main

import "fmt"

func main() {

	fmt.Println("Main func start")
	defer func() { fmt.Println("Main func first defer") }()

	foo()

	defer func() { fmt.Println("Main func second defer") }()
	fmt.Println("Main func end")

}

func foo() {
	fmt.Println("Foo func start")
	defer func() { fmt.Println("Foo func first defer") }()

	bar()

	defer func() { fmt.Println("Foo func second defer") }()
	fmt.Println("Foo func end")
}

func bar() {
	fmt.Println("Bar func start")
	defer func() { fmt.Println("Bar func first defer") }()
	panic("Ops..")
	defer func() { fmt.Println("Bar func second defer") }()
	fmt.Println("Bar func end")
	//defer func() { fmt.Println("bar func second defer") }()
}
