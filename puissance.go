package main

import "fmt"

func main() {
	var n int
	fmt.Println("donner nombre")
	fmt.Scanln(&n)
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
