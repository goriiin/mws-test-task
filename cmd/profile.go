/*
Copyright © 2025 Koshenkov Dmitry
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Группа команд для работы с профилями",
	Long:  `Позволяет создавать, просматривать, перечислять и удалять профили конфигурации.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(args)
		cmd.Println("Необходимо указать подкоманду: create, get, list, delete.")
		cmd.Println("Используйте './mws profile --help' для просмотра списка команд.")
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
