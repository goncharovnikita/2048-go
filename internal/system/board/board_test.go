package board

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_moveLeft(t *testing.T) {
	cases := []struct {
		name     string
		row      [][]*cell
		expected [][]*cell
	}{
		{
			name:     "basic shift",
			row:      makeRow(Empty, CS2, CS4, CS8),
			expected: makeRow(CS2, CS4, CS8, Empty),
		},
		{
			name:     "shift with gaps",
			row:      makeRow(Empty, CS2, Empty, CS4, CS8),
			expected: makeRow(CS2, CS4, CS8, Empty, Empty),
		},
		{
			name:     "single merge",
			row:      makeRow(CS2, CS2, CS8),
			expected: makeRow(CS4, CS8, Empty),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			moveLeft(tc.row[0])

			require.Equal(t, debugRows(tc.expected), debugRows(tc.row))
		})
	}
}

func Test_moveRight(t *testing.T) {
	cases := []struct {
		name     string
		row      [][]*cell
		expected [][]*cell
	}{
		{
			name:     "basic shift",
			row:      makeRow(CS2, CS4, CS8, Empty),
			expected: makeRow(Empty, CS2, CS4, CS8),
		},
		{
			name:     "shift with gaps",
			row:      makeRow(CS2, CS4, CS8, Empty, Empty),
			expected: makeRow(Empty, Empty, CS2, CS4, CS8),
		},
		{
			name:     "single merge",
			row:      makeRow(CS2, CS2, CS8),
			expected: makeRow(Empty, CS4, CS8),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			moveRight(tc.row[0])

			require.Equal(t, debugRows(tc.expected), debugRows(tc.row))
		})
	}
}

func Test_moveUp(t *testing.T) {
	cases := []struct {
		name     string
		row      [][]*cell
		expected [][]*cell
	}{
		{
			name:     "basic shift",
			row:      makeColumn(Empty, CS2, CS4, CS8),
			expected: makeColumn(CS2, CS4, CS8, Empty),
		},
		{
			name:     "shift with gaps",
			row:      makeColumn(Empty, CS2, Empty, CS4, CS8),
			expected: makeColumn(CS2, CS4, CS8, Empty, Empty),
		},
		{
			name:     "single merge",
			row:      makeColumn(Empty, CS2, Empty, CS2, CS8),
			expected: makeColumn(CS4, CS8, Empty, Empty, Empty),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			moveUp(tc.row, 0)

			require.Equal(t, debugRows(tc.expected), debugRows(tc.row))
		})
	}
}

func Test_moveDown(t *testing.T) {
	cases := []struct {
		name     string
		row      [][]*cell
		expected [][]*cell
	}{
		{
			name:     "basic shift",
			row:      makeColumn(CS2, CS4, CS8, Empty),
			expected: makeColumn(Empty, CS2, CS4, CS8),
		},
		{
			name:     "shift with gaps",
			row:      makeColumn(CS2, CS4, CS8, Empty, Empty),
			expected: makeColumn(Empty, Empty, CS2, CS4, CS8),
		},
		{
			name:     "single merge",
			row:      makeColumn(CS2, CS2, CS8, Empty, Empty),
			expected: makeColumn(Empty, Empty, Empty, CS4, CS8),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			moveDown(tc.row, 0)

			require.Equal(t, debugRows(tc.expected), debugRows(tc.row))
		})
	}
}

func makeRow(sizes ...CellSize) [][]*cell {
	row := make([]*cell, 0, len(sizes))
	res := make([][]*cell, 0, 1)

	for _, size := range sizes {
		row = append(row, &cell{size: size})
	}

	res = append(res, row)

	return res
}

func makeColumn(sizes ...CellSize) [][]*cell {
	res := make([][]*cell, 0, len(sizes))

	for _, size := range sizes {
		res = append(res, []*cell{{size: size}})
	}

	return res
}

func debugRows(rows [][]*cell) string {
	parts := make([]string, 0)

	for _, row := range rows {
		rr := make([]string, 0)

		for _, cell := range row {
			rr = append(rr, fmt.Sprintf("(%s)", cell.Size().Stirng()))
		}

		parts = append(parts, strings.Join(rr, "-"))
	}

	return strings.Join(parts, "\n")
}
