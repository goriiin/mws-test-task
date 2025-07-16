package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

type ProfileYAMLRepo struct {
	fileType string
	baseDir  string
}

func NewProfileYAMLRepo(baseDir string) (*ProfileYAMLRepo, error) {
	if baseDir == "" {
		return nil, dirNotSpecified
	}

	cleanDir, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить абсолютный путь для '%s': %w", baseDir, err)
	}

	if err = os.MkdirAll(cleanDir, 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать или получить доступ к директории '%s': %w", cleanDir, err)
	}

	return &ProfileYAMLRepo{
		fileType: ".yaml",
		baseDir:  baseDir,
	}, nil
}
