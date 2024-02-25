package entity

type Question struct {
	ID              uint
	Question        string
	PossibleAnswers []PossibleAnswers
	CorrectAnswer   PossibleAnswers
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswers uint8

func (p PossibleAnswers) IsValid() bool {
	if p >= PossibleAnswersA && p <= PossibleAnswersD {
		return true
	}
	return false
}

const (
	PossibleAnswersA PossibleAnswers = iota + 1
	PossibleAnswersB
	PossibleAnswersC
	PossibleAnswersD
)

type QuestionDifficulty uint

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (d QuestionDifficulty) IsValid() bool {
	if d >= 1 && d <= 3 {
		return true
	}
	return false
}

func (d QuestionDifficulty) String() string {
	switch d {
	case 1:
		return "Easy"
	case 2:
		return "Medium"
	case 3:
		return "Hard"
	default:
		return ""
	}
}