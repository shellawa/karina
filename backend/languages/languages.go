package languages

import (
	"context"
	"karina/backend/languages/common"
	"karina/backend/languages/python"
	"karina/backend/models"
	"karina/backend/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Service struct {
	ctx     context.Context
	runners map[string]common.Runner
}

func (s *Service) Initialize(ctx context.Context) {
	s.ctx = ctx
	s.runners = map[string]common.Runner{
		"python": &python.Runner{},
	}
}

type SolveProgressUpdate struct {
	ParticipantId string `json:"participantId"`
	Language      string `json:"language"`
	TestNo        int    `json:"testNo"`
}

type SolveContext struct {
	ParticipantId string
	SubmissionId  string
	SolveId       string
	SubmissionDir string
	SolveDir      string
	Tests         []TestContext
}

type TestContext struct {
	Id     string
	Dir    string
	Input  string
	Output string
}

func (s *Service) RunAllParticipants(problemId string, maxTime int, maxMemory int, IOMode int) {
	modelsService := &models.Service{}
	modelsService.Initialize(s.ctx)
	participants := modelsService.GetParticipants(problemId)

	runtime.EventsEmit(
		s.ctx,
		"solve:run_status",
		"running",
	)

	var solves []SolveContext
	problemDir := filepath.Join("data", "problems", problemId)

	for _, participant := range participants {
		participantDir := filepath.Join(problemDir, "participants", participant.Id)

		submissionId := strconv.Itoa(
			utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions")),
		)
		solveId := strconv.Itoa(
			utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions", submissionId, "solves")) + 1,
		)

		submissionDir := filepath.Join(participantDir, "submissions", submissionId)
		solveDir := filepath.Join(submissionDir, "solves", solveId)
		os.MkdirAll(solveDir, 0755)

		tests := modelsService.GetTestCases(problemId)
		testContexts := []TestContext{}
		for _, testCase := range tests {
			testDir := filepath.Join(solveDir, testCase.Id)
			os.MkdirAll(testDir, 0755)
			testContexts = append(testContexts, TestContext{
				Id:     testCase.Id,
				Dir:    testDir,
				Input:  testCase.Input,
				Output: testCase.Output,
			})
		}

		solves = append(solves, SolveContext{
			ParticipantId: participant.Id,
			SubmissionId:  submissionId,
			SolveId:       solveId,
			SubmissionDir: submissionDir,
			SolveDir:      solveDir,
			Tests:         testContexts,
		})
	}

	runtime.EventsEmit(
		s.ctx,
		"solve:test_run_finish",
	)

	// wait for the frontend to re-sort
	time.Sleep(time.Second)

	for _, solve := range solves {
		for i, test := range solve.Tests {

			if IOMode == 1 {
				// fileio: copy .inp file into working dir
				utils.CopyFile(
					filepath.Join(problemDir, "tests", test.Id, problemId+".inp"),
					filepath.Join(solve.SubmissionDir, problemId+".inp"),
					0644,
				)
			}

			runtime.EventsEmit(s.ctx, "solve:test_run_start", SolveProgressUpdate{
				ParticipantId: solve.ParticipantId,
				Language:      "Python",
				TestNo:        i + 1,
			})

			// build test run data
			runData := common.TestRunData{
				ProblemId:  problemId,
				WorkingDir: solve.SubmissionDir,
				MaxTime:    maxTime,
				MaxMemory:  maxMemory,
				TestInput:  test.Input,
				TestOutput: test.Output,
				IOMode:     IOMode,
			}

			testSolveResult := s.runners["python"].Run(runData)

			// compare the outputs
			if !utils.CompareOutputs(test.Output, testSolveResult.TestRunOutput) {
				testSolveResult.Verdict = "WA"
			}

			if IOMode == 1 {
				// fileio: copy out file to test run folder
				utils.CopyFile(
					filepath.Join(solve.SubmissionDir, problemId+".out"),
					filepath.Join(test.Dir, problemId+".out"),
					0644,
				)

				// cleanup input/output for next test
				os.Remove(filepath.Join(solve.SubmissionDir, problemId+".inp"))
				os.Remove(filepath.Join(solve.SubmissionDir, problemId+".out"))
			} else {
				// stdio: copy captured output to test run folder
				outFile := filepath.Join(test.Dir, problemId+".out")
				os.WriteFile(outFile, []byte(testSolveResult.TestRunOutput), 0644)
			}

			models.DB.Write(test.Dir, "result", testSolveResult)

			runtime.EventsEmit(
				s.ctx,
				"solve:test_run_finish",
			)
		}

		time.Sleep(300 * time.Millisecond)
	}

	runtime.EventsEmit(
		s.ctx,
		"solve:run_status",
		"finished",
	)
}

// only support python with stdio for now
func (s *Service) GenerateTests(problemId string, testNum int) {
	generatorDir := filepath.Join("data", "problems", problemId, "generator")

	cwd, _ := os.Getwd()
	pythonRuntime := filepath.Join(cwd, "assets", "runtimes", "python", "pythonw.exe")

	cmd := exec.Command(pythonRuntime, "generate.py")
	cmd.Dir = generatorDir

	for i := 0; i < testNum; i++ {
		stdoutStderr, _ := cmd.CombinedOutput()
		os.WriteFile(filepath.Join(generatorDir, problemId+".inp"), stdoutStderr, 0644)
	}
}
