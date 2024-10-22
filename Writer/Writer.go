package writer

import (
	"os"
	"strconv"
)

func WriteInFile(board *[9][9]byte) {

	file, err := os.OpenFile("./Boards/boards.txt", 1, os.ModeAppend)
	charCounter := 0
	colCounter := 0

	if err != nil {
		panic(err)
	}

	file.WriteString("\n")
	file.WriteString("\n")
	file.WriteString("\t------------------------------------------------- \n")
	for i := 0; i < len(board); i++ {

		if colCounter == 3 {
			file.WriteString("\t-------------------------------------------------\n")
			colCounter = 0
		}
		var Line string
		Line += "|\t"
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				Line += strconv.Itoa(int(board[i][j]))
			} else {
				Line += "?"
			}
			Line += "\t"
			charCounter++
			if charCounter == 3 {
				Line += "|\t"
				charCounter = 0
			}
		}

		colCounter++
		file.WriteString("\t")
		file.WriteString(Line)
		file.WriteString("\n")
	}
	file.WriteString("\t------------------------------------------------- \n")
	file.WriteString("\n")
	file.WriteString("\n")
}
