package domain

type AddProblemService struct {
	ProblemRepository ProblemRepository
}

func (service AddProblemService) AddProblem(question string, answer string) (*Problem, error) {
	problem, err := NewProblem(question, answer)
	if err != nil {
		return nil, err
	}
	service.ProblemRepository.AddProblem(problem)
	return problem, nil
}
