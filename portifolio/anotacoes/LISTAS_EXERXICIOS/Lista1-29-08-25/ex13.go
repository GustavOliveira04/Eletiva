// Escreva um algoritmo que retorne a moda de uma sequencia numerica

package main

import "fmt"

func moda(nums []int) int {
	freq := make(map[int]int)
	maiorFreq, moda := 0, nums[0]
	for _, n := range nums {
		freq[n]++
		if freq[n] > maiorFreq {
			maiorFreq = freq[n]
			moda = n
		}
	}
	return moda
}

func main() {
	fmt.Println("Moda:", moda([]int{1, 2, 2, 3, 3, 3, 4}))
}
