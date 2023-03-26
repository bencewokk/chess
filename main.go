package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// var bbishop, bking, bknight, bpawn, bqueen, brook, wbishop, wking, wknight, wpawn, wqueen, wrook *ebiten.Image
var img *ebiten.Image

/*
rook 1
knight 2
bishops 3
king 4
queen 5
pawn 6
/

	func init() {
		var err error
		wknight, _, err = ebitenutil.NewImageFromFile("wknight.png")
		if err != nil {
			log.Fatal(err)
		}
		bknight, _, err = ebitenutil.NewImageFromFile("bknight.png")
		if err != nil {
			log.Fatal(err)
		}
		bking, _, err = ebitenutil.NewImageFromFile("bking.png")
		if err != nil {
			log.Fatal(err)
		}
		bbishop, _, err = ebitenutil.NewImageFromFile("bbishop.png")
		if err != nil {
			log.Fatal(err)
		}
		bpawn, _, err = ebitenutil.NewImageFromFile("bpawn.png")
		if err != nil {
			log.Fatal(err)
		}
		bqueen, _, err = ebitenutil.NewImageFromFile("bqueen.png")
		if err != nil {
			log.Fatal(err)
		}
		brook, _, err = ebitenutil.NewImageFromFile("brook.png")
		if err != nil {
			log.Fatal(err)
		}
		wbishop, _, err = ebitenutil.NewImageFromFile("wbishop.png")
		if err != nil {
			log.Fatal(err)
		}
		wking, _, err = ebitenutil.NewImageFromFile("wking.png")
		if err != nil {
			log.Fatal(err)
		}
		wpawn, _, err = ebitenutil.NewImageFromFile("wpawn.png")
		if err != nil {
			log.Fatal(err)
		}
		wqueen, _, err = ebitenutil.NewImageFromFile("wqueen.png")
		if err != nil {
			log.Fatal(err)
		}
		wrook, _, err = ebitenutil.NewImageFromFile("wrook.png")
		if err != nil {
			log.Fatal(err)
		}
	}
*/
type Game struct{}

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

	fen := []string{
		"wr", "wn", "wb", "wq", "wk", "wb", "wn", "wr",
		"wp", "wp", "wp", "wp", "wp", "wp", "wp", "wp",
		"8",
		"8",
		"8",
		"8",
		"bp", "bp", "bp", "bp", "bp", "bp", "bp", "bp",
		"br", "bn", "bb", "bq", "bk", "bb", "bn", "br"}

	var boardindex int
	for i := 0; i < len(fen); i++ {
		if fen[i] == "wr" {
			board[boardindex] = "wr"
		} else if fen[i] == "wn" {
			board[boardindex] = "wn"
		} else if fen[i] == "wb" {
			board[boardindex] = "wb"
		} else if fen[i] == "wk" {
			board[boardindex] = "wk"
		} else if fen[i] == "wq" {
			board[boardindex] = "wq"
		} else if fen[i] == "wp" {
			board[boardindex] = "wp"

		} else if fen[i] == "br" {
			board[boardindex] = "br"
		} else if fen[i] == "bn" {
			board[boardindex] = "bn"
		} else if fen[i] == "bb" {
			board[boardindex] = "bb"
		} else if fen[i] == "bk" {
			board[boardindex] = "bk"
		} else if fen[i] == "bq" {
			board[boardindex] = "bq"
		} else if fen[i] == "bp" {
			board[boardindex] = "bp"

		} else if fen[i] == "8" {
			boardindex += 7
		} else if fen[i] == "7" {
			boardindex += 6
		} else if fen[i] == "6" {
			boardindex += 5
		} else if fen[i] == "5" {
			boardindex += 4
		} else if fen[i] == "4" {
			boardindex += 3
		} else if fen[i] == "3" {
			boardindex += 2
		} else if fen[i] == "2" {
			boardindex += 1
		}
		boardindex++
	}

	for i := 0; i != 0; i++ {

		if i%8 == 0 && i != 0 {
			fmt.Print("\n")
		}
		fmt.Print(board[i])

	}
	//fmt.Println(board)
}

func getPiece(board []string, i int) string {
	returnstr := ""

	if board[i] == "wr" {
		returnstr = "wrook.png"
	} else if board[i] == "wk" {
		returnstr = "wknight.png"
	} else if board[i] == "wb" {
		returnstr = "wbishop.png"
	} else if board[i] == "wk" {
		returnstr = "wking.png"
	} else if board[i] == "wq" {
		returnstr = "wqueen.png"
	} else if board[i] == "wp" {
		returnstr = "wpawn.png"

	} else if board[i] == "br" {
		returnstr = "brook.png"
	} else if board[i] == "bk" {
		returnstr = "bknight.png"
	} else if board[i] == "bb" {
		returnstr = "bbishop.png"
	} else if board[i] == "bk" {
		returnstr = "bking.png"
	} else if board[i] == "bq" {
		returnstr = "bqueen.png"
	} else if board[i] == "bp" {
		returnstr = "bpawn.png"
	}

	return returnstr
}

func (g *Game) Draw(screen *ebiten.Image) {
	res := float32(100)
	boardMade := 0
	board := make([]string, 64)
	if boardMade != 1 {
		fillBoard(board)
		boardMade++
	}

	fmt.Println(board)

	greenColor := color.RGBA{200, 255, 200, 255}
	offwhiteColor := color.RGBA{255, 248, 220, 255}
	multiplier := 1
	row := 0
	var piece string
	var x, y float32

	var opts ebiten.DrawImageOptions

	for i := 0; i < 64; i++ {
		//fmt.Println(i, "---")
		//fmt.Println((i + row) % 2)
		piece = getPiece(board, i)
		//opts.GeoM.Translate(float64(x), float64(y))
		opts.GeoM.Translate(float64(x), float64(y))
		fmt.Println(y, x)
		var err error

		if piece != "" {
			img, _, err = ebitenutil.NewImageFromFile("wpawn.png")
			if err != nil {
				log.Fatal(err)
			}
			screen.DrawImage(img, &opts)
		}
		if (i+row)%2 == 0 {
			vector.DrawFilledRect(screen, x, y, res, res, offwhiteColor, false)
		} else {
			vector.DrawFilledRect(screen, x, y, res, res, greenColor, false)
		}

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
