package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(MinimumSkew(text))
	for i := range "Gello" {
		fmt.Println(i)
	}
}

// FrequencyMap : given a text and length of a substring, returns a map of key=substrings with value being their frequencies
func FrequencyMap(text string, k int) map[string]int {
	freq := make(map[string]int)
	for j := 0; j < len(text)-k+1; j++ {
		pattern := text[j : j+k]
		freq[pattern]++
	}
	return freq
}

// MaxValue : given a frequency map, returns the maximum frequency
func MaxValue(freqMap map[string]int) int {
	m := 0
	firstTimeThrough := true

	for _, value := range freqMap {
		if firstTimeThrough || value > m {
			m = value
			firstTimeThrough = false
		}
	}
	return m
}

// MinArray : given an array of integers, returns the minimum value
func MinArray(nums []int) int {
	starter := nums[0]
	for i := range nums {
		if nums[i] < starter {
			starter = nums[i]
		}
	}
	return starter
}

// Contains : given an array of texts and a certain item, returns whether the array contains the item
func Contains(texts []string, item string) bool {
	for i := range texts {
		if texts[i] == item {
			return true
		}
	}
	return false
}

// FindFrequentWords : given a string text and the length of a k-mer, returns an array containg the most frequent k-mers
func FindFrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)
	freqMap := FrequencyMap(text, k)
	max := MaxValue(freqMap)
	for pattern, val := range freqMap {
		if val == max {
			freqPatterns = append(freqPatterns, pattern)
		}
	}
	return freqPatterns
}

// StartingIndices : given a patterna and a text, returns an array of indices of occurences of pattern in text
func StartingIndices(text, pattern string) []int {
	positions := make([]int, 0)
	n := len(text)
	k := len(pattern)
	for i := 0; i <= n-k; i++ {
		if text[i:i+k] == pattern {
			positions = append(positions, i)
		}
	}
	return positions
}

// FindClumps : given a text, a length of a k-mer, a window length, and a frequency threshold -> returns nucleotide sequences that are clumped together
func FindClumps(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	n := len(text)
	for i := 0; i <= n-L; i++ {
		window := text[i : i+L]
		freqMap := FrequencyMap(window, k)
		for key, value := range freqMap {
			if value >= t && !Contains(patterns, key) {
				patterns = append(patterns, key)
			}
		}
	}
	return patterns
}

// Skew : given a symbol, returns an integer adhering to the idea of the difference between Guanine and Cytosine.
func Skew(symbol string) int {
	if symbol == "G" {
		return 1
	} else if symbol == "C" {
		return -1
	}
	return 0
}

// SkewArray : given a genome, implements the skew method over an array and returns the array.
func SkewArray(genome string) []int {
	n := len(genome)
	array := make([]int, n+1)
	array[0] = 0
	for i := 1; i <= n; i++ {
		array[i] = array[i-1] + Skew(string(genome[i-1]))
	}
	return array
}

// MinimumSkew : given a genome, returns an array of the particular indices of the genome where the skew hits the global minimum.
func MinimumSkew(genome string) []int {
	indices := make([]int, 0)
	n := len(genome)
	array := SkewArray(genome)
	m := MinArray(array)
	for i := 0; i <= n; i++ {
		if array[i] == m {
			indices = append(indices, i)
		}
	}
	return indices
}
