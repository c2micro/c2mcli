package beacon

import (
	"strings"

	"github.com/c2micro/c2mcli/internal/scripts"
	"github.com/c2micro/c2mcli/internal/scripts/aliases"
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func aliasCommands(*console.Console) []*cobra.Command {
	cmds := make([]*cobra.Command, 0)
	for k, v := range aliases.Aliases {
		cmds = append(cmds, &cobra.Command{
			Use:   k,
			Short: v.GetDescription(),
			Run: func(cmd *cobra.Command, args []string) {
				rawCmd := k + " " + strings.Join(args, " ")
				if err := scripts.ProcessCommand(beacon.ActiveBeacon.GetId(), rawCmd); err != nil {
					color.Red(err.Error())
				}
			},
		})
	}
	return cmds
}