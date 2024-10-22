package main

import (
	creator "Sudoku/Creator/Creator"
	"fmt"
)

func main() {
	fmt.Println("Starting ...")
	fmt.Println("the first you must choose your Level !")

	fmt.Println("1- Very Easy")
	fmt.Println("2- Easy")
	fmt.Println("3- Normal")
	fmt.Println("4- Hard")

	var Level int = 2

	_, err := fmt.Scanf("%d", &Level)

	if err != nil {
		panic("Your level should be enter between 1 to 4")
	}

	if Level < 5 && Level > 0 {
		fmt.Println("Please Waiting ...")
		creator.Initialize(Level)
	} else {
		fmt.Println("you should enter number between 1 to 4")
	}

	fmt.Println("End thanks for waiting ! \a")
}
