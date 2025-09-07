package models

import (
	"context"
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

type ProblemService struct {
	ctx context.Context
}

func (s *ProblemService) Initialize(ctx context.Context) {
	s.ctx = ctx
}

func (s *ProblemService) WriteProblem(p Problem) {
	if len(p.Id) == 0 {
		return
	}
	db.Write("problems/"+p.Id, "specs", p)
	runtime.EventsEmit(s.ctx, "problem:change", p.Id)
}

func (s *ProblemService) GetProblems() []Problem {
	problems := []Problem{}
	entries, _ := os.ReadDir("data/problems")

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		specs := Problem{}
		db.Read("problems/"+entry.Name(), "specs", &specs)
		problems = append(problems, specs)
	}
	return problems
}

func (s *ProblemService) AddSubmission(filePath string, participantId string, problemId string) {
	_, fileName := filepath.Split(filePath)
	submissionContent, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
		return
	}

	targetDir := filepath.Join("data", "participants", participantId, problemId)
	print(filepath.Join(targetDir, fileName))
	os.WriteFile(filepath.Join(targetDir, fileName), submissionContent, 0644)
}
