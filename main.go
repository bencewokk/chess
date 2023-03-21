package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var img *ebiten.Image

type Game struct {
	img *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100 * 8, 100 * 8
}

func (g *Game) Draw(screen *ebiten.Image) {
	res := float32(100)

	greenColor := color.RGBA{200, 255, 200, 255}
	offwhiteColor := color.RGBA{255, 248, 220, 255}
	var iColor color.Color
	multiplier := 1
	row := 0

	var x, y float32

	for i := 0; i < 64; i++ {
		fmt.Println(i, "---")
		fmt.Println((i + row) % 2)

		if (i+row)%2 == 0 {
			iColor = offwhiteColor

		} else {
			iColor = greenColor

		}

		vector.DrawFilledRect(screen, x, y, res, res, iColor, false)

		x += res

		/*
			fmt.Println(i % 8)
			fmt.Println(x, y)
			fmt.Println(multiplier, "ff")
		*/

		if i == 8*multiplier-1 {
			x = 0
			row++
			y += res
			multiplier += 1
		}
	}
}

func main() {

	//pos

	img, _, err := ebitenutil.NewImageFromFile("redcar.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{img: img}

	ebiten.SetWindowSize(100*8, 100*8)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
