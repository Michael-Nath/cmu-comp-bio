package Alignment

type Alignment [2]string

// func main() {
// 	fmt.Println(GlobalAlignment("ATCGATCGT", "ATCGGCTAC", 1.0, 1.0, 0.5))
// }

//GlobalAlignment takes two strings, along with match reward, and mismatch, and gap penalties.
//It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str0, str1 string, match, mismatch, gap float64) Alignment {
	GlobalBacktrack := BackTrack(str0, str1, match, mismatch, gap)
	GAlign := OutputGlobal(str0, str1, GlobalBacktrack)
	return GAlign
}

// OutputGlobal :  Does Something
func OutputGlobal(str0, str1 string, backtrack [][]string) Alignment {
	//I'm going to chop off symbols as we go. When I run out of symbols in one string, I'm done
	newStr0 := ""
	newStr1 := ""
	for len(str0) > 0 || len(str1) > 0 {
		row := len(str0)
		col := len(str1)
		if backtrack[row][col] == "UP" {
			//chop off one symbol of str0
			newStr0 = string(str0[len(str0)-1]) + newStr0
			str0 = str0[:len(str0)-1]
			newStr1 = "-" + newStr1

		} else if backtrack[row][col] == "LEFT" {
			newStr1 = string(str1[len(str1)-1]) + newStr1
			str1 = str1[:len(str1)-1]
			newStr0 = "-" + newStr0
		} else if backtrack[row][col] == "DIAG" {
			//we could be at a match, or at a mismatch
			// if str0[len(str0)-1] == str1[len(str1)-1] {
			// 	//match
			// 	//add a symbol to the start of s
			// 	s = string(str0[len(str0)-1]) + s
			// }
			newStr0 = string(str0[len(str0)-1]) + newStr0
			newStr1 = string(str1[len(str1)-1]) + newStr1

			//now chop off a symbol of both str0 and str1
			str0 = str0[:len(str0)-1]
			str1 = str1[:len(str1)-1]
		} else {
			panic("Error: invalid value in backtrack array.")
		}
	}
	ans := Alignment{newStr0, newStr1}

	return ans
}

// BackTrack : backtracks
func BackTrack(str0, str1 string, match, mismatch, gap float64) [][]string {
	//we want a backtracking "pointer" for every node.
	numRows := len(str0) + 1
	numCols := len(str1) + 1

	//make the matrix
	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	//generate the scoring table using the function we already wrote
	scoringMatrix := GlobalScoreTable(str0, str1, match, mismatch, gap)

	//fill the matrix

	//first, the easiest backtracking pointers to set are the 0-th row and column
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}
	for i := 1; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	//traverse the rest of the table and set pointers based on neighboring values
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//check the value that was used to give me scoringMatrix[i][j]
			if scoringMatrix[i][j] == scoringMatrix[i-1][j]-gap {
				backtrack[i][j] = "UP"
			} else if scoringMatrix[i][j] == scoringMatrix[i][j-1]-gap {
				backtrack[i][j] = "LEFT"
			} else { // assume that the scoring matrix is correct
				backtrack[i][j] = "DIAG"
			}
		}
	}

	return backtrack
}

// TACG
// CACG-
