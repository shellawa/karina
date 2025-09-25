package models

import (
	"os"
	"path/filepath"

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

func (s *Service) AddSubmission(filePath string, participantId string, problemId string) {
	_, fileName := filepath.Split(filePath)
	submissionContent, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
		return
	}

	targetDir := filepath.Join("data", "participants", participantId, problemId)

	os.WriteFile(filepath.Join(targetDir, fileName), submissionContent, 0644)
}
