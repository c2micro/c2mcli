package base

import (
	"fmt"
	"strconv"

	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func useCommand(c *console.Console) *cobra.Command {
	useCmd := &cobra.Command{
		Use:                   "use",
		Short:                 "switch on beacon shell",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseInt(args[0], 16, 32)
			if err != nil {
				color.Red("invalid beacon id")
				return
			}
			b := beacon.Beacons.GetById(uint32(id))
			if b == nil {
				color.Red("unknown beacon id")
				return
			}
			if err := service.PollBeaconTasks(b); err != nil {
				color.Red("unable start polling tasks for beacon: %s", err.Error())
				return
			}
			beacon.ActiveBeacon = b
			c.Menu(constants.BeaconMenuName).Prompt().Primary = func() string { return fmt.Sprintf("[%s] > ", color.MagentaString(args[0])) }
			c.SwitchMenu(constants.BeaconMenuName)
		},
	}
	// генерация позиционного комплитера
	carapace.Gen(useCmd).PositionalCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var suggestions []string
		for _, v := range beacon.Beacons.Get() {
			suggestions = append(suggestions, v.GetIdHex())
		}
		return carapace.ActionValues(suggestions...)
	}))
	return useCmd
}
