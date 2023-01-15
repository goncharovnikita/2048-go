package cell

import (
	"github.com/goncharovnikita/2048-go/internal/renderer/cell"
	"github.com/goncharovnikita/2048-go/internal/system/board"
	"golang.org/x/image/font"
)

type Factory interface {
	Make(size board.CellSize) cell.Renderer
}

type factory struct {
	font font.Face
}

func New(
	font font.Face,
) *factory {
	return &factory{
		font: font,
	}
}

func (f *factory) Make(size board.CellSize) cell.Renderer {
	return cell.New(size, f.font)
}
