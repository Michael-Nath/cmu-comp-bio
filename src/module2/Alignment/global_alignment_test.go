package Alignment

import (
	"strconv"
	"testing"
)

type GlobalAlignmentInput struct {
	str1     string
	str2     string
	match    float64
	mismatch float64
	gap      float64
}

type globaltestpair struct {
	input    GlobalAlignmentInput
	solScore float64
}

var globalTests = []globaltestpair{
	{GlobalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		1.000},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		4.000},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		3.000},

	{GlobalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		4.000},

	{GlobalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		4.500},

	{GlobalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		9.000},

	{GlobalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		6.000}}

func TestGlobalAlignment(t *testing.T) {
	for _, pair := range globalTests {
		v := GlobalAlignment(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		score := computeScore(v, pair.input.match, pair.input.mismatch, pair.input.gap)
		if score != pair.solScore {
			t.Error(
				"For", pair.input,
				"expected alignment with score", strconv.FormatFloat(pair.solScore, 'f', 3, 64),
				"got", v, "with a score of", strconv.FormatFloat(score, 'f', 3, 64),
			)
		}
	}
}

func computeScore(alignment Alignment, match float64, mismatch float64, gap float64) float64 {
	score := 0.0
	str1 := alignment[0]
	str2 := alignment[1]
	gapC := "-"

	for i, c := range str1 {
		if string(c) == string(str2[i]) {
			score += match
		} else if string(c) == gapC || string(str2[i]) == gapC {
			score += gap
		} else {
			score += mismatch
		}
	}
	return score
}
