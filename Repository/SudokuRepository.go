package repository

import (
	models "Sudoku/Creator/Models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SudokuCollection struct {
	Sudoku *mongo.Collection
}

func (sc *SudokuCollection) Sudokus() []models.SudokuModel {
	sudokus := []models.SudokuModel{}

	res, err := sc.Sudoku.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	err = res.All(context.Background(), &sudokus)

	if err != nil {
		panic(err)
	}

	return sudokus
}

func (sc *SudokuCollection) GetSudokusByNumber(number int) models.SudokuModel {
	var sudoku models.SudokuModel

	err := sc.Sudoku.FindOne(context.Background(), bson.D{{Key: "number", Value: number}}).Decode(&sudoku)

	if err != nil {
		panic(err)
	}

	return sudoku
}

func (sc *SudokuCollection) GetLastNumber() int64 {
	var count int64
	opts := options.Count().SetHint("_id_")
	count, err := sc.Sudoku.CountDocuments(context.Background(), bson.D{}, opts)

	if err != nil {
		panic(err)
	}

	return count
}

func (sc *SudokuCollection) InsertSudoku(sm *models.SudokuModel) error {
	var duplicate models.SudokuModel
	filter := bson.A{
		"$and",
		bson.D{{Key: "digits", Value: sm.Digits}},
		bson.D{{Key: "location", Value: sm.Location}},
	}
	sc.Sudoku.FindOne(context.Background(), filter).Decode(&duplicate)

	if duplicate.ID != "" {
		return nil
	}

	_, err := sc.Sudoku.InsertOne(context.Background(), &sm)

	if err != nil {
		return err
	}
	return nil
}

func (sc *SudokuCollection) DeleteSudokuById(ID string) (bool, error) {
	res, err := sc.Sudoku.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: ID}})

	if err != nil {
		return false, err
	}

	if res.DeletedCount > 0 {
		return true, nil
	}
	return false, nil
}

func (sc *SudokuCollection) DeleteSudokuWithNumber(Number int) (bool, error) {

	res, err := sc.Sudoku.DeleteOne(context.Background(), bson.D{{Key: "number", Value: Number}})

	if err != nil {
		return false, err
	}

	if res.DeletedCount < 1 {
		return false, nil
	}

	return true, nil
}
