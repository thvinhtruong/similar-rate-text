package core

import (
	"fmt"
	"strings"
)

// VectorizeText converts text into a sparse vector representation
func VectorizeText(text string) map[string]int {
	vector := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		vector[word]++
	}
	fmt.Println(vector)
	return vector
}
