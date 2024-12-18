package beacon

import (
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func Commands(app *console.Console) console.Commands {
	return func() *cobra.Command {
		rootCmd := &cobra.Command{
			DisableFlagsInUseLine: true,
		}

		// last
		rootCmd.AddCommand(lastCommand(app))
		// tasks
		rootCmd.AddCommand(tasksCommand(app))
		// exit
		rootCmd.AddCommand(exitCommand(app))
		// алиасы
		for _, v := range aliasCommands(app) {
			rootCmd.AddCommand(v)
		}

		rootCmd.InitDefaultHelpCmd()
		rootCmd.CompletionOptions.DisableDefaultCmd = true
		return rootCmd
	}
}
