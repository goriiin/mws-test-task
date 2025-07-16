/*
Copyright © 2025 Koshenkov Dmitry
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Выводит список всех существующих профилей.",
	Long:  `Сканирует текущую директорию на наличие файлов с расширением .yaml и выводит их имена (без расширения) как список доступных профилей.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := repo.List()
		if err != nil {
			return err
		}

		for _, v := range all {
			out(cmd, v)
		}

		return nil
	},
}

func init() {
	profileCmd.AddCommand(listCmd)
}
