package main

//BinarizeMatrix takes a matrix of values and a threshold.
//It binarizes the matrix according to the threshold.
//If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx1 [][]float64, threshold float64) [][]int {
		mtx := InitializeIntMatrix(len(mtx1))
		for i := 0; i < len(mtx); i++ {
				for j := 0; j < len(mtx1[i]); j++ {
						if mtx1[i][j] >= threshold && mtx1[j][i] >= threshold {
								if mtx1[i][j] > mtx1[j][i] {
										mtx[i][j] = 1
										mtx[j][i] = 0
								} else {
										mtx[i][j] = 0
										mtx[j][i] = 1
								}
						} else if mtx1[i][j] >= threshold && mtx1[j][i] < threshold{
								mtx[i][j] = 1
								mtx[j][i] = 0
						} else if mtx1[j][i] >= threshold && mtx1[i][j] < threshold{
								mtx[i][j] = 0
								mtx[j][i] = 1
						} else {
								mtx[i][j] = 0
								mtx[j][i] = 0
						}

				}
		}
		return mtx
}


func InitializeIntMatrix(n int) [][]int {
		mtx := make([][]int, n)
		for i := range mtx {
				mtx[i] = make([]int, n)
		}
		return mtx
}
