package repo

import "mws/domain"

type dtoProfile struct {
	User    string `yaml:"user"`
	Project string `yaml:"project"`
}

func domainToDTO(profile domain.Profile) dtoProfile {
	return dtoProfile{
		User:    profile.User,
		Project: profile.Project,
	}
}

func dtoToDomain(name string, profile dtoProfile) domain.Profile {
	return domain.Profile{
		User:    profile.User,
		Project: profile.Project,
		Name:    name,
	}
}
