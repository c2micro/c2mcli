package base

import (
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func Commands(app *console.Console) console.Commands {
	return func() *cobra.Command {
		rootCmd := &cobra.Command{
			DisableFlagsInUseLine: true,
		}

		// exit
		rootCmd.AddCommand(exitCommand(app))
		// beacon
		rootCmd.AddCommand(beaconsCommand(app))
		// use
		rootCmd.AddCommand(useCommand(app))
		// script
		rootCmd.AddCommand(scriptsCommand(app))

		rootCmd.InitDefaultHelpCmd()
		rootCmd.CompletionOptions.DisableDefaultCmd = true
		return rootCmd
	}
}
