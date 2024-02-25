package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	PlayerIDs   []uint
	WinnerID    uint
	StartTime   time.Time
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   uint
	Answers []PossibleAnswers
}