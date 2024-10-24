package controller

import (
	models "Sudoku/Creator/Models"
	repository "Sudoku/Creator/Repository"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSudokuWithNumber(collection *mongo.Collection, Number int) [9][9]byte {
	board := [9][9]byte{}

	repo := repository.SudokuCollection{Sudoku: collection}

	sudoku := repo.GetSudokusByNumber(Number)

	digitPointer := 0
	xPointer := 0
	yPointer := 1
	xAxis := 0
	yAxis := 1

	digit := 0

	for digitPointer < len(sudoku.Digits) {

		digit = int(sudoku.Digits[digitPointer])
		xAxis = int(sudoku.Location[xPointer])
		yAxis = int(sudoku.Location[yPointer])

		board[xAxis][yAxis] = byte(digit)

		xPointer += 2
		yPointer += xPointer + 1
		digitPointer++
	}

	return board
}

func InsertSudoku(collection *mongo.Collection, board [9][9]byte) error {
	Digits := ""
	Location := ""

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				Digits += string(board[i][j])
				Location += strconv.Itoa(i)
				Location += strconv.Itoa(j)
			}
		}
	}

	repo := repository.SudokuCollection{Sudoku: collection}

	var NewSudoku models.SudokuModel

	NewSudoku.ID = uuid.NewString()
	NewSudoku.Digits = Digits
	NewSudoku.Location = Location
	NewSudoku.Number = int32(repo.GetLastNumber())

	err := repo.InsertSudoku(&NewSudoku)
	if err != nil {
		return err
	}
	return nil
}
