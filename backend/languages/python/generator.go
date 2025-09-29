package python

import (
	"context"
	"karina/backend/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// generator only support python with stdio for now, and the solution file must use fileio as io mode (crazy I know)
func (s *Python) Generate(ctx context.Context, problemId string, testNum int) {
	problemDir := filepath.Join("data", "problems", problemId)
	generatorDir := filepath.Join(problemDir, "generator")
	testsDir := filepath.Join(problemDir, "tests")

	cwd, _ := os.Getwd()
	pythonRuntime := filepath.Join(cwd, "assets", "runtimes", "python", "pythonw.exe")

	for i := 0; i < testNum; i++ {
		testName := strconv.Itoa(utils.GetLargestNumberedFolderName(testsDir) + 1)
		testDir := filepath.Join(testsDir, testName)
		os.MkdirAll(testDir, 0644)

		genInpCmd := exec.Command(pythonRuntime, "generate.py")
		genInpCmd.Dir = generatorDir

		genOutCmd := exec.Command(pythonRuntime, "solution.py")
		genOutCmd.Dir = generatorDir

		stdoutStderr, _ := genInpCmd.CombinedOutput()
		os.WriteFile(filepath.Join(generatorDir, problemId+".inp"), stdoutStderr, 0644)

		genOutCmd.Run()

		utils.CopyFile(filepath.Join(generatorDir, problemId+".inp"), filepath.Join(testDir, problemId+".inp"), 0644)
		utils.CopyFile(filepath.Join(generatorDir, problemId+".out"), filepath.Join(testDir, problemId+".out"), 0644)

		os.Remove(filepath.Join(generatorDir, problemId+".inp"))
		os.Remove(filepath.Join(generatorDir, problemId+".out"))

		runtime.EventsEmit(ctx, "generate:test_generated")
	}
}
