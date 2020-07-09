package main


//AverageOutDegree takes the adjacency list of an overlap network as input. It returns
//the average number of elements to which a given string is "connected"
//in the network.
func AverageOutDegree(adjList map[string][]string) float64 {
	sum := 0
	for _, val := range adjList {
		sum += len(val)
	}
	return float64(sum) / float64(len(adjList))
}
