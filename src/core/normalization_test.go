package core

import (
	"math"
	"reflect"
	"testing"
)

func TestL2NormalizeEmptyInputVector(t *testing.T) {
	inputVector1 := map[string]int{}
	expectedOutput1 := map[string]float64{}
	output1 := L2Normalize(inputVector1)
	if !reflect.DeepEqual(output1, expectedOutput1) {
		t.Errorf("L2Normalize(%v) = %v, expected %v", inputVector1, output1, expectedOutput1)
	}
}

func TestL2NormalizeNonEmptyInputVector(t *testing.T) {
	inputVector2 := map[string]int{"a": 2, "b": 3, "c": 4}
	expectedOutput2 := map[string]float64{"a": 2.0 / 5.385, "b": 3.0 / 5.385, "c": 4.0 / 5.385}
	output2 := L2Normalize(inputVector2)
	for key, value := range expectedOutput2 {
		if math.Abs(output2[key]-value) > 0.001 {
			t.Errorf("L2Normalize(%v) = %v, expected %v", inputVector2, output2, expectedOutput2)
			break
		}
	}
}
