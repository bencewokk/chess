package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

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

func PosToId(pos pos) int {
	id := pos.x/100 + pos.y/100*8

	if id < 0 || id >= 64 {
		log.Panic("PosToId: Invalid pos")
	}

	return id
}

func IdToPos(id int) pos {
	if id < 0 || id >= 64 {
		log.Panic("IdToPos: Invalid id")
	}

	x := (id % 8) * 100
	y := (id / 8) * 100

	return pos{x, y}
}

func callPopUps(screen *ebiten.Image) {
	if len(popups) == 5 {
		log.Println("callPopUps: No popups")
	}

	for i := 0; i < len(popups); i++ {
		switch popups[i].ty {
		case "piecrea":
			if popups[i].color == "b" {
				vector.DrawFilledRect(screen, 50, 50, 200, 200, color.Black, false)
			}

		}
	}
}

var (
	firstPos     pos
	curPos       pos
	isFirstPress bool = true

	popups []popup
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

	//BOARD

	for i := 0; i < 64; i++ {

		if (row+i)%2 == 0 {

			vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), greenColor, false)
		} else {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), offwhiteColor, false)
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

	//CURSOR AND PIECES

	for i := 0; i < 64; i++ {

		xcurs, ycurs := ebiten.CursorPosition()
		curPos = createPos(xcurs, ycurs)
		pos := ebiten.DrawImageOptions{}

		//fmt.Println(ebiten.IsMouseButtonPressed(ebiten.MouseButton0)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) == true && isFirstPress == true {
			isFirstPress = false

			firstPos.x = xcurs
			firstPos.y = ycurs

		}

		if tablep[PosToId(curPos)].img != nil {
			//fmt.Println(tablep[PosToId(curPos)].piece)
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) == false &&
			isFirstPress == false {
			isFirstPress = true
			if moveIsLegal(firstPos, curPos, screen) == true {
				tablep[PosToId(firstPos)].pos = IdToPos(PosToId(curPos))
				tablep[PosToId(curPos)] = tablep[PosToId(firstPos)]
				tablep[PosToId(firstPos)] = piece{}
			}
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) == false {
			tablep[i].offsetX, tablep[i].offsetY = 0, 0
		} else {
			tablep[PosToId(firstPos)].offsetX, tablep[PosToId(firstPos)].offsetY =
				firstPos.x-xcurs, firstPos.y-ycurs
		}

		if tablep[i].img != nil {
			pos.GeoM.Translate(
				float64(tablep[i].pos.x-tablep[i].offsetX),
				float64(tablep[i].pos.y-tablep[i].offsetY))
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

		callPopUps(screen)
	}
}

func main() {
	game := &Game{}
	tablepp := &tablep
	*tablepp = fillTable(tableUnread)
	ebiten.SetWindowSize(800, 800)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
