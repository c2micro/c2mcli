package beacon

import (
	"fmt"
	"strconv"

	"github.com/c2micro/c2mcli/internal/storage/task"
	"github.com/fatih/color"
	"github.com/reeflective/console"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

// вывод результатов в таск группе
func tasksGetCommand(c *console.Console) *cobra.Command {
	tasksGetCmd := &cobra.Command{
		Use:                   "get",
		Short:                 "get output for task",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
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
	// генерация позиционного комплитера
	carapace.Gen(tasksGetCmd).PositionalCompletion(tasksCommandCompleter())
	return tasksGetCmd
}

func tasksCommand(c *console.Console) *cobra.Command {
	tasksCmd := &cobra.Command{
		Use:                   "tasks",
		Short:                 "show tasks for beacon",
		Aliases:               []string{"t"},
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range task.TaskGroups.Get() {
				timestamp := v.GetCreatedAt().Format("01/02 15:04:05")
				fmt.Printf("[%s] (%4d) %s: %s\n",
					timestamp,
					v.GetId(),
					v.GetAuthor(),
					v.GetCmd(),
				)
			}
		},
	}

	tasksCmd.AddCommand(
		// get
		tasksGetCommand(c),
	)

	return tasksCmd
}

func tasksCommandCompleter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var suggestions []string
		for _, v := range task.TaskGroups.Get() {
			suggestions = append(suggestions, strconv.Itoa(int(v.GetId())))
		}
		return carapace.ActionValues(suggestions...)
	})
}
