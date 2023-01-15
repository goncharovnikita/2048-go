package asset

import (
	"io"
	"io/fs"
	"log"
	"path"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type fileSystem interface {
	Open(name string) (fs.File, error)
}

func GetFont(
	fileSystem fileSystem,
	fontSize float64,
	dpi float64,
) font.Face {
	r, err := fileSystem.Open(path.Join("assets", "font.ttf"))
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	fontData, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	opts := &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingNone,
	}

	fontFace, err := opentype.NewFace(tt, opts)
	if err != nil {
		log.Fatal(err)
	}

	return fontFace
}
