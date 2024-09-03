package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for i := 0; i <= len(nums2); i++ {
		for j := 0; j <= len(nums1); j++ {
			if nums2[i] > nums1[j] {
				arr := []int{}
				arr = append(arr, nums2[i])
				nums1 = append(nums1[:j+1], append(arr, nums1[j+1:]...)...)
				fmt.Println(nums1)
				break
			}
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	n, m := 3, 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1, nums2)
}
