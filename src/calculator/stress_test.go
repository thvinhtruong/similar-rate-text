package calculator

import "testing"

func BenchmarkCalculateSimilarRate(b *testing.B) {
	basedText := "I want to go to the zoo"
	inputText := "I go to the zoo"
	for i := 0; i < b.N; i++ {
		CalculateSimilarRate(basedText, inputText)
	}
}

func BenchmarkCalculateSimilarRateWithNormalize(b *testing.B) {
	basedText := "I want to go to the zoo"
	inputText := "I go to the zoo"
	for i := 0; i < b.N; i++ {
		CalculateSimilarRateWithNormalize(basedText, inputText)
	}
}

func BenchmarkCalculateSimilarRateAsync(b *testing.B) {
	basedText := "I want to go to the zoo"
	inputText := "I go to the zoo"
	for i := 0; i < b.N; i++ {
		resultChan := make(chan float64)
		CalculateSimilarRateAsync(basedText, inputText, true, resultChan)
	}
}
