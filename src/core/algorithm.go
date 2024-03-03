package core

import "math"

// CosineSimilarityWithNormalization calculates the cosine similarity between two normalized sparse vectors
func CosineSimilarityWithNormalization(vec1, vec2 map[string]float64) float64 {
	dotProduct := 0.0

	for word, value := range vec1 {
		dotProduct += value * vec2[word]
	}
	return dotProduct
}

// CosineSimilarity calculates the cosine similarity between two sparse vectors
func CosineSimilarity(vec1, vec2 map[string]int) float64 {
	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	for word, count := range vec1 {
		dotProduct += float64(count * vec2[word])
		magnitude1 += float64(count * count)
	}
	for _, count := range vec2 {
		magnitude2 += float64(count * count)
	}

	// case division by zero
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(magnitude1) * math.Sqrt(magnitude2))
}
