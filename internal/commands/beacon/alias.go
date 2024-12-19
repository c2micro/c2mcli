package beacon

import (
	"fmt"
	"strings"

	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/c2micro/c2mcli/internal/scripts"
	"github.com/c2micro/c2mcli/internal/scripts/aliases"
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

// регистрация алиасов
func aliasCommands(*console.Console) []*cobra.Command {
	cmds := make([]*cobra.Command, 0)
	for k, v := range aliases.Aliases {
		cmd := &cobra.Command{
			Use:                   k,
			Short:                 v.GetDescription(),
			GroupID:               constants.AliasGroupId,
			DisableFlagsInUseLine: true,
			DisableFlagParsing:    true,
			Run: func(cmd *cobra.Command, args []string) {
				rawCmd := k + " " + strings.Join(args, " ")
				if err := scripts.ProcessCommand(beacon.ActiveBeacon.GetId(), rawCmd); err != nil {
					color.Red(err.Error())
				}
			},
		}
		cmd.SetHelpTemplate(fmt.Sprintf("%s\n\n%s\n", v.GetDescription(), v.GetUsage()))
		cmds = append(cmds, cmd)
	}
	return cmds
}
