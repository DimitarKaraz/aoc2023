package main

import "fmt"

var timeToDistanceRecords = map[int]int{
	// 7: 9,
	// 15: 40,
	// 30: 200,
	// 71530: 940200,
	34908986: 204171312101780,
	// 34: 204,
	// 90: 1713,
	// 89: 1210,
	// 86: 1780,
}

func main() {
	multiplicationResult := 1

	for time, _ := range timeToDistanceRecords {
		multiplicationResult *= calcWaysToBeatRecord(time)
	}

	fmt.Println(multiplicationResult)
}

func calcWaysToBeatRecord(raceTime int) int {
	waysToBeats := 0

	for speed := 0; speed <= raceTime; speed++ {
		if speed*(raceTime-speed) > timeToDistanceRecords[raceTime] {
			waysToBeats++
		}
	}
	fmt.Println(raceTime, " > ", waysToBeats)
	return waysToBeats
}
