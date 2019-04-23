package main

import "fmt"

const (
	_ = iota // 0
	kB = 1 << (iota * 10) // 1 << (1 * 10)
	mB = 1 << (iota * 10) // 1 << (2 * 20)
	gB = 1 << (iota * 10) // 1 << (3 * 30)
)

func main() {
	fmt.Println("binary\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", kB)
	fmt.Printf("%d\n", kB)
	fmt.Printf("%b\t\t\t", mB)
	fmt.Printf("%d\n", mB)
	fmt.Printf("%b\t\t\t", gB)
	fmt.Printf("%d\n", gB)
}