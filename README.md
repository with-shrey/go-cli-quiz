# CLI Quiz App in GoLang

## Introduction
I have tried to implement Domain driven design along with a miniature hexagonal architecture

- `/domain` - contains domain entity and domain services 
  - domain entity also contain repository interface which are out bound ports 
- `/adpaters` have implementation on in-memory storage implementation of repositories 
  - csv_problem_adapter is an in bound adapter that uses domain service as port
- `/cmd` contains all executables  

## How to run 
- copy `/problems.csv` to `cmd/cli/problems.csv`
- `go mod download`
- `go build github.com/with-shrey/go-quiz/cmd/cli`
```
  Usage of cli:
  -csv path (absolute or relative)
        csv file path for importing problems (default "problems.csv")
  -timeout string
        Timeout seconds for answering questions (default "10")  
  -h help
        list of possible flag the app supports
```

## Run Tests
- Acceptance tests
  - located at : `/cmd/cli/cli_test`
  - go test github.com/with-shrey/go-quiz/cmd/cli
