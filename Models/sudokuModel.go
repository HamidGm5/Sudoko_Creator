package models

type SudokuModel struct {
	ID       string `json:"id,omitempty" bson:"_id" `
	Number   int32  `json:"number,omitempty" bson:"number"`
	Digits   string `json:"digits,omitempty" bson:"digits"`     // the every characters have two character for location
	Location string `json:"location,omitempty" bson:"location"` // the even digits is the x axis and odd digits is y axis
}
