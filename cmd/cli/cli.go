package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/with-shrey/go-quiz/adapter"
	"github.com/with-shrey/go-quiz/domain"
)

var (
	problemRepository domain.ProblemRepository
	scoreRepository domain.ScoreCardRepository
	addProblemService domain.AddProblemService
	csvAdapter adapter.CsvProblemAdapter
	userProblemService domain.GetUserProblemService
	evaluateProblemService domain.EvaluateAddProblemService
)

func initializeDependencies() {
	problemRepository = &adapter.ProblemInMemoryAdapter{
		Problems: make([]domain.Problem, 0),
	}
	scoreRepository = &adapter.ScoreCardInMemoryAdapter{
		Lock: sync.Mutex{},
		Score: *domain.NewScoreCard(),
	}
	addProblemService = domain.AddProblemService{
		ProblemRepository: problemRepository,
	}
	csvAdapter = adapter.CsvProblemAdapter{
		AddProblemService: addProblemService,
	}
	userProblemService = domain.GetUserProblemService{
		ProblemRepository: problemRepository,
		ScoreCardRepository: scoreRepository,
	}
	evaluateProblemService = domain.EvaluateAddProblemService{
		ProblemRepository: problemRepository,
		ScoreCardRepository: scoreRepository,
		UserAddProblemService: userProblemService,
	}
}

func showQuestions(channel chan<- bool) {
	for {
		problem, err := userProblemService.GetProblemToShow()
		if problem == nil {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(problem.GetFormattedQuestion())
		var answer string
		fmt.Scanln(&answer)
		evaluateProblemService.EvaluateProblem(answer)
	}
	channel <- true
}

const (
	DefaultCsvFile = "problems.csv"
	DefaultTimeout = "10"
)


func startCLI(defaultCsvFile string, defaultTimeout string) {
	// parse cli args
	csvPathArg := flag.String("csv", defaultCsvFile, "csv file path for importing problems")
	timeoutArg := flag.String("timeout", defaultTimeout, "Timeout seconds for answering questions")
	flag.Parse()
	csvPath := string(*csvPathArg)
	timeout, errorConv := strconv.Atoi(*timeoutArg)
	if errorConv != nil {
		panic(errorConv);
	}
	// initialize dependencies
	initializeDependencies()
	
	// csv handling
	

	if !filepath.IsAbs(csvPath) {
		cwd, err := os.Getwd();
		if err != nil {
			panic("Cannot get current directory")
		}
		csvPath = filepath.Join(cwd, csvPath) 
	}
	
	err := csvAdapter.ImportProblems(csvPath)
	if err != nil {
		panic(err)
	}
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	channel := make(chan bool)
	defer close(channel)
	go showQuestions(channel)
	select {
		case <-channel:
		case <-timer.C:
			fmt.Println("Time Out")
	}
	
	fmt.Println(scoreRepository.GetScore().GetFormattedResult())
	fmt.Printf("Total Question: %d\n", problemRepository.Count())
}

func main() {
	startCLI(DefaultCsvFile, DefaultTimeout)
}