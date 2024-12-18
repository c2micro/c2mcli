package beacon

import (
	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func Commands(app *console.Console) console.Commands {
	return func() *cobra.Command {
		rootCmd := &cobra.Command{
			DisableFlagsInUseLine: true,
		}

		rootCmd.AddGroup(
			&cobra.Group{ID: constants.AliasGroupId, Title: constants.AliasGroupId},
			&cobra.Group{ID: constants.CoreGroupId, Title: constants.CoreGroupId},
		)

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
		rootCmd.SetHelpCommandGroupID(constants.CoreGroupId)
		rootCmd.CompletionOptions.DisableDefaultCmd = true
		return rootCmd
	}
}
