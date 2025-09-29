package models

import (
	"karina/backend/utils"
	"os"
	"path/filepath"
)

func (s *Service) WriteGeneratorScript(problemId string, script string) {
	generatorDir := filepath.Join("data", "problems", problemId, "generator")
	os.MkdirAll(generatorDir, 0755)

	scriptPath := filepath.Join(generatorDir, "generate.py")
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

func (s *Service) WriteSolution(problemId string, sourcePath string) {
	generatorDir := filepath.Join("data", "problems", problemId, "generator")
	os.MkdirAll(generatorDir, 0755)

	targetPath := filepath.Join(generatorDir, "solution.py")
	utils.CopyFile(sourcePath, targetPath, 0644)
}
