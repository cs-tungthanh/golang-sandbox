package utils_test

import (
	"strings"
	"testing"

	"github.com/cstungthanh/sandbox/src/utils"
)

func TestMap(t *testing.T) {
	s := []string{"a", "b", "c"}
	expected := []string{"A", "B", "C"}
	result := utils.Map(s, strings.ToUpper)
	if len(result) != len(expected) {
		t.Errorf("Expected length of result to be %d, but got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected result[%d] to be %s, but got %s", i, expected[i], result[i])
		}
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4}
	expected := []int{2, 4}
	result := utils.Filter(s, func(v int) bool {
		return v%2 == 0
	})
	if len(result) != len(expected) {
		t.Errorf("Expected length of result to be %d, but got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected result[%d] to be %d, but got %d", i, expected[i], result[i])
		}
	}
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4}
	expected := 10
	result := utils.Reduce(s, 0, func(cur, next int) int {
		return cur + next
	})
	if result != expected {
		t.Errorf("Expected result to be %d, but got %d", expected, result)
	}
}
