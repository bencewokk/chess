package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//piecrea-piece creation

type popup struct {
	color string
	ty    string
}

const (
	n  = 8
	nw = 7
	ne = 9

	w = -1
	e = 1

	s  = -8
	sw = -7
	se = -9
)

func createPopUp(color, ty string) popup {
	return popup{color: color, ty: ty}
}

func checkWayRook(firstPos, curPos pos) string {

	offset := (PosToId(firstPos) - PosToId(curPos))

	if (PosToId(firstPos)-PosToId(curPos))%8 == 0 {
		if PosToId(firstPos)-PosToId(curPos) > 0 {
			return "n"
		} else {
			return "s"
		}
	} else {
		if offset > 0 {
			return "w"
		} else {
			return "e"
		}
	}

	//log.Println("checkWayRook: should not reach this point")

}

func checkCollision(way string, distance int) bool {
	switch way {
	case "n":

		for i := 0; i < distance; i++ {
			fmt.Println(PosToId(firstPos) + i*8)
			if tablep[PosToId(firstPos)+i*8].img == nil {

				fmt.Println("kurvanayad4")
				return false

			}
		}
	}

	return true //should not reach this point

}

func isInSameRow(curPos, firstPos pos) bool {
	if PosToId(curPos)/8 == PosToId(firstPos)/8 {
		return true
	} else {
		return false
	}

}

func isInLegalId(ids []int, id int) bool {
	for i := 0; i < len(ids); i++ {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func moveIsLegal(firstPos, curPos pos, screenat *ebiten.Image) bool {

	piece := tablep[PosToId(firstPos)]
	switch piece.piece {
	case "bp":
		bpids := []int{7, 9}

		//move forward & attack & replace pawn
		if PosToId(piece.pos) < 16 &&
			PosToId(firstPos)-PosToId(curPos) == 8 &&
			tablep[PosToId(curPos)].img == nil {

			return true
		}
		if PosToId(piece.pos) < 16 &&
			isInLegalId(bpids, PosToId(firstPos)-PosToId(curPos)) == true &&
			tablep[PosToId(curPos)].color == "w" {

			popups = append(popups, createPopUp("b", "piecrea"))

			return true
		}

		//move forward & attack
		if PosToId(piece.pos) > 15 &&
			PosToId(firstPos)-PosToId(curPos) == 8 &&
			tablep[PosToId(curPos)].img == nil {

			return true
		}
		if PosToId(piece.pos) > 15 &&
			isInLegalId(bpids, PosToId(firstPos)-PosToId(curPos)) == true &&
			tablep[PosToId(curPos)].color == "w" {

			return true
		}

	case "br":

		if (PosToId(curPos)-PosToId(firstPos))%-8 == 0 || isInSameRow(curPos, firstPos) == true { //is in same row
			if tablep[PosToId(curPos)].img == nil || tablep[PosToId(curPos)].color == "w" { //curpos isnt black
				if checkCollision(checkWayRook(firstPos, curPos), PosToId(firstPos)-PosToId(curPos)) == false {
					//fmt.Println(checkWayRook(firstPos, curPos))

					return true
				}
			}
		} else {
			fmt.Println("kurvanayad3")
		}

	case "bn":

	case "bb":

	case "bq":

	case "bk":

	default:
		// Unknown piece type, should not reach this point
		log.Println("moveIsLegal: unknown piece type")
	}
	return false
}
