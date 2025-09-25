package models

import (
	"karina/backend/utils"
	"os"
	"path/filepath"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Problem struct {
	Id          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Time        int    `json:"time"`
	Memory      int    `json:"memory"`
}

func (s *Service) WriteProblem(p Problem) {
	if len(p.Id) == 0 {
		return
	}
	db.Write(filepath.Join("problems", p.Id), "specs", p)
	runtime.EventsEmit(s.ctx, "problem:change", p.Id)
}

func (s *Service) GetProblems() []Problem {
	problems := []Problem{}
	entries, _ := os.ReadDir(filepath.Join("data", "problems"))

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		specs := Problem{}
		db.Read(filepath.Join("problems", entry.Name()), "specs", &specs)
		problems = append(problems, specs)
	}
	return problems
}

func (s *Service) AddSubmission(sourcePath string, participantId string, problemId string) {
	_, fileName := filepath.Split(sourcePath)
	submissionsDir := filepath.Join("data", "participants", participantId, problemId, "submissions")
	currentSubmissionDir := filepath.Join(submissionsDir, strconv.Itoa(utils.GetLargestNumberedFolderName(submissionsDir)+1))

	os.MkdirAll(currentSubmissionDir, 0755)
	utils.CopyFile(
		sourcePath,
		filepath.Join(currentSubmissionDir, fileName),
		0644,
	)
}
