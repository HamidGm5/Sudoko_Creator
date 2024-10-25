package main

import (
	controller "Sudoku/Creator/Controller"
	creator "Sudoku/Creator/Creator"
	"context"
	"fmt"
	"log"

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
	fmt.Println("Starting ...")
	fmt.Println("the first you must choose your Level !")

	fmt.Println("1- Very Easy")
	fmt.Println("2- Easy")
	fmt.Println("3- Normal")
	fmt.Println("4- Hard")

	var Level int = 0

	_, err := fmt.Scanf("%d", &Level)

	if err != nil {
		panic("Your level should be enter between 1 to 4")
	}

	if Level < 5 && Level > 0 {
		fmt.Println("Please Waiting ...")
		board := creator.Initialize(Level)
		InitializeDbContext()
		insertToDataBase(board)
	} else {
		fmt.Println("you should enter number between 1 to 4")
	}

	fmt.Println("End thanks for waiting ! \a")
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
