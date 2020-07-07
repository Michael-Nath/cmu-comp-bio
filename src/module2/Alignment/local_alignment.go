// Compass Quad 2020
package Alignment
//LocalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score local alignment of the strings corresponding to these penalties.

//LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for local alignment with these penalties.
// type Alignment [2]string
func LocalAlignment(str0, str1 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
  backtrack, end0, end1 := LocalBacktrack(str0, str1, match, mismatch, gap)
  ans := Alignment{"", ""}
  start0 := 0
  start1 := 0
  ans[0], ans[1], start0, start1 = OutputLocal(str0, str1, backtrack, end0, end1)
  return ans, start0, end0, start1, end1
}


func LocalScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {
  if len(str0) == 0 || len(str1) == 0 {
    panic("Error: String of Length Zero Given to LocalScoreTable")
  }
  numRows := len(str0) + 1
  numCols := len(str1) + 1
  scoringMatrix := make([][]float64, numRows)
	for i := range scoringMatrix {
		scoringMatrix[i] = make([]float64, numCols)
	}

  for i := 1; i < numRows; i++ {
    for j := 1; j < numCols; j++ {
      //Recurrence Relation
      top := scoringMatrix[i-1][j] - gap
      left := scoringMatrix[i][j-1] - gap
      diag := scoringMatrix[i-1][j-1]
      if str0[i-1] == str1[j-1] {
        diag += match
      } else {
        diag -= mismatch
      }
      scoringMatrix[i][j] = MaxFloat(top, left, diag, 0.0)
    }
  }
  return scoringMatrix
}


func LocalBacktrack(str0, str1 string, match, mismatch, gap float64) ([][]string, int, int) {
  numRows := len(str0) + 1
  numCols := len(str1) + 1
  scoringMatrix := LocalScoreTable(str0, str1, match, mismatch, gap)
  backtrack := make([][]string, numRows)
  end0 := 0
  end1 := 0
  max := scoringMatrix[0][0]
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}
  for j := 1; j < numCols; j++ {
		backtrack[0][j] = "HOME"
	}
  for i := 1; i < numRows; i++ {
    backtrack[i][0] = "HOME"
  }

  for i := 1; i < numRows; i++ {
    for j := 1; j < numCols; j++ {
      if scoringMatrix[i][j] > max {
        max = scoringMatrix[i][j]
        end0 = i
        end1 = j
      }
      if scoringMatrix[i][j] == 0 {
        backtrack[i][j] = "HOME"
      } else if scoringMatrix[i][j] == scoringMatrix[i-1][j] - gap {
        backtrack[i][j] = "UP"
      } else if scoringMatrix[i][j] == scoringMatrix[i][j-1] - gap {
        backtrack[i][j] = "LEFT"
      } else {
        backtrack[i][j] = "DIAG"
      }
    }
  }
  return backtrack, end0, end1
}

func OutputLocal(str0, str1 string, backtrack [][]string, end0, end1 int) (string, string, int, int) {
  start0 := end0
  start1 := end1
  newStr0 := ""
  newStr1 := ""
  for backtrack[start0][start1] != "HOME" && start0 >= 0 && start1 > 0 {
    ans := backtrack[start0][start1]
    if ans == "UP" {
      newStr1 = "-" + newStr1
      newStr0 = string(str0[start0-1]) + newStr0
      start0 -= 1
    } else if ans == "LEFT" {
      newStr0 = "-" + newStr0
      newStr1 = string(str1[start1-1]) + newStr1
      start1 -= 1
    } else if ans == "DIAG" {
      newStr0 = string(str0[start0-1]) + newStr0
      newStr1 = string(str1[start1-1]) + newStr1
      start0 -= 1
      start1 -= 1
    }
  }
  return newStr0, newStr1, start0, start1
}
