package gomath

import "testing"

func TestNewMatrix(t *testing.T) {
	rows := []int{10, 20, 50, 1000, 10000}
	cols := []int{10, 21, 50, 1001, 10000}

	for i := 0; i < len(rows); i++ {
		expected := make([]float64, rows[i]*cols[i])
		m, err := NewMatrix(rows[i], cols[i])
		if m.data != expected {
			t.Error(
				"For", rows[i], cols[i],
				"expected", expected,
				"got", m, err,
			)
		}
	}

	rows = []int{10001, 10000, 10001, 10000}
	cols = []int{10001, 10000, 10000, 10001}
	expected := "Invalid dimensions (too big)"

	for i := 0; i < len(rows); i++ {
		expected := make([]float64, rows[i]*cols[i])
		m, err := NewMatrix(rows[i], cols[i])
		if err != expected {
			t.Error(
				"For", rows[i], cols[i],
				"expected", expected,
				"got", m, err,
			)
		}
	}

	rows = []int{0, -1, 20}
	cols = []int{-1, 0, -10}
	expected = "Invalid dimensions (negative)"

	for i := 0; i < len(rows); i++ {
		expected := make([]float64, rows[i]*cols[i])
		m, err := NewMatrix(rows[i], cols[i])
		if err != expected {
			t.Error(
				"For", rows[i], cols[i],
				"expected", expected,
				"got", m, err,
			)
		}
	}

	expected = "Invalid dimensions (both zero)"
	m, err := NewMatrix(0, 0)
	if err != expected {
		t.Error(
			"For", 0, 0,
			"expected", expected,
			"got", m, err,
		)
	}
}
