/*
Copyright © 2025 Koshenkov Dmitry
*/
package cmd

import (
	"log"
	"mws/domain"

	"github.com/spf13/cobra"
)

var createName, createUser, createProject string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создает новый профиль конфигурации.",
	Long: `Создает YAML-файл с указанным именем (--name).
Внутри файла сохраняются поля user и project на основе соответствующих флагов.

Пример:
./mws profile create --name=test --user=example --project=new-project
Создаст файл 'test.yaml' с содержимым:
user: example
project: new-project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p := domain.Profile{
			User:    createUser,
			Project: createProject,
			Name:    createName,
		}

		err := repo.Create(p)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	profileCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(
		&createName,
		nameFlag,
		defaultName,
		"название профиля (имя файла без расширения)",
	)
	err := createCmd.MarkFlagRequired(nameFlag)
	if err != nil {
		log.Printf(" [ createCmd.init ] ERROR: %s \n", err.Error())

		return
	}

	createCmd.Flags().StringVar(
		&createUser,
		userFlag,
		defaultUser,
		"имя пользователя для профиля",
	)
	err = createCmd.MarkFlagRequired(userFlag)
	if err != nil {
		log.Printf(" [ createCmd.init ] ERROR: %s \n", err.Error())

		return
	}

	createCmd.Flags().StringVar(
		&createProject,
		projectFlag,
		defaultProject,
		"название проекта для профиля",
	)
	err = createCmd.MarkFlagRequired(projectFlag)
	if err != nil {
		log.Printf(" [ createCmd.init ] ERROR: %s \n", err.Error())

		return
	}
}
