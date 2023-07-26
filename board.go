package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	tableUnread = []string{
		"wr", "wn", "wb", "wq", "wk", "wb", "wn", "wr",
		"wp", "wp", "wp", "wp", "wp", "wp", "wp", "wp",
		"8",
		"8",
		"8",
		"8",
		"bp", "bp", "bp", "bp", "bp", "bp", "bp", "bp",
		"br", "bn", "bb", "bq", "bk", "bb", "bn", "br"}

	tablep []piece

	wrook, _, _   = ebitenutil.NewImageFromFile("pieces/wrook.png")
	wqueen, _, _  = ebitenutil.NewImageFromFile("pieces/wqueen.png")
	wpawn, _, _   = ebitenutil.NewImageFromFile("pieces/wpawn.png")
	wknight, _, _ = ebitenutil.NewImageFromFile("pieces/wknight.png")
	wbishop, _, _ = ebitenutil.NewImageFromFile("pieces/wbishop.png")
	wking, _, _   = ebitenutil.NewImageFromFile("pieces/wking.png")

	brook, _, _   = ebitenutil.NewImageFromFile("pieces/brook.png")
	bqueen, _, _  = ebitenutil.NewImageFromFile("pieces/bqueen.png")
	bpawn, _, _   = ebitenutil.NewImageFromFile("pieces/bpawn.png")
	bknight, _, _ = ebitenutil.NewImageFromFile("pieces/bknight.png")
	bbishop, _, _ = ebitenutil.NewImageFromFile("pieces/bbishop.png")
	bking, _, _   = ebitenutil.NewImageFromFile("pieces/bking.png")
)

type piece struct {
	x, y     int
	piece    string
	img      *ebiten.Image
	legalIDs []int
}

func createPiece(pieceType string, img *ebiten.Image, x, y int) piece {
	return piece{piece: pieceType, img: img, x: x, y: y}
}

func fillTable(table []string) []piece {

	var (
		tablep = make([]piece, 64)
		x, y   int
	)

	for u, i := 0, 0; i < len(table); i++ {

		switch table[i] {
		case "1":
			u++
		case "2":
			u += 2
		case "3":
			u += 3
		case "4":
			u += 4
		case "5":
			u += 5
		case "6":
			u += 6
		case "7":
			u += 7
		case "8":
			u += 8

		case "wr":
			tablep[u] = createPiece("wr", wrook, x, y)
			u++
		case "wq":
			tablep[u] = createPiece("wq", wqueen, x, y)
			u++

		case "wp":
			tablep[u] = createPiece("wp", wpawn, x, y)
			u++

		case "wk":
			tablep[u] = createPiece("wk", wking, x, y)
			u++

		case "wn":
			tablep[u] = createPiece("wn", wknight, x, y)
			u++

		case "wb":
			tablep[u] = createPiece("wb", wbishop, x, y)
			u++

		case "br":
			tablep[u] = createPiece("br", brook, x, y)
			u++

		case "bq":
			tablep[u] = createPiece("bq", bqueen, x, y)
			u++

		case "bp":
			tablep[u] = createPiece("bp", bpawn, x, y)
			u++

		case "bk":
			tablep[u] = createPiece("bk", bking, x, y)
			u++

		case "bn":
			tablep[u] = createPiece("bn", bknight, x, y)
			u++

		case "bb":
			tablep[u] = createPiece("bb", bbishop, x, y)
			u++

		}
		var x2, y2 int
		for i := 0; i < 64; i++ {

			if tablep[i].img != nil {
				tablep[i].x = x2
				tablep[i].y = y2
			}

			x2 += 100

			if (i+1)%8 == 0 && i != 0 {
				x2 = 0
				y2 += 100
			}

		}

	}
	return tablep
}
