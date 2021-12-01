package day1

import "fmt"

type Measurements = []int

func SlideMeasurements(measurements Measurements) (slided Measurements) {
	measurementLength := len(measurements)
	fmt.Println(measurementLength)

	for key, measurement := range measurements {
		if measurementLength <= key + 2 {
			break
		}

		total := measurement + measurements[key + 1] + measurements[key + 2]
		slided = append(slided, total)
	}

	return slided
}

func CountIncrements(measurements Measurements) int {
	count := 0

	for key, measurement := range measurements {
		if key == 0 {
			continue
		}

		prev := measurements[key - 1]

		if measurement > prev {
			count++
		}
	}

	return count
}
