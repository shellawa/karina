package common

type TestSolveVerdict struct {
	Verdict string
	Time    int64
	Memory  int
}

type TestRunData struct {
	ProblemId  string
	WorkingDir string
	MaxTime    int
	MaxMemory  int
}

type Runner interface {
	Name() string
	Run(d TestRunData) TestSolveVerdict
}
