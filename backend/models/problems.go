package models

import (
	"context"
	"os"

	"github.com/sdomino/scribble"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Problem struct {
	Id          string   `json:"id"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Time        int      `json:"time"`
	Memory      int      `json:"memory"`
	Tests       []string `json:"tests"`
}

type ProblemService struct {
	ctx context.Context
}

func (s *ProblemService) Initialize(ctx context.Context) {
	s.ctx = ctx
}

var db, _ = scribble.New("data/problems", nil)

func (s *ProblemService) WriteProblem(p Problem) {
	db.Write(p.Id, "specs", p)
	runtime.EventsEmit(s.ctx, "problem:add", p.Id)
}

func (s *ProblemService) GetProblems() []Problem {
	problems := []Problem{}
	entries, _ := os.ReadDir("data/problems")

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		specs := Problem{}
		db.Read(entry.Name(), "specs", &specs)
		problems = append(problems, specs)
	}
	return problems
}
