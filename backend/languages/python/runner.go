package python

import (
	"context"
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

func (p *Runner) Run(d common.TestRunData) common.TestSolveResult {
	cwd, _ := os.Getwd()
	pythonRuntime := filepath.Join(cwd, "assets", "runtimes", "python", "python.exe")
	verdict := "AC"

	// context to interrupt process with time limit
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(d.MaxTime))
	defer cancel()

	cmd := exec.CommandContext(ctx, pythonRuntime, d.ProblemId+".py")
	cmd.Dir = d.WorkingDir

	startTime := time.Now()
	takenTime := int64(d.MaxTime)

	_, err := cmd.CombinedOutput()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			verdict = "TLE"
		} else {
			verdict = "Other"
		}
	} else {
		takenTime = time.Since(startTime).Milliseconds()
	}

	return common.TestSolveResult{
		Verdict: verdict,
		Time:    takenTime,
		Memory:  456,
	}
}
