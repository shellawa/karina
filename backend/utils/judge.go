package utils

import "strings"

func CompareOutputs(expected, actual string) bool {
	expLines := strings.Split(strings.TrimSpace(expected), "\n")
	actLines := strings.Split(strings.TrimSpace(actual), "\n")

	if len(expLines) != len(actLines) {
		return false
	}

	for i := range expLines {
		e := strings.TrimRight(expLines[i], " \t\r")
		a := strings.TrimRight(actLines[i], " \t\r")
		if e != a {
			return false
		}
	}
	return true
}
