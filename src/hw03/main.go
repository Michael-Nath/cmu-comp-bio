package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Monkey Moo")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(RandomWalk(10, 6))
}

// InitializeMatrix : given #rows and #cols, returns a 2d array with those dimensions.
func InitializeMatrix(numRows, numCols int) [][]int {
	myMatrix := make([][]int, numRows)
	for i := range myMatrix {
		myMatrix[i] = make([]int, numCols)
	}
	return myMatrix
}

// BinarizeMatrix : given a matrix and a threshold, returns a new matrix such that b[i][j] is true if matrix[i][j] >= threshold
func BinarizeMatrix(mtx [][]int, threshold int) [][]bool {
	rows := getRows(mtx)
	cols := getCols(mtx)

	// create b
	B := make([][]bool, rows)
	for j := range B {
		B[j] = make([]bool, cols)
	}

	// set contents of B depending on whether corresponding position in mtx is at least the threshold value
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mtx[i][j] >= threshold {
				B[i][j] = true
			} else {
				B[i][j] = false
			}
		}
	}
	return B
}

// getRows : returns the # of rows in a 2d matrix
func getRows(mtx [][]int) int {
	return len(mtx)
}

// getCols : returns the # of columns in a 2d matrix
func getCols(mtx [][]int) int {
	if len(mtx[0]) == 0 {
		panic("This isn't a valid matrix")
	}
	return len(mtx[0])
}

// ConnectFour : given an matrix B,  return true if there are four consecutive values in the same row, column, or diagonal that have the same value; it should return false otherwise.
func ConnectFour(B [][]int) bool {
	rows := getRows(B) // get rows
	cols := getCols(B) // get columns

	// traverse through the matrix B
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if DetermineHorizontal(B, i, j) || DetermineVertical(B, i, j) || DetermineDiagonal(B, i, j) {
				return true
			}
		}
	}

	// if no such index is found, return true
	return false
}

// DetermineHorizontal : subroutine to determine if 4 consecutive values are the same horizontally
func DetermineHorizontal(B [][]int, i, j int) bool {
	count := 0
	// traverse columns
	for m := j + 1; m < getCols(B); m++ {
		if B[i][m] == B[i][j] {
			count++
		}
	}
	if count >= 3 {
		return true
	}
	return false
}

// DetermineVertical : subroutine to determine if 4 consecutive values are the same vertically
func DetermineVertical(B [][]int, i, j int) bool {
	count := 0
	// traverse columns
	for m := i + 1; m < getRows(B); j++ {
		if B[m][j] == B[i][j] {
			count++
		}
	}
	if count >= 3 {
		return true
	}
	return false
}

// DetermineDiagonal : subroutine to determine if 4 consecutive values are the same diagonally
func DetermineDiagonal(B [][]int, i, j int) bool {
	count := 0
	// traverse columns
	for m := 1; i+m < getRows(B) && j+m < getCols(B); m++ {
		if B[i+m][j+m] == B[i][j] {
			count++
		}
	}
	if count >= 3 {
		return true
	}
	return false
}

// RandomWalk : given a grid and a number of iterations, simulates a bacterial walk
func RandomWalk(numSteps, width int) [][]int {
	curX := 0
	curY := 0
	myWalks := InitializeMatrix(numSteps, 2)
	for i := 0; i < numSteps; i++ {
		curX, curY = DetermineWalk(curX, curY, width)
		myWalks[i][0] = curX
		myWalks[i][1] = curY
	}
	return myWalks
}

// DetermineWalk : subroutine to determine the next random move
func DetermineWalk(x, y, width int) (int, int) {
	moveX := rand.Intn(3) - 1
	moveY := rand.Intn(3) - 1
	rand.Seed(time.Now().UnixNano())
	for (moveX+x > width || moveX+x < (-1*width)) || (moveY+y > width || moveY+y < (-1*width)) {
		rand.Seed(time.Now().UnixNano())
		if moveX+x > width || moveX+x < (-1*width) {
			moveX = rand.Intn(3) - 1
		}
		if moveY+y > width || moveY+y < (-1*width) {
			moveY = rand.Intn(3) - 1
		}
	}
	return x + moveX, y + moveY
}
