package Alignment

type Alignment [2]string

//GlobalScoreTable takes two strings, along with a match reward and mismatch and gap
//alignment penalties. It returns a 2-D array
//holding dynamic programming scores for global alignment with these penalties.

//GlobalAlignment takes two strings, along with match reward, and mismatch, and gap penalties.
//It returns a maximum score global alignment of the strings corresponding to these penalties.
