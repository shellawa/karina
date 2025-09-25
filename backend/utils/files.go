package utils

import (
	"os"
	"strconv"
)

func CopyFile(sourcePath string, targetPath string, perm os.FileMode) {
	fileContent, _ := os.ReadFile(sourcePath)
	os.WriteFile(targetPath, fileContent, perm)
}

func GetLargestNumberedFolderName(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
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
