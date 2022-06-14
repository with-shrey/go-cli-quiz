package domain

import (
	"errors"
	"strings"
)

type Problem struct {
	ID       int
	Question string
	Answer   string
}

type ProblemRepository interface {
	AddProblem(problem *Problem) error
	FindProblemByID(id int) (*Problem, error)
	Count() int
}

var (
	ErrInvalidQuestion = errors.New("question should not be empty")
)

func NewProblem(Question string, Answer string) (*Problem, error) {
	if len(Question) == 0 {
		return nil, ErrInvalidQuestion
	}
	return &Problem{
		ID:       0,
		Question: Question,
		Answer:   Answer,
	}, nil
}

func (problem Problem) Evaluate(answer string) bool {
	return strings.EqualFold(answer, problem.Answer)
}

func (problem Problem) GetFormattedQuestion() string {
	return "Question: " + problem.Question
}
