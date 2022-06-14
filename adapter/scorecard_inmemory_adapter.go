package adapter

import (
	"sync"

	"github.com/with-shrey/go-quiz/domain"
)

type ScoreCardInMemoryAdapter struct {
	Lock  sync.Mutex
	Score domain.ScoreCard
}

func (scoreRepository *ScoreCardInMemoryAdapter) AnsweredCorrectly() {
	scoreRepository.Lock.Lock()
	scoreRepository.Score.CorrectAnswers = scoreRepository.Score.CorrectAnswers + 1
	scoreRepository.Score.CurrentProblemIndex = scoreRepository.Score.CurrentProblemIndex + 1
	scoreRepository.Lock.Unlock()
}

func (scoreRepository *ScoreCardInMemoryAdapter) IncorrectAnswer() {
	scoreRepository.Lock.Lock()
	scoreRepository.Score.WrongAnswers = scoreRepository.Score.WrongAnswers + 1
	scoreRepository.Score.CurrentProblemIndex = scoreRepository.Score.CurrentProblemIndex + 1
	scoreRepository.Lock.Unlock()
}

func (scoreRepository *ScoreCardInMemoryAdapter) GetProblemID() int {
	return scoreRepository.Score.CurrentProblemIndex
}

func (scoreRepository *ScoreCardInMemoryAdapter) GetScore() domain.ScoreCard {
	return scoreRepository.Score
}
