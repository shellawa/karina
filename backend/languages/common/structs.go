package common

type TestSolveResult struct {
	Verdict string `json:"verdict"`
	Time    int64  `json:"time"`
	Memory  int    `json:"memory"`
}

type TestRunData struct {
	ProblemId  string
	WorkingDir string
	MaxTime    int
	MaxMemory  int
	TestInput  string
	TestOutput string
}

type Runner interface {
	Name() string
	Run(d TestRunData) TestSolveResult
}
