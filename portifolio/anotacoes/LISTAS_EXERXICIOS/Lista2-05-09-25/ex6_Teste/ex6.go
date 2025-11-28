package main

import (
	"fmt"
	"sort"
)

func ordenaNumeros(nums []int) {
	sort.Ints(nums)
	fmt.Println("Ordenados: ", nums)
}

func main() {
	ordenaNumeros([]int{6, 8, 2, 1, 4, 0, 9, 3, 5, 7})
}
