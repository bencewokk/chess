package main

import (
	"fmt"
	"log"
)

//piecrea-piece creation

type popup struct {
	color string
	ty    string
}

/*const (
	n  = 8
	nw = 7
	ne = 9

	w = -1
	e = 1

	s  = -8
	sw = -7
	se = -9
)*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func noCollisionRook(curPos, firstPos pos) bool {
	idCur := PosToId(curPos)
	idFirst := PosToId(firstPos)

	rowCur, colCur := idCur/8, idCur%8
	rowFirst, colFirst := idFirst/8, idFirst%8

	if colCur == colFirst {
		for row := min(rowCur, rowFirst) + 1; row < max(rowCur, rowFirst); row++ {
			if tablep[row*8+colCur].img != nil {
				return false
			}
		}
	} else if rowCur == rowFirst {
		for col := min(colCur, colFirst) + 1; col < max(colCur, colFirst); col++ {
			if tablep[rowCur*8+col].img != nil {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkCollisionKnight(curPos, destPos pos) bool {
	// Convert current and destination positions to their corresponding indices
	idCur := PosToId(curPos)
	idDest := PosToId(destPos)

	// Calculate the row and column for current and destination positions
	rowCur, colCur := idCur/8, idCur%8
	rowDest, colDest := idDest/8, idDest%8

	// Calculate the difference between current and destination positions
	rowDiff := abs(rowCur - rowDest)
	colDiff := abs(colCur - colDest)

	// Check if the move is a valid L-shaped move for a knight
	if (rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2) {
		return true // Valid knight move without checking the target square
	}

	return false // Not a valid knight move
}

func checkCollisionBishop(curPos, destPos pos) bool {
	// Convert current and destination positions to their corresponding indices
	idCur := PosToId(curPos)
	idDest := PosToId(destPos)

	// Calculate the row and column for current and destination positions
	rowCur, colCur := idCur/8, idCur%8
	rowDest, colDest := idDest/8, idDest%8

	// Calculate the difference between current and destination positions
	rowDiff := abs(rowCur - rowDest)
	colDiff := abs(colCur - colDest)

	// Check if the move is a valid diagonal move for a bishop
	if rowDiff != colDiff {
		return false // Not a valid diagonal move
	}

	// Determine the direction of movement
	rowStep := 1
	if rowDest < rowCur {
		rowStep = -1
	}
	colStep := 1
	if colDest < colCur {
		colStep = -1
	}

	// Check all squares along the diagonal path
	row, col := rowCur+rowStep, colCur+colStep
	for row != rowDest && col != colDest {
		if tablep[row*8+col].img != nil {
			return false // Collision detected
		}
		row += rowStep
		col += colStep
	}

	return true // No collision detected, valid bishop move
}

func checkCollisionQueen(curPos, destPos pos) bool {
	// Convert current and destination positions to their corresponding indices
	idCur := PosToId(curPos)
	idDest := PosToId(destPos)

	// Calculate the row and column for current and destination positions
	rowCur, colCur := idCur/8, idCur%8
	rowDest, colDest := idDest/8, idDest%8

	// Calculate the difference between current and destination positions
	rowDiff := abs(rowCur - rowDest)
	colDiff := abs(colCur - colDest)

	// Check if the move is a valid horizontal/vertical move (like a rook)
	if rowCur == rowDest {
		// Horizontal move
		step := 1
		if colDest < colCur {
			step = -1
		}
		for col := colCur + step; col != colDest; col += step {
			if tablep[rowCur*8+col].img != nil {
				return false // Collision detected
			}
		}
		return true // No collision detected in horizontal path
	} else if colCur == colDest {
		// Vertical move
		step := 1
		if rowDest < rowCur {
			step = -1
		}
		for row := rowCur + step; row != rowDest; row += step {
			if tablep[row*8+colCur].img != nil {
				return false // Collision detected
			}
		}
		return true // No collision detected in vertical path
	} else if rowDiff == colDiff {
		// Diagonal move
		rowStep := 1
		if rowDest < rowCur {
			rowStep = -1
		}
		colStep := 1
		if colDest < colCur {
			colStep = -1
		}
		row, col := rowCur+rowStep, colCur+colStep
		for row != rowDest && col != colDest {
			if tablep[row*8+col].img != nil {
				return false // Collision detected
			}
			row += rowStep
			col += colStep
		}
		return true // No collision detected in diagonal path
	}

	return false // Not a valid queen move
}

func checkCollisionKing(curPos, destPos pos) bool {
	// Convert current and destination positions to their corresponding indices
	idCur := PosToId(curPos)
	idDest := PosToId(destPos)

	// Calculate the row and column for current and destination positions
	rowCur, colCur := idCur/8, idCur%8
	rowDest, colDest := idDest/8, idDest%8

	// Calculate the differences in rows and columns
	rowDiff := abs(rowCur - rowDest)
	colDiff := abs(colCur - colDest)

	// Check if the move is valid for a king (one square in any direction)
	if (rowDiff <= 1 && colDiff <= 1) && !(rowDiff == 0 && colDiff == 0) {
		// Check if the target square is occupied by another piece
		if tablep[idDest].img != nil {
			return false // Collision detected
		}
		return true // No collision detected; valid king move
	}

	return false // Not a valid king move
}

func isInSameRow(curPos, firstPos pos) bool {
	if PosToId(curPos)/8 == PosToId(firstPos)/8 {
		return true
	} else {
		return false
	}

}
func isInSameLine(curPos, firstPos pos) bool {
	if PosToId(curPos)%8 == PosToId(firstPos)%8 {
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

// curpos where the the pice is moving to
// firstpos where the pice is moving from
func moveIsLegal(firstPos, curPos pos) bool {
	piece := tablep[PosToId(firstPos)]
	switch piece.piece {
	case "bp": // Black Pawn
		blackPawnIds := []int{7, 9}
		if PosToId(piece.pos) < 16 && // Upgrading pawn
			isInLegalId(blackPawnIds, PosToId(firstPos)-PosToId(curPos)) &&
			tablep[PosToId(curPos)].color == "w" {
			fmt.Println("debug2")
			// TODO: add popups later on pawn promotion
			return true
		}

		if PosToId(piece.pos) > 15 && // Move forward without attacking
			PosToId(firstPos)-PosToId(curPos) == 8 &&
			tablep[PosToId(curPos)].img == nil {
			fmt.Println("debug3")
			return true
		}
		if PosToId(piece.pos) > 15 && // Move forward or attacking
			isInLegalId(blackPawnIds, PosToId(firstPos)-PosToId(curPos)) &&
			tablep[PosToId(curPos)].color == "w" {
			fmt.Println("debug4")
			return true
		}

	case "wp": // White Pawn
		whitePawnIds := []int{-7, -9}
		if PosToId(piece.pos) > 47 && // Upgrading pawn
			isInLegalId(whitePawnIds, PosToId(firstPos)-PosToId(curPos)) &&
			tablep[PosToId(curPos)].color == "b" {
			fmt.Println("debug5")
			// TODO: add popups later on pawn promotion
			return true
		}

		if PosToId(piece.pos) < 48 && // Move forward without attacking
			PosToId(firstPos)-PosToId(curPos) == -8 &&
			tablep[PosToId(curPos)].img == nil {
			fmt.Println("debug6")
			return true
		}
		if PosToId(piece.pos) < 48 && // Move forward or attacking
			isInLegalId(whitePawnIds, PosToId(firstPos)-PosToId(curPos)) &&
			tablep[PosToId(curPos)].color == "b" {
			fmt.Println("debug7")
			return true
		}

	case "br": // Black Rook
		if isInSameLine(curPos, firstPos) || isInSameRow(curPos, firstPos) {
			if noCollisionRook(curPos, firstPos) && tablep[PosToId(curPos)].color != "b" {
				// Moving without attacking -- target cell is empty
				return true
			}
		}

	case "wr": // White Rook
		if isInSameLine(curPos, firstPos) || isInSameRow(curPos, firstPos) {
			if noCollisionRook(curPos, firstPos) && tablep[PosToId(curPos)].color != "w" {
				// Moving without attacking -- target cell is empty
				return true
			}
		}

	case "bn": // Black Knight
		if checkCollisionKnight(firstPos, curPos) && tablep[PosToId(curPos)].color != "b" {
			return true
		}
	case "wn": // White Knight
		if checkCollisionKnight(firstPos, curPos) && tablep[PosToId(curPos)].color != "w" {
			return true
		}

	case "bb": // Black Bishop
		if checkCollisionBishop(firstPos, curPos) && tablep[PosToId(curPos)].color != "b" {
			return true
		}
	case "wb": // White Bishop
		if checkCollisionBishop(firstPos, curPos) && tablep[PosToId(curPos)].color != "w" {
			return true
		}

	case "bq": // Black Queen
		if checkCollisionQueen(firstPos, curPos) && tablep[PosToId(curPos)].color != "b" {
			return true
		}
	case "wq": // White Queen
		if checkCollisionQueen(firstPos, curPos) && tablep[PosToId(curPos)].color != "w" {
			return true
		}

	case "bk": // Black King
		if checkCollisionKing(firstPos, curPos) && tablep[PosToId(curPos)].color != "b" {
			return true
		}
	case "wk": // White King
		if checkCollisionKing(firstPos, curPos) && tablep[PosToId(curPos)].color != "w" {
			return true
		}

	default:
		// Unknown piece type, should not reach this point
		log.Println("moveIsLegal: unknown piece type")
	}
	return false
}
