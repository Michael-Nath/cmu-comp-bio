package main

//MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
//It returns adjacency lists of reads; edges are only included
//in the overlap graph is the alignment score of the two reads is at least the threshold.

//get scoring matrix
//binarize the matrix given
//create map
//input the values into the map

func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	mtx := OverlapScoringMatrix(reads, match, mismatch, gap)
	mtx1 := BinarizeMatrix(mtx, threshold)
	adjacencyList := make(map[string][]string)

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx); j++ {
			if mtx1[i][j] == 1 {
				adjacencyList[reads[i]] = append(adjacencyList[reads[i]], reads[j])
			}
		}
	}
	// fill in details of the function here.

	return adjacencyList
}

/*
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
			} else if mtx1[i][j] >= threshold && mtx1[j][i] < threshold {
				mtx[i][j] = 1
				mtx[j][i] = 0
			} else if mtx1[j][i] >= threshold && mtx1[i][j] < threshold {
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
/*
/*
func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]float64 {
	mtx := InitializeSquareMatrix(len(reads))
	numRows := len(mtx)
	numCols := len(mtx[0])
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if i != j {
				mtx[i][j] = ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap)
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

func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)
	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}
*/
