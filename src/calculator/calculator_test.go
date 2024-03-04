package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateSimilarRate_AsExpected(t *testing.T) {
	basedText := "I want to go to the zoo"
	inputText := "I go to the zoo"
	result := CalculateSimilarRate(basedText, inputText)
	require.NotZero(t, result)
	require.Less(t, result, 1.0)
}

func TestCalculateSimilarRate_EmptyInput_AsExpected(t *testing.T) {
	basedText := ""
	inputText := "hello"
	result := CalculateSimilarRate(basedText, inputText)
	require.Zero(t, result)
}

func TestCalculateSimilarRateWithNormalize_UnrelatedWords(t *testing.T) {
	basedText := "I want to go to the zoo"
	inputText := "I am a developer"
	result := CalculateSimilarRate(basedText, inputText)
	require.Less(t, result, 0.5)
}

func TestCalculateSimilarRateWithNormalize_AsExpected(t *testing.T) {
	basedText := "I want to go to the zoo"
	inputText := "I go to the zoo"
	result := CalculateSimilarRateWithNormalize(basedText, inputText)
	require.NotZero(t, result)
	require.Less(t, result, 1.0)
}

func TestCalculateSimilarRateWithNormalize_EmptyInput_AsExpected(t *testing.T) {
	basedText := ""
	inputText := "hello"
	result := CalculateSimilarRateWithNormalize(basedText, inputText)
	require.Zero(t, result)
}
