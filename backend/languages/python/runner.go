package python

import (
	"fmt"
	"karina/backend/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type Runner struct{}

func (p *Runner) Name() string {
	return "python"
}

func (p *Runner) Run(sourcePath string) string {
	participantId := "john_smith"
	problemId := "two_sum"
	submissionId := strconv.Itoa(1)
	solveId := strconv.Itoa(1)
	testId := "test01"

	cwd, _ := os.Getwd()
	problemDir := filepath.Join(cwd, "data", "problems", problemId)
	participantDir := filepath.Join(cwd, "data", "participants", participantId)

	os.MkdirAll(filepath.Join(participantDir, problemId, "submissions", submissionId, "solves", solveId, testId), 0755)

	// copy input file to solve
	utils.CopyFile(
		filepath.Join(problemDir, "tests", testId, problemId+".inp"),
		filepath.Join(participantDir, problemId, "submissions", submissionId, problemId+".inp"),
		0644,
	)

	// actual python related part
	pythonRuntime := filepath.Join(cwd, "assets", "runtimes", "python", "python.exe")
	println(pythonRuntime)
	cmd := exec.Command(pythonRuntime, problemId+".py")
	cmd.Dir = filepath.Join(participantDir, problemId, "submissions", submissionId)

	startTime := time.Now()
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		print(err)
	}
	takenTime := time.Since(startTime).String()
	fmt.Printf("%s\n", stdoutStderr)
	// that's it

	// copy out file to the solve's test folder
	utils.CopyFile(
		filepath.Join(participantDir, problemId, "submissions", submissionId, problemId+".out"),
		filepath.Join(participantDir, problemId, "submissions", submissionId, "solves", solveId, testId, problemId+".out"),
		0644,
	)

	// remove inp and out file of the test to prepare for the next one
	os.Remove(filepath.Join(participantDir, problemId, "submissions", submissionId, problemId+".inp"))
	os.Remove(filepath.Join(participantDir, problemId, "submissions", submissionId, problemId+".out"))

	return takenTime
}
