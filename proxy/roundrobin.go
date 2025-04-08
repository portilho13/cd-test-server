package main

var roundRobinIndex *int = new(int) // Will set to 0 without creating complementary var


func RoundRobin(serversQueueSize int) int {
	if *roundRobinIndex >= serversQueueSize - 1 {
		*roundRobinIndex = 0
	} else {
		*roundRobinIndex += 1
	}
	return *roundRobinIndex
}