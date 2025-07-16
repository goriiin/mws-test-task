/*
Copyright © 2025 Koshenkov Dmitry
*/
package cmd

import (
	"mws/domain"
	"os"

	"github.com/spf13/cobra"
)

//go:generate mockgen -source=root.go -destination=../cmd/mock/mock_repo.go -package=mock
type ProfileRepo interface {
	Create(profile domain.Profile) error
	List() ([]domain.Profile, error)
	Get(name string) (domain.Profile, error)
	Delete(name string) error
}

var repo ProfileRepo

func NewProfileRepo(r ProfileRepo) {
	repo = r
}

const (
	nameFlag    = "name"
	projectFlag = "project"
	userFlag    = "user"

	defaultName    = ""
	defaultUser    = ""
	defaultProject = ""
)

var rootCmd = &cobra.Command{
	Use:   "mws",
	Short: "CLI для управления профилями конфигурации.",
	Long: `mws - это утилита командной строки для создания и управления
профилями конфигурации в виде YAML-файлов.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
