package models

import (
	"os"
	"path/filepath"
)

func (s *Service) WriteGeneratorScript(problemId string, script string) {
	scriptPath := filepath.Join("data", "problems", problemId, "generator", "generate.py")
	os.WriteFile(scriptPath, []byte(script), 0644)
}

func (s *Service) GetGeneratorScript(problemId string) string {
	if problemId == "" {
		return ""
	}

	scriptDir := filepath.Join("data", "problems", problemId, "generator")
	scriptPath := filepath.Join(scriptDir, "generate.py")
	os.MkdirAll(scriptDir, 0755)
	b, err := os.ReadFile(scriptPath)
	if err != nil {
		return ""
	}
	return string(b)
}
