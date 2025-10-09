package common

import "context"

type TestSolveResult struct {
	Verdict       string `json:"verdict"`
	Time          int64  `json:"time"`
	Memory        int    `json:"memory"`
	TestRunOutput string `json:"actual_output"`
}

type TestRunData struct {
	ProblemId  string
	WorkingDir string
	MaxTime    int
	MaxMemory  int
	TestInput  string
	TestOutput string
	IOMode     int
}

type Language interface {
	Name() string
	Run(ctx context.Context, d TestRunData) TestSolveResult
	Generate(ctx context.Context, problemId string, testNum int)
}
