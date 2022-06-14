package adapter

import (
	"errors"

	"github.com/with-shrey/go-quiz/domain"
)

type ProblemInMemoryAdapter struct {
	Problems []domain.Problem
}

func (problemRepository *ProblemInMemoryAdapter) AddProblem(problem *domain.Problem) error {
	problem.ID = len(problemRepository.Problems)
	problemRepository.Problems = append(problemRepository.Problems, *problem)
	return nil
}

func (problemRepository *ProblemInMemoryAdapter) FindProblemByID(id int) (*domain.Problem, error) {
	if len(problemRepository.Problems) <= id {
		return nil, errors.New("problem with this ID not found")
	}
	return &problemRepository.Problems[id], nil
}

func (problemRepository *ProblemInMemoryAdapter) Count() int {
	return len(problemRepository.Problems)
}
