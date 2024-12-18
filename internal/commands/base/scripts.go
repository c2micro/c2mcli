package base

import (
	"github.com/c2micro/c2mcli/internal/scripts"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func scriptsLoadCommand(*console.Console) *cobra.Command {
	scriptsLoadCmd := &cobra.Command{
		Use:                   "load",
		Short:                 "load script by path on FS",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := scripts.RegisterExternalByPath(args[0]); err != nil {
				color.Red(err.Error())
				return
			}
			color.Green("script successfully registered")
		},
	}
	// генерация позиционного комплитера
	carapace.Gen(scriptsLoadCmd).PositionalCompletion(scriptsLoadCommandCompleter())
	return scriptsLoadCmd
}

func scriptsListCommand(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "list",
		Short:                 "list registred scripts",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			registeredScripts := scripts.GetScripts()
			if len(registeredScripts) == 0 {
				color.Yellow("no scripts registered")
				return
			}
			for _, v := range registeredScripts {
				timestamp := v.GetAddedAt().Format("01/02 15:04:05")
				c.Printf("[%s] %s\n",
					timestamp,
					v.GetPath(),
				)
			}
		},
	}
}

func scriptsRemoveCommand(*console.Console) *cobra.Command {
	scriptsRemoveCmd := &cobra.Command{
		Use:                   "remove",
		Short:                 "remove registred scripts",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := scripts.RemoveExternalByPath(args[0]); err != nil {
				color.Red(err.Error())
				return
			}
			color.Green("script %s removed", args[0])
		},
	}
	// генерация позиционного комплитера
	carapace.Gen(scriptsRemoveCmd).PositionalCompletion(externalScriptsCompleter())
	return scriptsRemoveCmd
}

func scriptsReloadCommand(*console.Console) *cobra.Command {
	scriptsReloadCmd := &cobra.Command{
		Use:                   "reload",
		Short:                 "reload script/all scripts",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// релоад всех скриптов
				if err := scripts.Rebuild(); err != nil {
					color.Red(err.Error())
					return
				}
				color.Green("all scripts reloaded")
				return
			}
			if err := scripts.ReloadExternalByPath(args[0]); err != nil {
				color.Red(err.Error())
				return
			}
			color.Green("script %s reloaded", args[0])
		},
	}
	// генерация позиционного комплитера
	carapace.Gen(scriptsReloadCmd).PositionalCompletion(externalScriptsCompleter())
	return scriptsReloadCmd
}

func scriptsCommand(c *console.Console) *cobra.Command {
	scriptsCmd := &cobra.Command{
		Use:                   "scripts",
		Short:                 "manage scripts",
		DisableFlagsInUseLine: true,
	}

	scriptsCmd.AddCommand(
		scriptsLoadCommand(c),
		scriptsListCommand(c),
		scriptsReloadCommand(c),
		scriptsRemoveCommand(c),
	)

	return scriptsCmd
}

func externalScriptsCompleter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var suggestions []string
		for _, v := range scripts.GetScripts() {
			suggestions = append(suggestions, v.GetPath())
		}
		return carapace.ActionValues(suggestions...)
	})
}

func scriptsLoadCommandCompleter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionFiles()
	})
}
