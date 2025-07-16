package repo

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"mws/domain"
	"os"
)

func (p *ProfileYAMLRepo) Get(name string) (domain.Profile, error) {
	if !isNameValid(name) {
		return domain.Profile{}, invalidName
	}

	filename := name + p.fileType
	fullPath := p.baseDir + "/" + filename

	yml, err := os.ReadFile(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return domain.Profile{}, fmt.Errorf("профиль %s не найден", name)
		}

		return domain.Profile{}, err
	}

	var dto dtoProfile
	if err = yaml.Unmarshal(yml, &dto); err != nil {
		return domain.Profile{}, err
	}

	pr := dtoToDomain(name, dto)

	return pr, nil
}
