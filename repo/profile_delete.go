package repo

import (
	"fmt"
	"os"
)

func (p *ProfileYAMLRepo) Delete(name string) error {
	if !isNameValid(name) {
		return invalidName
	}

	fileName := name + p.fileType
	fullPath := p.baseDir + "/" + fileName
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("профиль с именем %s не существует, папка %s", fileName, p.baseDir)
	}

	return os.Remove(fullPath)
}
