package main

import "fmt"

func isInLegalId(ids []int, id int) bool {
	for i := 0; i < len(ids); i++ {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func moveIsLegal(firstPos, curPos pos) bool {

	piece := tablep[PosToId(firstPos)]
	switch piece.piece {
	case "bp":
		bpids := []int{7, 9}

		if PosToId(piece.pos) < 16 &&
			PosToId(firstPos)-PosToId(curPos) == 8 &&
			tablep[PosToId(curPos)].img == nil {

			return true
		}
		if PosToId(piece.pos) < 16 &&
			isInLegalId(bpids, PosToId(firstPos)-PosToId(curPos)) == true &&
			tablep[PosToId(curPos)].color == "w" {

			return true
		}

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

	case "bn":

	case "bb":

	case "bq":

	case "bk":

	default:
		// Unknown piece type, should not reach this point
		fmt.Println("moveIsLegal unknown piece type")
	}
	return false
}
