package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var img *ebiten.Image

/*
rook 1
knight 2
bishops 3
king 4
queen 5
pawn 6
*/
type rook struct {
	legalmoves []int
	pos        []int
}

type knight struct {
	legalmoves []int
	pos        []int
}

type bishop struct {
	legalmoves []int
	pos        []int
}

type king struct {
	checkpos   []int
	legalmoves []int
	pos        []int
}

type queen struct {
	legalmoves []int
	pos        []int
}

type pawn struct {
	legalmoves []int
	pos        []int
}

type Game struct {
	//img *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100 * 8, 100 * 8
}

/*
rook 1
night 2
bishops 3
king 4
queen 5
pawn 6
*/

func fillBoard(board []string) {

	fen := []string{"r", "n", "b", "q", "k", "b", "n", "r", "/", "p", "p", "p", "p",
		"p", "p", "p", "p", "/", "8", "/", "8", "/", "8", "/", "8", "/",
		"P", "P", "P", "P", "P", "P", "P", "P", "/", "R", "N", "K", "B", "Q", "K", "B", "N", "R"}

	for i := 0; i < len(fen); i++ {
		if fen[i] == "r" {
			board[i] = "r"
		} else if fen[i] == "n" {
			board[i] = "n"
		} else if fen[i] == "b" {
			board[i] = "b"
		} else if fen[i] == "k" {
			board[i] = "k"
		} else if fen[i] == "q" {
			board[i] = "q"
		}
	}
	fmt.Println(board)
}

func (g *Game) Draw(screen *ebiten.Image) {
	res := float32(100)

	board := make([]string, 64)
	fillBoard(board)

	fmt.Println(board)
	greenColor := color.RGBA{200, 255, 200, 255}
	offwhiteColor := color.RGBA{255, 248, 220, 255}
	var iColor color.Color
	multiplier := 1
	row := 0

	var x, y float32

	for i := 0; i < 64; i++ {
		//fmt.Println(i, "---")
		//fmt.Println((i + row) % 2)

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
	//fmt.Println("aaa")
	game := &Game{}

	ebiten.SetWindowSize(100*8, 100*8)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
