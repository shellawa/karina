package python

import (
	"context"
	"karina/backend/languages/common"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func (p *Python) Run(_ context.Context, d common.TestRunData) common.TestSolveResult {
	cwd, _ := os.Getwd()
	pythonRuntime := filepath.Join(cwd, "assets", "runtimes", "python", "pythonw.exe")

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(d.MaxTime))
	defer cancel()

	cmd := exec.CommandContext(ctx, pythonRuntime, d.ProblemId+".py")
	cmd.Dir = d.WorkingDir

	var testRunOutput string

	if d.IOMode == 0 {
		cmd.Stdin = strings.NewReader(d.TestInput)
	}

	startTime := time.Now()
	takenTime := int64(d.MaxTime)

	out, err := cmd.CombinedOutput()
	if d.IOMode == 0 {
		testRunOutput = string(out)
	} else {
		outFile, err := os.ReadFile(filepath.Join(d.WorkingDir, d.ProblemId+".out"))
		if err == nil {
			testRunOutput = string(outFile)
		}
	}

	verdict := "AC"
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			verdict = "TLE"
		} else {
			verdict = "RE"
		}
	} else {
		takenTime = time.Since(startTime).Milliseconds()
	}

	return common.TestSolveResult{
		Verdict:       verdict,
		Time:          takenTime,
		Memory:        727,
		TestRunOutput: testRunOutput,
	}
}
