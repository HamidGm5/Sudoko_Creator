package controller

import (
	models "Sudoku/Creator/Models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sudokuCollection struct {
	sudokus *mongo.Collection
}

var (
	DataBase         string = "Sudoku_Generator"
	SudokuCollection string = "Sudokus"
)

func (sc *sudokuCollection) Sudokus() []models.SudokuModel {
	sudokus := []models.SudokuModel{}

	res, err := sc.sudokus.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	err = res.All(context.Background(), &sudokus)

	if err != nil {
		panic(err)
	}

	return sudokus
}

func (sc *sudokuCollection) GetSudokusByNumber(number int) models.SudokuModel {
	var sudoku models.SudokuModel

	err := sc.sudokus.FindOne(context.Background(), bson.D{{Key: "number", Value: number}}).Decode(&sudoku)

	if err != nil {
		panic(err)
	}

	return sudoku
}

func (sc *sudokuCollection) GetLastNumber() int64 {
	var count int64
	opts := options.Count().SetHint("_id")
	count, err := sc.sudokus.CountDocuments(context.Background(), bson.D{}, opts)

	if err != nil {
		panic(err)
	}
	return count
}

func (sc *sudokuCollection) InsertSudoku(sm *models.SudokuModel) error {
	var duplicate models.SudokuModel
	filter := bson.A{
		"$and",
		bson.D{{Key: "digits", Value: sm.Digits}},
		bson.D{{Key: "location", Value: sm.Location}},
	}
	err := sc.sudokus.FindOne(context.Background(), filter).Decode(&duplicate)

	if err != nil {
		return err
	}

	if duplicate.ID != "" {
		return nil
	}

	_, err = sc.sudokus.InsertOne(context.Background(), &sm)
	if err != nil {
		return err
	}
	return nil
}

func (sc *sudokuCollection) DeleteSudokuById(ID string) (bool, error) {
	res, err := sc.sudokus.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: ID}})

	if err != nil {
		return false, err
	}

	if res.DeletedCount > 0 {
		return true, nil
	}
	return false, nil
}
