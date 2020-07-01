// Compass Quad 2020 (Abby M-P, Michael Nath, Kevin Fu)

package BetaDiversity

import (
	"Metagenomics/BrayCurtis"
	"Metagenomics/Jaccard"
)

// Write the functions BetaDiversityMatrix() with the following input and output
// Input: a slice of frequency maps samples and a distance metric ("Bray-Curtis" or "Jaccard")
// Output: a distance matrix where D_i,j is the distance between
// sample_i and sample_j according to the given distance metric


func BetaDiversityMatrix(allMaps []map[string]int, distMetric string) [][]float64{
		// creates distanceMatrix
		distanceMatrix := InitializeMatrix(len(allMaps), len(allMaps))

		// iterates over the matrix, assinging every spot to the value of the distance metric.
		for i := 0; i < len(distanceMatrix); i++ {
			for j := 0; j < len(distanceMatrix); j++ {
				if distMetric == "Bray-Curtis" {
					distanceMatrix[i][j] = BrayCurtis.BrayCurtisDistance(allMaps[i], allMaps[j])
				} else {
					distanceMatrix[i][j] = Jaccard.JaccardDistance(allMaps[i],allMaps[j])
				}
			}
		}
		return distanceMatrix
}

func InitializeMatrix(numRows, numCols int) [][]float64 {
	myMatrix := make([][]float64, numRows)
	for i := range myMatrix {
		myMatrix[i] = make([]float64, numCols)
	}
	return myMatrix
}
