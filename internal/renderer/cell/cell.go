package cell

import (
	"image/color"

	"github.com/goncharovnikita/2048-go/internal/system/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Renderer interface {
	Render(to *ebiten.Image)
	Update(size board.CellSize)
}

type renderer struct {
	img   *ebiten.Image
	font  font.Face
	color color.Color

	text      string
	textColor color.Color

	cellSize board.CellSize

	updated bool

	xOffset int
	yOffset int
}

func New(
	size board.CellSize,
	font font.Face,
) *renderer {
	rdr := &renderer{
		font:      font,
		color:     sizeToColor(size),
		text:      sizeToText(size),
		textColor: color.RGBA{119, 110, 101, 255},
		cellSize:  size,
		updated:   true,
	}

	rdr.updateOffsets()

	return rdr
}

func (c *renderer) Render(to *ebiten.Image) {
	if !c.updated {
		return
	}

	h, w := to.Size()

	if c.img == nil {
		img := ebiten.NewImage(h, w)

		c.img = img
	}

	c.img.Fill(c.color)

	x := (w / 2) - c.xOffset
	y := (h / 2) + c.yOffset

	text.Draw(c.img, c.text, c.font, x, y, c.textColor)

	to.DrawImage(c.img, nil)

	c.updated = false
}

func (c *renderer) Update(size board.CellSize) {
	if c.cellSize == size {
		return
	}

	c.cellSize = size
	c.text = sizeToText(size)
	c.color = sizeToColor(size)

	c.updateOffsets()

	c.updated = true
}

func (c *renderer) updateOffsets() {
	rec := text.BoundString(c.font, c.text)

	c.yOffset = rec.Dy() / 2
	c.xOffset = rec.Dx() / 2
}

func sizeToColor(s board.CellSize) color.Color {
	switch s {
	case board.CS2:
		return color.RGBA{238, 228, 218, 255}
	case board.CS4:
		return color.RGBA{237, 224, 200, 255}
	case board.CS8:
		return color.RGBA{242, 177, 121, 255}
	case board.CS16:
		return color.RGBA{245, 149, 99, 255}
	case board.CS32:
		return color.RGBA{246, 124, 95, 255}
	case board.CS64:
		return color.RGBA{246, 94, 59, 255}
	case board.CS128:
		return color.RGBA{237, 207, 114, 255}
	case board.CS256:
		return color.RGBA{237, 204, 97, 255}
	case board.CS512:
		return color.RGBA{237, 200, 80, 255}
	case board.CS1024:
		return color.RGBA{237, 197, 63, 255}
	case board.CS2048:
		return color.RGBA{237, 194, 46, 255}

	default:
		return color.RGBA{238, 228, 218, 200}
	}
}

func sizeToText(s board.CellSize) string {
	switch s {
	case board.CS2:
		return "2"
	case board.CS4:
		return "4"
	case board.CS8:
		return "8"
	case board.CS16:
		return "16"
	case board.CS32:
		return "32"
	case board.CS64:
		return "64"
	case board.CS128:
		return "128"
	case board.CS256:
		return "256"
	case board.CS512:
		return "512"
	case board.CS1024:
		return "1024"
	case board.CS2048:
		return "2048"

	default:
		return ""
	}
}
