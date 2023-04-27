package valid_sudoku_start

import (
	"strconv"
)

const boardSize = 9

func isValidSudoku(board [][]byte) bool {
	return IsValidSudoku(board)
}

func IsValidSudoku(in [][]byte) bool {
	board := createBoard(in)

	return checkRows(board) && checkColumns(board) && checkSquares(board)
}

func checkSquares(board [][]*int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !checkSquare(board, i*3, j*3) {
				return false
			}
		}
	}
	return true
}

func checkSquare(board [][]*int, startingRow int, startingCol int) bool {
	incidence := make([]bool, 10)
	for i := startingRow; i < startingRow+3; i++ {
		for j := startingCol; j < startingCol+3; j++ {
			el := board[i][j]
			if el == nil {
				continue
			}
			if incidence[*el] {
				return false
			}
			incidence[*el] = true
		}
	}
	return true
}

func checkColumns(board [][]*int) bool {
	for i := range board {
		incidence := make([]bool, 10)
		for j := 0; j < boardSize; j++ {
			el := board[j][i]
			if el == nil {
				continue
			}
			if incidence[*el] {
				return false
			}
			incidence[*el] = true
		}
	}
	return true
}

func checkRows(board [][]*int) bool {
	for _, row := range board {
		incidence := make([]bool, 10)
		for _, el := range row {
			if el == nil {
				continue
			}
			if incidence[*el] {
				return false
			}
			incidence[*el] = true
		}
	}
	return true
}

func createBoard(oldBoard [][]byte) [][]*int {
	newBoard := make([][]*int, boardSize)
	for i, row := range oldBoard {
		newBoard[i] = make([]*int, boardSize)
		for j, el := range row {
			s := string(el)
			if s != "." {
				atoi, _ := strconv.Atoi(s)
				newBoard[i][j] = &atoi
			}
		}
	}
	return newBoard
}
