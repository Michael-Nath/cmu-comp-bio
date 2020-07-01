package MultipleRichness

import("Metagenomics/Richness")

// Write the function RichnessMatrix() with the following input and output
// Input: an array of frequency maps samples
// Output: an slice whose i-th element is the richness of the i-th frequency map

func RichnessMatrix(allMaps []map[string]int) []float64 {
    RichnessSlice := make([]float64, len(allMaps))
    for i := 0; i < len(AllMaps); i++ {
			RichnessSlice[i] = Richness.Richness(AllMaps[i])
		}
    return RichnessSlice
}
