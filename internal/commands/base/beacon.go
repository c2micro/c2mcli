package base

import (
	"fmt"

	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

// листинг биконов
func beaconListCommand(*console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "list",
		Short:                 "list beacons",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			beacons := beacon.Beacons.Get()
			for _, v := range beacons {
				os := v.GetOs().StringShort()
				if v.GetIsPrivileged() {
					os = color.RedString(v.GetOs().StringShort())
				}
				last := color.GreenString(utils.HumanDurationC(v.GetLast()))
				if v.IsDelay(0) {
					last = color.YellowString(utils.HumanDurationC(v.GetLast()))
				}
				if v.IsDead(0) {
					last = color.RedString(utils.HumanDurationC(v.GetLast()))
				}
				fmt.Printf("[%s] (%15s) %6s %-20s %-16s %s\n",
					os,
					last,
					v.GetIdHex(),
					v.GetUsername(),
					v.GetHostname(),
					v.GetIntIp(),
				)
			}
		},
	}
}

// работа с биконами
func beaconCommand(c *console.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "beacon",
		Short:                 "manage beacons",
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		beaconListCommand(c),
	)
	return cmd
}
