package repo

import (
	"gopkg.in/yaml.v3"
	"mws/domain"
	"os"
)

func (p *ProfileYAMLRepo) Create(profile domain.Profile) error {
	if !isNameValid(profile.Name) {
		return invalidName
	}

	yml, err := yaml.Marshal(domainToDTO(profile))
	if err != nil {
		return err
	}

	filename := profile.Name + p.fileType
	fullPath := p.baseDir + "/" + filename

	return os.WriteFile(fullPath, yml, os.ModePerm)
}
