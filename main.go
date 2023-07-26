package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func say(str string) {
	fmt.Println(str)
}

type Game struct{}

type pos struct {
	x, y int
}

func createPos(x, y int) pos {
	return pos{x: x, y: y}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100 * 8, 100 * 8
}

var (
	firstPos     pos
	isFirstPress bool = true
)

func (g *Game) Draw(screen *ebiten.Image) {

	var (
		x float64
		y float64

		width  float64 = 100
		height float64 = 100

		row int

		greenColor    = color.RGBA{150, 190, 180, 255}
		offwhiteColor = color.RGBA{240, 220, 200, 255}
	)

	for i := 0; i < 64; i++ {
		xcurs, ycurs := ebiten.CursorPosition()
		pos := ebiten.DrawImageOptions{}

		if (row+i)%2 == 0 {
			ebitenutil.DrawRect(screen, x, y, width, height, greenColor)
		} else {
			ebitenutil.DrawRect(screen, x, y, width, height, offwhiteColor)
		}

		//fmt.Println(ebiten.IsMouseButtonPressed(ebiten.MouseButton0))
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) == true &&
			isFirstPress == true {
			isFirstPress = false

			firstPos.x = xcurs
			firstPos.y = ycurs

			say("pressed")
			fmt.Println(firstPos)
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) == false {
			isFirstPress = true
		} else {
			fmt.Println(firstPos.x-xcurs, firstPos.y-ycurs, firstPos)
		}

		if tablep[i].img != nil {
			pos.GeoM.Translate(float64(tablep[i].x), float64(tablep[i].y))
			screen.DrawImage(tablep[i].img, &pos)
		}

		x += 100

		if (i+1)%8 == 0 && i != 0 {
			x = 0
			y += 100

			if row == 0 {
				row++
			} else {
				row--
			}
		}
	}
}

func main() {
	game := &Game{}
	tablepp := &tablep

	*tablepp = fillTable(tableUnread)

	fmt.Println(tablep)

	ebiten.SetWindowSize(800, 800)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
