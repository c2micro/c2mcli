package beacon

import (
	"strconv"

	"github.com/c2micro/c2mcli/internal/constants"
	"github.com/c2micro/c2mcli/internal/storage/task"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func commandCommand(c *console.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "command",
		Short:                 "show commands for beacon",
		Aliases:               []string{"t"},
		DisableFlagsInUseLine: true,
		GroupID:               constants.CoreGroupId,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				for _, v := range task.TaskGroups.Get() {
					timestamp := v.GetCreatedAt().Format("01/02 15:04:05")
					c.Printf("[%s] (%d) %s: %s\n",
						timestamp,
						v.GetId(),
						v.GetAuthor(),
						v.GetCmd(),
					)
				}
				return
			}
			// листинг команд (таск групп)
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				color.Red("invalid task id")
				return
			}
			// получение таск группы
			tg := task.TaskGroups.GetById(id)
			if tg == nil {
				color.Red("unknown task id")
				return
			}
			for _, v := range tg.GetData().Get() {
				printTaskGroupData(c, v)
			}
		},
	}
	carapace.Gen(cmd).PositionalCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var suggestions []string
		for _, v := range task.TaskGroups.Get() {
			suggestions = append(suggestions, strconv.Itoa(int(v.GetId())))
		}
		return carapace.ActionValues(suggestions...)
	}))
	return cmd
}
