package gomath

import "testing"

func TestNewMatrix(t *testing.T) {
	rows := []int{10, 20, 50, 1000, 10000}
	cols := []int{10, 21, 50, 1001, 10000}

	for i := 0; i < len(rows); i++ {
		expected := make([]float64, rows[i]*cols[i])
		m, err := NewMatrix(rows[i], cols[i])
		if m.data == nil || len(m.data) != rows[i]*cols[i] {
			t.Error(
				"For", rows[i], cols[i],
				"expected", expected,
				"got", m, err,
			)
		}
	}

	// Checking for errors
	rows = []int{10001, 10001, 10000, 0, -1, 20}
	cols = []int{10001, 10000, 10001, -1, 0, -10}

	for i := 0; i < len(rows); i++ {
		m, err := NewMatrix(rows[i], cols[i])
		if err == nil {
			t.Error(
				"For", rows[i], cols[i],
				"expected an error",
				"got", m, err,
			)
		}
	}
}
