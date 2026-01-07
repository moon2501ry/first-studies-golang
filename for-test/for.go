package main

import "fmt"

func main() {
	var numbers [3]int
	fmt.Scanf("%d,%d,%d", &numbers[0], &numbers[1], &numbers[2])
	for i, n := range numbers {
		fmt.Printf("(%d, %d)\n", i, n)
	}
}
