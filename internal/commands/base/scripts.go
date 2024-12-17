package base

import (
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func scriptsCommand(*console.Console) *cobra.Command {
	scriptsCmd := &cobra.Command{
		Use:                   "scripts",
		Short:                 "manage scripts",
		DisableFlagsInUseLine: true,
	}

	scriptsCmd.AddCommand()

	return scriptsCmd
}
