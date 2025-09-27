package languages

import (
	"context"
	"karina/backend/languages/common"
	"karina/backend/languages/python"
	"karina/backend/models"
	"karina/backend/utils"
	"os"
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

func (s *Service) RunAllParticipants(problemId string, maxTime int, maxMemory int) {
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
	time.Sleep(time.Millisecond * 1000)

	for _, solve := range solves {
		for i, test := range solve.Tests {
			// copy inp file to submission folder (working folder)
			utils.CopyFile(
				filepath.Join(problemDir, "tests", test.Id, problemId+".inp"),
				filepath.Join(solve.SubmissionDir, problemId+".inp"),
				0644,
			)

			runtime.EventsEmit(
				s.ctx,
				"solve:test_run_start",
				SolveProgressUpdate{
					ParticipantId: solve.ParticipantId,
					Language:      "Python",
					TestNo:        i + 1,
				})

			testSolveResult := s.runners["python"].Run(common.TestRunData{
				ProblemId:  problemId,
				WorkingDir: solve.SubmissionDir,
				MaxTime:    maxTime,
				MaxMemory:  maxMemory,
				TestInput:  test.Input,
				TestOutput: test.Output,
			})

			// compare the outputs
			if testSolveResult.Verdict == "AC" {
				solveOutputBytes, _ := os.ReadFile(filepath.Join(solve.SubmissionDir, problemId+".out"))
				solveOutput := string(solveOutputBytes)
				if !utils.CompareOutputs(test.Output, solveOutput) {
					testSolveResult.Verdict = "WA"
				}
			}

			// copy out file to the solve's test folder
			utils.CopyFile(
				filepath.Join(solve.SubmissionDir, problemId+".out"),
				filepath.Join(test.Dir, problemId+".out"),
				0644,
			)

			models.DB.Write(test.Dir, "result", testSolveResult)

			// remove inp and out file of the test to prepare for the next one
			os.Remove(filepath.Join(solve.SubmissionDir, problemId+".inp"))
			os.Remove(filepath.Join(solve.SubmissionDir, problemId+".out"))

			runtime.EventsEmit(
				s.ctx,
				"solve:test_run_finish",
			)
		}

		time.Sleep(time.Millisecond * 300)
	}

	runtime.EventsEmit(
		s.ctx,
		"solve:run_status",
		"finished",
	)

}
