package models

import (
	"context"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Participant struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

type ParticipantService struct {
	ctx context.Context
}

func (s *ParticipantService) Initialize(ctx context.Context) {
	s.ctx = ctx
}

func (s *ParticipantService) AddOneParticipant(p Participant, problemId string) {
	if len(p.Id) == 0 {
		return
	}

	db.Write(filepath.Join("participants", p.Id), "profile", p)
	os.MkdirAll(filepath.Join("data", "participants", p.Id, problemId, "solves"), 0755)
	runtime.EventsEmit(s.ctx, "participant:change")
}

func (s *ParticipantService) GetParticipants() []Participant {
	partipants := []Participant{}
	entries, _ := os.ReadDir(filepath.Join("data", "participants"))

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		profile := Participant{}
		db.Read(filepath.Join("participants", entry.Name()), "profile", &profile)
		partipants = append(partipants, profile)
	}

	return partipants
}
