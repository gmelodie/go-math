package gomath

import (
	"errors"
	"math"
)

const MAX_ROW_SIZE = 10000
const MAX_COL_SIZE = 10000

func NewMatrix(nrows, ncols int) (matrix, error) {
	m := matrix{nrows, ncols, nil}

	if nrows > MAX_ROW_SIZE || ncols > MAX_COL_SIZE {
		return m, errors.New("Invalid dimensions (too big)")
	}

	if nrows < 0 || ncols < 0 {
		return m, errors.New("Invalid dimensions (negative)")
	}

	if nrows == 0 && ncols == 0 {
		return m, errors.New("Invalid dimensions (both zero)")
	}

	m.data = make([]float64, nrows*ncols)
	return m, nil
}

type matrix struct {
	nrows int
	ncols int
	data  []float64
}

func (m *matrix) Get(row, col int) (float64, error) {
	if row >= m.nrows || col >= m.ncols || row < 0 || col < 0 {
		return 0, errors.New("Index out of bounds")
	}

	idx := m.ncols*row + col
	return m.data[idx], nil
}

func (m *matrix) Set(row, col int, val float64) error {
	if row >= m.nrows || col >= m.ncols || row < 0 || col < 0 {
		return errors.New("Index out of bounds")
	}

	idx := m.ncols*row + col
	m.data[idx] = val
	return nil
}

func (m *matrix) Det() (float64, error) {
	if m.nrows != m.ncols {
		return 0, errors.New("Not a square matrix")
	}

	dim := m.ncols // could be m.nrows since they are equal

	if dim <= 0 {
		return 0, errors.New("Invalid square matrix")
	}

	if dim == 1 {
		return m.data[0], nil
	}

	var det float64 = 0
	for i := 0; i < m.nrows; i++ {
		a, _ := m.Get(i, 0)
		det += a * m.coFactor(i, 0)
	}
	return det, nil
}

func (m *matrix) coFactor(row, col int) float64 {
	// build submatrix
	submatrix, _ := NewMatrix(m.nrows-1, m.ncols-1)
	for i := 0; i < m.nrows; i++ {
		if i == row {
			continue
		}
		for j := 0; j < m.ncols; j++ {
			if j == col {
				continue
			}
			val, _ := m.Get(i, j)
			submatrix.Set(i, j, val)
		}
	}

	d, _ := submatrix.Det()
	return math.Pow(float64(-1), float64(row+col)) * d
}
