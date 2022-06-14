package domain

type EvaluateAddProblemService struct {
	ProblemRepository     ProblemRepository
	ScoreCardRepository   ScoreCardRepository
	UserAddProblemService GetUserProblemService
}

func (service EvaluateAddProblemService) EvaluateProblem(answer string) error {
	problem, err := service.UserAddProblemService.GetProblemToShow()
	if err != nil {
		return err
	}
	if problem.Evaluate(answer) {
		service.ScoreCardRepository.AnsweredCorrectly()
	} else {
		service.ScoreCardRepository.IncorrectAnswer()
	}
	return nil
}
