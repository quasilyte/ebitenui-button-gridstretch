package main

import (
	"image"
	"os"

	uiimage "github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
)

func NineSliceImage(i *ebiten.Image, offsetX, offsetY int) *uiimage.NineSlice {
	size := i.Bounds().Size()
	w := size.X
	h := size.Y
	return uiimage.NewNineSlice(i,
		[3]int{offsetX, w - 2*offsetX, offsetX},
		[3]int{offsetY, h - 2*offsetY, offsetY},
	)
}

func LoadImage(filename string) *ebiten.Image {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
