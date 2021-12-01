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
