package models

import (
	"karina/backend/utils"
	"os"
	"path/filepath"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Participant struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Organization      string `json:"organization"`
	LastSubmittedCode string `json:"code,omitempty"`
}

func (s *Service) AddOneParticipant(p Participant, problemId string) {
	if len(p.Id) == 0 {
		return
	}

	DB.Write(filepath.Join("data", "problems", problemId, "participants", p.Id), "profile", p)
	os.MkdirAll(filepath.Join("data", "problems", problemId, "participants", p.Id, "submissions"), 0755)
	runtime.EventsEmit(s.ctx, "participant:change")
}

func (s *Service) GetParticipants(problemId string) []Participant {
	partipants := []Participant{}
	entries, _ := os.ReadDir(filepath.Join("data", "problems", problemId, "participants"))

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		profile := Participant{}
		DB.Read(filepath.Join("data", "problems", problemId, "participants", entry.Name()), "profile", &profile)
		profile.LastSubmittedCode = s.GetLastSubmittedCode(problemId, profile.Id)
		partipants = append(partipants, profile)
	}

	return partipants
}

func (s *Service) GetLastSubmittedCode(problemId string, participantId string) string {
	submissionsDir := filepath.Join("data", "problems", problemId, "participants", participantId, "submissions")
	lastSubmissionId := utils.GetLargestNumberedFolderName(submissionsDir)

	raw, _ := os.ReadFile(filepath.Join(submissionsDir, strconv.Itoa(lastSubmissionId), problemId+".py"))
	return string(raw)
}
