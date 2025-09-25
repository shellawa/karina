package utils

import (
	"os"
	"strconv"
)

func CopyFile(sourcePath string, targetPath string, perm os.FileMode) {
	fileContent, _ := os.ReadFile(sourcePath)
	os.WriteFile(targetPath, fileContent, perm)
}

func GetAllDirectories(dir string) []string {
	dirs := []string{}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
		}
	}

	return dirs
}

func GetLargestNumberedFolderName(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0 // in case folder didn't exist, it probably means that it's the first one to be created
	}

	maxNum := 0

	for _, e := range entries {
		if e.IsDir() {
			name := e.Name()
			n, err := strconv.Atoi(name)
			if err != nil {
				continue
			}
			if n > maxNum {
				maxNum = n
			}
		}
	}
	return maxNum
}
