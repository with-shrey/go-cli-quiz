package domain

type GetUserProblemService struct {
	ProblemRepository   ProblemRepository
	ScoreCardRepository ScoreCardRepository
}

func (service GetUserProblemService) GetProblemToShow() (*Problem, error) {
	problemId := service.ScoreCardRepository.GetProblemID()
	problem, err := service.ProblemRepository.FindProblemByID(problemId)
	if err != nil {
		return nil, err
	}
	return problem, nil
}
