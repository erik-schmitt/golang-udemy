package main

import "fmt"

func main()  {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)

	yards := meters * metersToYards
	fmt.Println(meters, " meters is", yards, " yards.")
}