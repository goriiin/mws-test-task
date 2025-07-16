package repo

import "strings"

func isNameValid(name string) bool {
	if name == "" {
		return false
	}

	if strings.Contains(name, "..") {
		return false
	}

	if strings.Contains(name, "/") {
		return false
	}

	return true
}
