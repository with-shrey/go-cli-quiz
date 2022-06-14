package domain

import "fmt"

type ScoreCard struct {
	CurrentProblemIndex int
	CorrectAnswers      int
	WrongAnswers        int
}

type ScoreCardRepository interface {
	AnsweredCorrectly()
	IncorrectAnswer()
	GetProblemID() int
	GetScore() ScoreCard
}

func NewScoreCard() *ScoreCard {
	return &ScoreCard{
		CurrentProblemIndex: 0,
		CorrectAnswers:      0,
		WrongAnswers:        0,
	}
}

func (score ScoreCard) GetFormattedResult() string {
	return fmt.Sprintf("Correct Answers: %d", score.CorrectAnswers)
}
