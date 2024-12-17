package base

import (
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func scriptCommand(*console.Console) *cobra.Command {
	scriptCmd := &cobra.Command{
		Use:                   "script",
		Short:                 "manage scripts",
		DisableFlagsInUseLine: true,
	}

	scriptCmd.AddCommand()

	return scriptCmd
}
