package beacon

import (
	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func exitCommand(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "exit",
		Short:                 "switch back on base console",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := service.UnpollBeaconTasks(beacon.ActiveBeacon); err != nil {
				color.Yellow("unable stop polling tasks for beacon: %s", err.Error())
			}
			beacon.ActiveBeacon = nil
			c.SwitchMenu(constants.BaseMenuName)
		},
	}
}