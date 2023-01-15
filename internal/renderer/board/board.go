package board

import (
	cellftr "github.com/goncharovnikita/2048-go/internal/factory/cell"
	cellrdr "github.com/goncharovnikita/2048-go/internal/renderer/cell"
	"github.com/goncharovnikita/2048-go/internal/system/board"
	"github.com/hajimehoshi/ebiten/v2"
)

type cell struct {
	img  *ebiten.Image
	opts *ebiten.DrawImageOptions

	rdr cellrdr.Renderer
}

type renderer struct {
	cells [][]*cell

	updated bool
}

func New(
	gridSize int,
	cellSize int,
	padding int,
	cellFactory cellftr.Factory,
) *renderer {
	return &renderer{
		cells: makeCells(cellFactory, gridSize, cellSize, padding),
	}
}

func (r *renderer) Render(to *ebiten.Image) {
	if !r.updated {
		return
	}

	for _, row := range r.cells {
		for _, cell := range row {
			cell.rdr.Render(cell.img)
			to.DrawImage(cell.img, cell.opts)
		}
	}
}

func (r *renderer) Update(b board.Board) {
	rows := b.Rows()

	for i, row := range rows {
		for j, cell := range row {
			r.cells[i][j].rdr.Update(cell.Size())
		}
	}

	r.updated = true
}

func makeCells(
	cf cellftr.Factory,
	gridSize int,
	cellSize int,
	padding int,
) [][]*cell {
	res := make([][]*cell, gridSize)

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			img := ebiten.NewImage(cellSize, cellSize)
			geom := ebiten.GeoM{}

			tx := float64(padding + (padding * j) + (cellSize * j))
			ty := float64(padding + (padding * i) + (cellSize * i))

			geom.Translate(tx, ty)

			opts := &ebiten.DrawImageOptions{
				GeoM: geom,
			}

			res[i] = append(res[i], &cell{
				img:  img,
				opts: opts,
				rdr:  cf.Make(board.Empty),
			})
		}
	}

	return res
}
