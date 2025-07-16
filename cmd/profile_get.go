/*
Copyright © 2025 Koshenkov Dmitry
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var getName string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Выводит информацию о конкретном профиле.",
	Long: `Находит профиль по имени, указанному в флаге --name,
и выводит содержимое соответствующего YAML-файла в консоль.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		get, err := repo.Get(getName)
		if err != nil {
			return err
		}

		out(cmd, get)

		return nil
	},
}

func init() {
	profileCmd.AddCommand(getCmd)

	getCmd.Flags().StringVar(&getName, nameFlag, defaultName, "название профиля")
	err := getCmd.MarkFlagRequired(nameFlag)
	if err != nil {
		log.Fatalf("[ getCmd.init ] ERROR: %s ", err.Error())
	}
}
