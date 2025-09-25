package python

import (
	"fmt"
	"karina/backend/languages/common"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Runner struct{}

func (p *Runner) Name() string {
	return "python"
}

func (p *Runner) Run(d common.TestRunData) common.TestSolveVerdict {
	cwd, _ := os.Getwd()

	pythonRuntime := filepath.Join(cwd, "assets", "runtimes", "python", "python.exe")
	cmd := exec.Command(pythonRuntime, d.ProblemId+".py")
	cmd.Dir = d.WorkingDir

	startTime := time.Now()
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		print(err)
	}
	takenTime := time.Since(startTime).Milliseconds()
	fmt.Printf("%s\n", stdoutStderr)

	return common.TestSolveVerdict{
		Verdict: "AC",
		Time:    takenTime,
		Memory:  456,
	}
}
