package beacon

import (
	"github.com/c2micro/c2mcli/internal/storage/task"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func lastCommand(c *console.Console) *cobra.Command {
	return &cobra.Command{
		Use:                   "last",
		Short:                 "get output of last task",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			tg := task.TaskGroups.GetLast()
			if tg == nil {
				color.YellowString("no tasks")
				return
			}
			for _, v := range tg.GetData().Get() {
				switch data := v.(type) {
				case *task.Message:
					c.Printf("%s\n", data.String())
				case *task.Task:
					preambule := data.StringStatus()
					if preambule != "" {
						c.Printf("%s\n", preambule)
					}
					if data.GetIsOutputBig() {
						color.YellowString("output of task is to big. Download it directly (TODO)")
						continue
					}
					output := data.GetOutputString()
					if output != "" {
						c.Printf("%s\n", output)
					}
				}
			}
		},
	}
}
