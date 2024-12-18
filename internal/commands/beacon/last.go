package beacon

import (
	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/c2micro/c2mcli/internal/storage/task"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func lastCommand(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "last",
		Aliases:               []string{"l"},
		Short:                 "get output of last task",
		DisableFlagsInUseLine: true,
		GroupID:               constants.CoreGroupId,
		Run: func(cmd *cobra.Command, args []string) {
			tg := task.TaskGroups.GetLast()
			if tg == nil {
				color.YellowString("no tasks")
				return
			}
			for _, v := range tg.GetData().Get() {
				printTaskGroupData(c, v)
			}
		},
	}
}
