package main

import (
	"fmt"

	"algorytms/selectionSort"
)

func main() {
	fmt.Println("Starting")

	list := []int{2, 3, 4, 1, 5, 7, 6}
	result := selectionSort.Execute(list)
	fmt.Println(result)
}
