/*
Copyright © 2025 Koshenkov Dmitry
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var deleteName string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Удаляет указанный профиль.",
	Long:  `Находит и удаляет YAML-файл профиля по имени, указанному в флаге --name.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		err := repo.Delete(deleteName)

		return err
	},
}

func init() {
	profileCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVar(
		&deleteName,
		nameFlag,
		defaultName,
		"название профиля для удаления",
	)

	err := deleteCmd.MarkFlagRequired(nameFlag)
	if err != nil {
		return
	}
}
