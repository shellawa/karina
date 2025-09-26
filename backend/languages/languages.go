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

func (s *Service) RunAllParticipants(problemId string, time int, memory int) {
	modelsService := &models.Service{}
	modelsService.Initialize(s.ctx)

	runtime.EventsEmit(
		s.ctx,
		"solve:run_status",
		"running",
	)

	participants := modelsService.GetParticipants(problemId)
	for _, participant := range participants {

		// sorting out ids and dirs
		problemDir := filepath.Join("data", "problems", problemId)
		participantDir := filepath.Join(problemDir, "participants", participant.Id)

		submissionId := strconv.Itoa(utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions")))
		solveId := strconv.Itoa(utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions", submissionId, "solves")) + 1) // this is the next solve

		submissionDir := filepath.Join(participantDir, "submissions", submissionId)
		solveDir := filepath.Join(submissionDir, "solves", solveId)
		os.MkdirAll(solveDir, 0755)

		tests := modelsService.GetTestCases(problemId)

		for i, test := range tests {
			testDir := filepath.Join(solveDir, test.Id)
			os.MkdirAll(testDir, 0755)

			// copy inp file to submission folder (working folder)
			utils.CopyFile(
				filepath.Join(problemDir, "tests", test.Id, problemId+".inp"),
				filepath.Join(submissionDir, problemId+".inp"),
				0644,
			)

			runtime.EventsEmit(
				s.ctx,
				"solve:test_run_start",
				SolveProgressUpdate{
					ParticipantId: participant.Id,
					Language:      "Python",
					TestNo:        i + 1,
				})

			testSolveResult := s.runners["python"].Run(common.TestRunData{
				ProblemId:  problemId,
				WorkingDir: submissionDir,
				MaxTime:    time,
				MaxMemory:  memory,
				TestInput:  test.Input,
				TestOutput: test.Output,
			})

			// compare the outputs
			if testSolveResult.Verdict == "AC" {
				solveOutputBytes, _ := os.ReadFile(filepath.Join(submissionDir, problemId+".out"))
				solveOutput := string(solveOutputBytes)
				if !utils.CompareOutputs(test.Output, solveOutput) {
					testSolveResult.Verdict = "WA"
				}
			}

			// copy out file to the solve's test folder
			utils.CopyFile(
				filepath.Join(submissionDir, problemId+".out"),
				filepath.Join(testDir, problemId+".out"),
				0644,
			)

			models.DB.Write(testDir, "result", testSolveResult)

			// remove inp and out file of the test to prepare for the next one
			os.Remove(filepath.Join(submissionDir, problemId+".inp"))
			os.Remove(filepath.Join(submissionDir, problemId+".out"))

			runtime.EventsEmit(
				s.ctx,
				"solve:test_run_finish",
			)
		}

	}
	runtime.EventsEmit(
		s.ctx,
		"solve:run_status",
		"finished",
	)

}
