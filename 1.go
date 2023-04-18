package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Создаем слайс из 10 случайных чисел
    nums := make([]int, 10)
    for i := range nums {
        nums[i] = rand.Intn(100)
    }

    fmt.Println("Исходный слайс:", nums)

    // Сортировка пузырьком
    bubbleSort(nums)
    fmt.Println("Слайс после сортировки пузырьком:", nums)

    // Сортировка вставками
    insertionSort(nums)
    fmt.Println("Слайс после сортировки вставками:", nums)

    // Сортировка слиянием
    mergeSort(nums)
    fmt.Println("Слайс после сортировки слиянием:", nums)
}

// Сортировка пузырьком
func bubbleSort(nums []int) {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}

// Сортировка вставками
func insertionSort(nums []int) {
    n := len(nums)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 && nums[j-1] > nums[j] {
            nums[j-1], nums[j] = nums[j], nums[j-1]
            j--
        }
    }
}

// Сортировка слиянием
func mergeSort(nums []int) {
    n := len(nums)
    if n < 2 {
        return
    }
    mid := n / 2
    left := make([]int, mid)
    right := make([]int, n-mid)
    for i := 0; i < mid; i++ {
        left[i] = nums[i]
    }
    for i := mid; i < n; i++ {
        right[i-mid] = nums[i]
    }
    mergeSort(left)
    mergeSort(right)
    merge(nums, left, right)
}

func merge(nums, left, right []int) {
    i, j, k := 0, 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            nums[k] = left[i]
            i++
        } else {
            nums[k] = right[j]
            j++
        }
        k++
    }
    for i < len(left) {
        nums[k] = left[i]
        i++
        k++
    }
    for j < len(right) {
        nums[k] = right[j]
        j++
        k++
    }
}
