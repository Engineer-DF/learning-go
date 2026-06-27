package main

import (
	"fmt"
)

func measureTime() {
	defer fmt.Println("end")
	fmt.Println("start")
}

func deferOrder() {
	defer fmt.Println("third")
	defer fmt.Println("second")
	defer fmt.Println("first")
}

func someFunc() {
	a := 5
	defer fmt.Println(a)
	a = 10
	fmt.Println(a)
}

func main() {
	measureTime()
	deferOrder()
	someFunc()
}
