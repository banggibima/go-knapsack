package main

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	capacity := 50

	result := knapsack(values, weights, capacity)

	expectedResult := 220

	if result != expectedResult {
		t.Errorf("Expected result: %d, but got: %d", expectedResult, result)
	}
}
