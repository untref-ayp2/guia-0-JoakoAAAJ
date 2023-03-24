package main

import "fmt"

func minMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return min, max
}

func main() {
	nums := []int{10, 5, 20, 15, 30}
	min, max := minMax(nums)
	fmt.Println("El número menor es:", min)
	fmt.Println("El número mayor es:", max)
}
