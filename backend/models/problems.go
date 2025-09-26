package models

import (
	"karina/backend/languages/common"
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

type TestCase struct {
	Id     string
	Input  string
	Output string
}

type ParticipantSolveResult struct {
	Id      string                   `json:"id"`
	Results []common.TestSolveResult `json:"results"`
}

func (s *Service) WriteProblem(p Problem) {
	if len(p.Id) == 0 {
		return
	}
	DB.Write(filepath.Join("data", "problems", p.Id), "specs", p)
	os.MkdirAll(filepath.Join("data", "problems", p.Id, "participants"), 0755)
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
		DB.Read(filepath.Join("data", "problems", entry.Name()), "specs", &specs)
		problems = append(problems, specs)
	}
	return problems
}

func (s *Service) AddSubmission(sourcePath string, participantId string, problemId string) {
	_, fileName := filepath.Split(sourcePath)
	submissionsDir := filepath.Join("data", "problems", problemId, "participants", participantId, "submissions")
	os.MkdirAll(submissionsDir, 0755)

	currentSubmissionDir := filepath.Join(submissionsDir, strconv.Itoa(utils.GetLargestNumberedFolderName(submissionsDir)+1))
	os.MkdirAll(currentSubmissionDir, 0755)

	utils.CopyFile(
		sourcePath,
		filepath.Join(currentSubmissionDir, fileName),
		0644,
	)
}

func (s *Service) GetTestCases(problemId string) []TestCase {
	testCases := []TestCase{}

	testsDirPath := filepath.Join("data", "problems", problemId, "tests")
	testNames := utils.GetAllDirectories(testsDirPath)

	for _, testName := range testNames {
		testPath := filepath.Join(testsDirPath, testName)

		testInput := ""
		testOutput := ""
		entries, _ := os.ReadDir(testPath)

		for _, e := range entries {
			if e.IsDir() {
				continue
			}

			if e.Name() == problemId+".inp" {
				testInputByte, _ := os.ReadFile(filepath.Join(testPath, problemId+".inp"))
				testInput = string(testInputByte)
			}

			if e.Name() == problemId+".out" {
				testOutputByte, _ := os.ReadFile(filepath.Join(testPath, problemId+".out"))
				testOutput = string(testOutputByte)
			}
		}

		if testInput == "" || testOutput == "" {
			continue
		}

		testCases = append(testCases, TestCase{
			Id:     testName,
			Input:  testInput,
			Output: testOutput,
		})
	}

	return testCases
}

// wouldn't recalling this everytime a new test is judge very costly?
func (s *Service) GetSolveResults(problemId string) []ParticipantSolveResult {
	participantSolveResults := []ParticipantSolveResult{}

	participants := s.GetParticipants(problemId)
	for _, participant := range participants {
		// sorting out ids and dirs
		problemDir := filepath.Join("data", "problems", problemId)
		participantDir := filepath.Join(problemDir, "participants", participant.Id)

		submissionId := strconv.Itoa(utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions")))
		solveId := strconv.Itoa(utils.GetLargestNumberedFolderName(filepath.Join(participantDir, "submissions", submissionId, "solves")))

		submissionDir := filepath.Join(participantDir, "submissions", submissionId)
		solveDir := filepath.Join(submissionDir, "solves", solveId)

		tests := s.GetTestCases(problemId)
		testSolveResults := []common.TestSolveResult{}

		for _, test := range tests {
			testDir := filepath.Join(solveDir, test.Id)
			testSolveResult := common.TestSolveResult{}
			DB.Read(testDir, "result", &testSolveResult)
			testSolveResults = append(testSolveResults, testSolveResult)
		}

		participantSolveResults = append(
			participantSolveResults,
			ParticipantSolveResult{Id: participant.Id, Results: testSolveResults})
	}

	return participantSolveResults
}
