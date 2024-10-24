package creator

import (
	writer "Sudoku/Creator/Writer"
	"math/rand"
	"slices"
)

func Initialize(Level int) {
	Count := CountForLevels(Level)
	sudoku := Create(Count)

	i := 0
	for i < 1 {
		if Check(sudoku) {
			writer.WriteInFile(&sudoku)
			i++
		} else {
			sudoku = Create(Count)
		}
	}

}

func Create(Count int) [9][9]byte {
	sudoku := [9][9]byte{}

	for i := 0; i < Count; i++ {
		randNum := rand.Intn(9)
		randNum += 1
		randx := rand.Intn(9)
		randy := rand.Intn(9)

		sudoku[randx][randy] = byte(randNum)
	}

	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku); j++ {
			if sudoku[i][j] == 0 {
				sudoku[i][j] = '.'
			}
		}
	}

	return sudoku
}

func Check(board [9][9]byte) bool {

	i := 0
	for ; i < len(board); i++ {
		// row checker
		row := []byte{}
		j := 0

		for ; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				if !slices.Contains(row, board[i][j]) {
					row = append(row, board[i][j])
				} else {
					return false
				}
			}
		}
		row = nil

		// column checker
		col := []byte{}

		for j = 0; j < len(board); j++ {
			if board[j][i] != '.' {
				if !slices.Contains(col, board[j][i]) && board[j][i] != '.' {
					col = append(col, board[j][i])
				} else {
					return false
				}
			}
		}
		col = nil
	}

	// home checker

	homes := [3][]byte{}

	for i := 0; i < 9; i += 3 {
		for j := i; j < i+3; j++ {
			for k := 0; k < len(board[j]); k++ {
				if board[j][k] != '.' {
					if k < 3 {
						if !slices.Contains(homes[0], board[j][k]) {
							homes[0] = append(homes[0], board[j][k])
						} else {
							return false
						}
					} else if k < 6 {
						if !slices.Contains(homes[1], board[j][k]) {
							homes[1] = append(homes[1], board[j][k])
						} else {
							return false
						}
					} else {
						if !slices.Contains(homes[2], board[j][k]) {
							homes[2] = append(homes[2], board[j][k])
						} else {
							return false
						}
					}
				}
			}
			// this is for read 3 homes 3 times
			if j == 2 || j == 5 {
				homes = [3][]byte{}
			}
		}
	}

	return true
}

func CountForLevels(Level int) int {

	if Level == 1 {
		return 45
	} else if Level == 2 {
		return 40
	} else if Level == 3 {
		return 35
	} else {
		return 30
	}
}
