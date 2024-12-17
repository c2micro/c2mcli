package base

import (
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func useCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "use",
		Short:                 "switch on beacon shell",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
