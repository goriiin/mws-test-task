package repo

import (
	"mws/domain"
	"os"
	"strings"
)

func (p *ProfileYAMLRepo) List() ([]domain.Profile, error) {
	files, err := os.ReadDir(p.baseDir)
	if err != nil {
		return nil, err
	}

	profiles := make([]domain.Profile, 0, 10)
	var profile domain.Profile

	for _, f := range files {
		filename := f.Name()
		profile, err = p.Get(filename)

		if !f.IsDir() && strings.HasSuffix(filename, p.fileType) {
			filename = strings.TrimSuffix(filename, p.fileType)

			profile, err = p.Get(filename)
			if err != nil {
				continue
			}

			profiles = append(profiles, profile)
		}
	}

	return profiles, nil
}
