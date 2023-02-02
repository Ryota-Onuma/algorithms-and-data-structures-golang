package main

import "fmt"

func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    target := 5
    index := binarySearch(nums, target)
    if index == -1 {
        fmt.Println("Target not found")
    } else {
        fmt.Println("Target found at index", index)
    }
}

