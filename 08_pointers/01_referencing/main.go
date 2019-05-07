package main

import "fmt"

func main()  {
	a := 43
	fmt.Println(a)   // 43
	fmt.Println(&a)  // 0xc042054080

	var b = &a
	fmt.Println(b)  // 0xc042054080
	fmt.Println(*b)  // 43

	*b = 42 // change the value at this address to 42
	fmt.Println(*b)
}