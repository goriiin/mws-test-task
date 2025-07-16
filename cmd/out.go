package cmd

import (
	"github.com/spf13/cobra"
	"mws/domain"
)

func out(cmd *cobra.Command, p domain.Profile) {
	cmd.Printf("name: %s\n\tuser: %s\n\tproject: %s\n", p.Name, p.User, p.Project)
}
