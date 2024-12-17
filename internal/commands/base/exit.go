package base

import (
	"os"

	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func exitCommand(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "exit",
		Short: "exit operator cli",
		Run: func(*cobra.Command, []string) {
			if utils.ExitConsole(c) {
				service.Close()
				os.Exit(0)
			}
		},
	}
}
