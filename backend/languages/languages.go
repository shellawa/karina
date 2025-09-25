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

func (s *Service) RunAllParticipants(problemId string, time int, memory int) {
	modelsService := &models.Service{}
	modelsService.Initialize(s.ctx)

	participants := modelsService.GetParticipants(problemId)
	for _, participant := range participants {
		cwd, _ := os.Getwd()

		// sorting out ids and dirs
		problemDir := filepath.Join(cwd, "data", "problems", problemId)
		participantDir := filepath.Join(problemDir, "participants", participant.Id)

		submissionId := strconv.Itoa(utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions")))
		solveId := strconv.Itoa(utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions", submissionId, "solves")) + 1)

		submissionDir := filepath.Join(participantDir, "submissions", submissionId)
		solveDir := filepath.Join(submissionDir, "solves", solveId)
		os.MkdirAll(solveDir, 0755)

		testIds := utils.GetAllDirectories(filepath.Join(problemDir, "tests"))

		for _, testId := range testIds {
			testDir := filepath.Join(solveDir, testId)
			os.MkdirAll(testDir, 0755)

			// copy inp file to submission folder (working folder)
			utils.CopyFile(
				filepath.Join(problemDir, "tests", testId, problemId+".inp"),
				filepath.Join(submissionDir, problemId+".inp"),
				0644,
			)

			s.runners["python"].Run(common.TestRunData{
				ProblemId:  problemId,
				WorkingDir: submissionDir,
				MaxTime:    time,
				MaxMemory:  memory,
			})

			// copy out file to the solve's test folder
			utils.CopyFile(
				filepath.Join(submissionDir, problemId+".out"),
				filepath.Join(testDir, problemId+".out"),
				0644,
			)

			// remove inp and out file of the test to prepare for the next one
			os.Remove(filepath.Join(submissionDir, problemId+".inp"))
			os.Remove(filepath.Join(submissionDir, problemId+".out"))
		}

	}
}
