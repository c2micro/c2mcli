package base

import (
	"github.com/c2micro/c2mcli/internal/scripts"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

// регистрация скрипта
func scriptLoadCommand(*console.Console) *cobra.Command {
	cmd := &cobra.Command{
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
	carapace.Gen(cmd).PositionalCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionFiles()
	}))
	return cmd
}

// листинг скриптов
func scriptListCommand(c *console.Console) *cobra.Command {
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

// удаление скриптов
func scriptRemoveCommand(*console.Console) *cobra.Command {
	cmd := &cobra.Command{
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
	carapace.Gen(cmd).PositionalCompletion(externalScriptsCompleter())
	return cmd
}

// перезагрузка скриптов
func scriptReloadCommand(*console.Console) *cobra.Command {
	cmd := &cobra.Command{
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
	carapace.Gen(cmd).PositionalCompletion(externalScriptsCompleter())
	return cmd
}

// работа со скриптам
func scriptCommand(c *console.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "scripts",
		Short:                 "manage scripts",
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		scriptLoadCommand(c),
		scriptListCommand(c),
		scriptReloadCommand(c),
		scriptRemoveCommand(c),
	)
	return cmd
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
