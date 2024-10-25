package controller

import (
	models "Sudoku/Creator/Models"
	repository "Sudoku/Creator/Repository"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type SudokuCollection struct {
	Collection *mongo.Collection
}

func (sc *SudokuCollection) GetSudokuWithNumber(Number int) [9][9]byte {
	board := [9][9]byte{}

	repo := repository.SudokuCollection{Sudoku: sc.Collection}

	sudoku := repo.GetSudokusByNumber(Number)

	digitPointer := 0
	xPointer := 0
	yPointer := 1
	xAxis := 0
	yAxis := 1
	digit := 0

	for digitPointer < len(sudoku.Digits) {

		digit, _ = strconv.Atoi(string(sudoku.Digits[digitPointer]))
		xAxis, _ = strconv.Atoi(string(sudoku.Location[xPointer]))
		yAxis, _ = strconv.Atoi(string(sudoku.Location[yPointer]))

		board[xAxis][yAxis] = byte(digit)

		xPointer += 2
		yPointer = xPointer + 1
		digitPointer++
	}

	for i := len(board) - 1; i >= 0; i-- {
		for j := len(board) - 1; j >= 0; j-- {
			if board[i][j] == 0 {
				board[i][j] = '.'
			}
		}
	}

	return board
}

func (sc *SudokuCollection) InsertSudoku(board [9][9]byte) error {
	Digits := ""
	Location := ""

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				Digits += strconv.Itoa(int(board[i][j]))
				Location += strconv.Itoa(i)
				Location += strconv.Itoa(j)
			}
		}
	}

	repo := repository.SudokuCollection{Sudoku: sc.Collection}

	var NewSudoku models.SudokuModel

	NewSudoku.ID = uuid.NewString()
	NewSudoku.Digits = Digits
	NewSudoku.Location = Location
	NewSudoku.Number = int32(repo.GetLastNumber()) + 1

	err := repo.InsertSudoku(&NewSudoku)
	if err != nil {
		return err
	}
	return nil
}

func (sc *SudokuCollection) DeleteSudokuById(ID string) (bool, error) {

	repo := repository.SudokuCollection{Sudoku: sc.Collection}
	res, err := repo.DeleteSudokuById(ID)

	if err != nil {
		return false, err
	}
	return res, nil
}

func (sc *SudokuCollection) DeleteSudokuByNumber(Number int) (bool, error) {
	repo := repository.SudokuCollection{Sudoku: sc.Collection}

	res, err := repo.DeleteSudokuWithNumber(Number)

	if err != nil {
		return false, err
	}
	return res, nil
}
