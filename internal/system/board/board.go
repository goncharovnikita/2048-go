package board

import (
	"fmt"
	"math/rand"
)

type CellSize int

const (
	Empty CellSize = 1 + iota
	CS2
	CS4
	CS8
	CS16
	CS32
	CS64
	CS128
	CS256
	CS512
	CS1024
	CS2048
)

func (s CellSize) Stirng() string {
	switch s {
	case Empty:
		return "x"
	case CS2:
		return "2"
	case CS4:
		return "4"
	case CS8:
		return "8"
	case CS16:
		return "16"
	case CS32:
		return "32"
	case CS64:
		return "64"
	case CS128:
		return "128"
	case CS256:
		return "256"
	case CS512:
		return "512"
	case CS1024:
		return "1024"
	case CS2048:
		return "2048"
	default:
		return "0"
	}

}

type Cell interface {
	Size() CellSize
	Empty() bool
}

type Board interface {
	Rows() [][]Cell
	HasSpace() bool
	Moved() bool
	GenerateCell()

	Left()
	Right()
	Up()
	Down()
}

type cell struct {
	size CellSize
}

type board struct {
	rows [][]*cell

	moved bool
}

func New(size int) *board {
	return &board{
		rows: makeCells(size, size, true),
	}
}

func (b *board) Rows() [][]Cell {
	res := make([][]Cell, 0, len(b.rows))

	for _, row := range b.rows {
		r := make([]Cell, 0, len(row))

		for _, cell := range row {
			r = append(r, cell)
		}

		res = append(res, r)
	}

	return res
}

func (b *board) HasSpace() bool {
	for _, row := range b.rows {
		for _, cell := range row {
			if cell.Empty() {
				return true
			}
		}

	}

	return false
}

func (b *board) Moved() bool {
	return b.moved
}

func (b *board) GenerateCell() {
	cellSize := CS2

	if rand.Intn(100) > 50 {
		cellSize = CS4
	}

	emptyCells := make([]*cell, 0)

	for _, row := range b.rows {
		for _, cell := range row {
			if cell.Empty() {
				emptyCells = append(emptyCells, cell)
			}
		}
	}

	emptyCells[rand.Intn(len(emptyCells))].size = cellSize
	b.moved = false
}

func (b *board) Left() {
	for _, row := range b.rows {
		if moveLeft(row) {
			b.moved = true
		}
	}
}

func (b *board) Right() {
	for _, row := range b.rows {
		if moveRight(row) {
			b.moved = true
		}
	}
}

func (b *board) Up() {
	row := b.rows[0]

	for i := range row {
		if moveUp(b.rows, i) {
			b.moved = true
		}
	}
}

func (b *board) Down() {
	row := b.rows[0]

	for i := range row {
		if moveDown(b.rows, i) {
			b.moved = true
		}
	}
}

func (b *board) Print() {
	rows := b.Rows()
	fmt.Printf("rows len: %d\n", len(rows))

	for _, row := range rows {
		fmt.Print("|")

		for _, cell := range row {
			fmt.Printf(" %d |", cell.Size())
		}

		fmt.Print("\n")
	}
}

func (c *cell) Size() CellSize {
	return c.size
}

func (c *cell) Empty() bool {
	return c.size == Empty
}

func canJoin(a, b Cell) bool {
	return !a.Empty() && !b.Empty() && a.Size() == b.Size()
}

func moveLeft(row []*cell) bool {
	tail := 0
	head := 1
	moved := false

	for head < len(row) {
		if canJoin(row[head], row[tail]) {
			mergeCells(row[head], row[tail])
			moved = true

			tail++
		}

		if !row[head].Empty() {
			tail = head
		}

		head++
	}

	tail = 0
	head = 1

	for head < len(row) {
		for !row[tail].Empty() && tail < head {
			tail++
		}

		if !row[head].Empty() && row[tail].Empty() {
			row[head], row[tail] = row[tail], row[head]
			moved = true
		}

		head++
	}

	return moved
}

func moveRight(row []*cell) bool {
	tail := len(row) - 1
	head := len(row) - 2
	moved := false

	for head >= 0 {
		if canJoin(row[head], row[tail]) {
			mergeCells(row[tail], row[head])
			moved = true

			tail--
		}

		if !row[head].Empty() {
			tail = head
		}

		head--
	}

	tail = len(row) - 1
	head = len(row) - 2

	for head >= 0 {
		for !row[tail].Empty() && tail > head {
			tail--
		}

		if !row[head].Empty() && row[tail].Empty() {
			row[head], row[tail] = row[tail], row[head]
			moved = true
		}

		head--
	}

	return moved
}

func moveUp(rows [][]*cell, row int) bool {
	tail := 0
	head := 1
	moved := false

	for head < len(rows) {
		a, b := rows[head][row], rows[tail][row]

		if canJoin(a, b) {
			mergeCells(b, a)
			moved = true

			tail++
		}

		if !rows[head][row].Empty() {
			tail = head
		}

		head++
	}

	tail = 0
	head = 1

	for head < len(rows) {
		for !rows[tail][row].Empty() && tail < head {
			tail++
		}

		if !rows[head][row].Empty() && rows[tail][row].Empty() {
			rows[head][row], rows[tail][row] = rows[tail][row], rows[head][row]
			moved = true
		}

		head++
	}

	return moved
}

func moveDown(rows [][]*cell, row int) bool {
	tail := len(rows) - 1
	head := len(rows) - 2
	moved := false

	for head >= 0 {
		a, b := rows[head][row], rows[tail][row]

		if canJoin(a, b) {
			mergeCells(b, a)
			moved = true

			tail--
		}

		if !rows[head][row].Empty() {
			tail = head
		}

		head--
	}

	tail = len(rows) - 1
	head = len(rows) - 2

	for head >= 0 {
		for !rows[tail][row].Empty() && tail > head {
			tail--
		}

		if !rows[head][row].Empty() && rows[tail][row].Empty() {
			rows[head][row], rows[tail][row] = rows[tail][row], rows[head][row]
			moved = true
		}

		head--
	}

	return moved
}

func makeCells(width, height int, withRandom bool) [][]*cell {
	if height == 0 || width == 0 {
		return nil
	}

	res := make([][]*cell, 0, height)

	for i := 0; i < height; i++ {
		row := make([]*cell, 0, width)

		for j := 0; j < width; j++ {
			s := Empty

			if withRandom {
				r := rand.Intn(100)
				if r > 85 {
					s = CS2
				} else if r > 75 {
					s = CS4
				}
			}

			row = append(row, &cell{
				size: s,
			})
		}

		res = append(res, row)
	}

	return res
}

func nextSize(s CellSize) CellSize {
	switch s {
	case CS2:
		return CS4

	case CS4:
		return CS8

	case CS8:
		return CS16

	case CS16:
		return CS32

	case CS32:
		return CS64

	case CS64:
		return CS128

	case CS128:
		return CS256

	case CS256:
		return CS512

	case CS512:
		return CS1024

	case CS1024:
		return CS2048

	case CS2048:
		return CellSize(0)

	default:
		return CellSize(0)
	}
}

func mergeCells(from, to *cell) {
	to.size = nextSize(to.size)
	from.size = Empty
}
