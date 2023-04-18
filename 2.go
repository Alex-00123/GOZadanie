package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	rows int
	cols int
	data [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}
	return &Matrix{rows, cols, data}
}

func (m *Matrix) Set(row, col, val int) {
	m.data[row][col] = val
}

func (m *Matrix) Get(row, col int) int {
	return m.data[row][col]
}

func (m *Matrix) FillRandom() {
	used := make(map[int]bool)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			var num int
			for {
				num = rand.Intn(m.rows * m.cols)
				if !used[num] {
					used[num] = true
					break
				}
			}
			m.Set(i, j, num)
		}
	}
}

func (m *Matrix) FindMaxRow() int {
	maxSum := 0
	maxRow := -1
	for i := 0; i < m.rows; i++ {
		rowSum := 0
		for j := 0; j < m.cols; j++ {
			rowSum += m.Get(i, j)
		}
		if rowSum > maxSum {
			maxSum = rowSum
			maxRow = i
		}
	}
	return maxRow
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rows, cols := 5, 10
	matrix := NewMatrix(rows, cols)
	matrix.FillRandom()

	maxRow := matrix.FindMaxRow()

	fmt.Println("Массив:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%3d ", matrix.Get(i, j))
		}
		fmt.Println()
	}
	fmt.Printf("Строка max sum: %d\n", maxRow)
}