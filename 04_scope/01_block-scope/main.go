package main

import (
	"fmt"
	"github.com/GoesToEleven/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)


func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max)

	fmt.Println(vis.MyName)
	vis.PrintVar()
}

