package main

import (
	"embed"
	"image/color"
	"log"
	"os"

	"github.com/goncharovnikita/2048-go/internal/asset"
	cellfactory "github.com/goncharovnikita/2048-go/internal/factory/cell"
	boardrenderer "github.com/goncharovnikita/2048-go/internal/renderer/board"
	"github.com/goncharovnikita/2048-go/internal/system/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	blockSize     = 80
	blocksPadding = 10
	gridSize      = 4
)

//go:embed assets/*
var gameAssets embed.FS

var (
	windowSize = (blocksPadding * (gridSize + 1)) + (blockSize * gridSize)
)

type renderer interface {
	Render(*ebiten.Image)
}

type boardUpdater interface {
	Update(board.Board)
}

type boardRenderer interface {
	renderer
	boardUpdater
}

type Game struct {
	boardRenderer boardRenderer
	boardSystem   board.Board

	bkg *ebiten.Image
}

func newGame(
	boardRenderer boardRenderer,
	boardSystem board.Board,
) *Game {
	return &Game{
		boardRenderer: boardRenderer,
		boardSystem:   boardSystem,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.boardSystem.Left()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		g.boardSystem.Right()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		g.boardSystem.Down()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		g.boardSystem.Up()
	}

	if g.boardSystem.Moved() {
		g.boardSystem.GenerateCell()
		g.boardRenderer.Update(g.boardSystem)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.bkg == nil {
		bkg := ebiten.NewImage(screen.Size())
		bkg.Fill(color.RGBA{187, 173, 160, 255})

		g.bkg = bkg
	}

	g.boardRenderer.Render(g.bkg)

	screen.DrawImage(g.bkg, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowSize, windowSize
}

func main() {
	boardSystem := board.New(gridSize)
	baseFont := asset.GetFont(gameAssets, 32, 72)
	cellFactory := cellfactory.New(baseFont)
	boardRenderer := boardrenderer.New(
		gridSize,
		blockSize,
		blocksPadding,
		cellFactory,
	)

	boardRenderer.Update(boardSystem)

	game := newGame(
		boardRenderer,
		boardSystem,
	)

	ebiten.SetWindowSize(windowSize, windowSize)
	ebiten.SetWindowTitle("2048")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
