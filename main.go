package main

import (
	controller "Sudoku/Creator/Controller"
	creator "Sudoku/Creator/Creator"
	writer "Sudoku/Creator/Writer"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

var (
	DataBase         = "Sudoku_Generator"
	URI              = "mongodb://localhost:27017/" + DataBase
	SudokuCollection = "Sudokus"
)

func main() {
	fmt.Println("Hello wellcome to Sudoku game ")
	Start()
}

func Start() {
	fmt.Println("the first choose number 1 for going sudoku with number and number 2 for random sudoku")
	fmt.Println("1-Level")
	fmt.Println("2-Random")

	var choose int = 0
	_, err := fmt.Scanf("%d", &choose)

	if err != nil {
		Restart()
	}

	if choose == 1 {

		var number int = 0
		fmt.Println("Enter number of your sudoku you want : ")
		_, err = fmt.Scanf("%d", &number)

		if err != nil {
			Restart()
		}

		fmt.Println()

		if number < 1 {
			fmt.Println("your number is not valid")
			Restart()
		}

		InitializeDbContext()
		mongoCollection := mongoClient.Database(DataBase).Collection(SudokuCollection)
		defer mongoClient.Disconnect(context.Background())

		Controller := controller.SudokuCollection{Collection: mongoCollection}

		board := Controller.GetSudokuWithNumber(number)

		writer.WriteWithBoard(&board, "ConstantBoards.txt")
		fmt.Println("Your sudoku written")
	} else if choose == 2 {

		fmt.Println("Enter level for random sudoku ")
		fmt.Println("1- Very Easy")
		fmt.Println("2- Easy")
		fmt.Println("3- Normal")
		fmt.Println("4- Hard")

		fmt.Scanf("choose :", &choose)

		if choose < 5 && choose > 0 {
			fmt.Println("Please Waiting ...")
			board := creator.Initialize(choose)
			writer.WriteWithBoard(&board, "RandomBoards.txt")
			InitializeDbContext()
			insertToDataBase(board)
		} else {
			fmt.Println("you should enter number between 1 to 4")
		}

	} else {
		fmt.Println("your number is not valid")
		Restart()
	}

}

func Restart() {
	fmt.Println("if you want play again enter Y and for close game enter N")
	var choose byte
	_, err := fmt.Scanf("%c", &choose)

	if err != nil {
		os.Exit(0)
	}

	if bytes.ToLower([]byte{choose})[0] == 'y' {
		Start()
	}
	os.Exit(0)
}

func InitializeDbContext() {
	MongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))

	if err != nil {
		log.Fatal("Cann't connect to database")
	}

	err = MongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("we cann't ping database !")
	}
	mongoClient = MongoClient
	log.Print("Connect to data base was successfully")
}

func insertToDataBase(newBoard [9][9]byte) {
	mongoCollection := mongoClient.Database(DataBase).Collection(SudokuCollection)
	defer mongoClient.Disconnect(context.Background())

	Controller := controller.SudokuCollection{Collection: mongoCollection}

	err := Controller.InsertSudoku(newBoard)

	if err != nil {
		log.Fatal("somthing went wrong !")
	}
}
