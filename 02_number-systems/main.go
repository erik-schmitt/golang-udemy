package main

import "fmt"

func main() {
	fmt.Println(42) // Decimal

	fmt.Printf("%d - %b \n", 42, 42) // Decimal - Binary

	fmt.Printf("%d - %b - %x\n", 42, 42, 42) // Decimal - Binary - Hexadecimal

	fmt.Printf("%d - %b - %#x\n", 42, 42, 42) // Decimal - Binary - Hexadecimal with leading 0x

	fmt.Printf("%d - %b - %#X\n", 42, 42, 42) // Decimal - Binary - Hexadecimal with leading 0X and capital letters

	for i := 0;  i < 200; i++ {
		fmt.Printf("%d \t %b \t  %x \t %q \n", i , i, i, i) // %q is UTF-8
	}
}
