package core

import "math"

// L2Normalize normalizes a sparse vector using L2-normalization
func L2Normalize(vec map[string]int) map[string]float64 {
	normalizedVec := make(map[string]float64)
	magnitude := 0.0

	for _, count := range vec {
		magnitude += float64(count * count)
	}
	magnitude = math.Sqrt(magnitude)

	for word, count := range vec {
		normalizedVec[word] = float64(count) / magnitude
	}
	return normalizedVec
}
