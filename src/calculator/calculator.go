package calculator

import "github.com/thvinhtruong/similar-rate-text/src/core"

const (
	// Timeout for the function
	Timeout = 1000
)

// CalculateSimilarRate calculates the similarity rate between two texts using trivial cosine similarity method
func CalculateSimilarRate(basedText string, inputText string) float64 {
	baseVector := core.VectorizeText(basedText)
	inputVector := core.VectorizeText(inputText)

	return core.CosineSimilarity(baseVector, inputVector)
}

// CalculateSimilarRateWithNormalize calculates the similarity rate between two texts using cosine similarity method with L2 normalization
func CalculateSimilarRateWithNormalize(basedText string, inputText string) float64 {
	basedVector := core.L2Normalize(core.VectorizeText(basedText))
	inputVector := core.L2Normalize(core.VectorizeText(inputText))

	return core.CosineSimilarityWithNormalization(basedVector, inputVector)
}

// Async version of CalculateSimilarRate to save time for main thread
func CalculateSimilarRateAsync(basedText string, inputText string, allowNormalize bool, resultChan chan float64) {
	go func() {
		if allowNormalize {
			resultChan <- CalculateSimilarRateWithNormalize(basedText, inputText)
			return
		}
		resultChan <- CalculateSimilarRate(basedText, inputText)
	}()

	// Timeout
	go func() {
		resultChan <- 0
	}()
}
