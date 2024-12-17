package beacon

import (
	"fmt"

	"github.com/c2micro/c2mcli/internal/storage/task"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

func tasksCommand(*console.Console) *cobra.Command {
	tasksCmd := &cobra.Command{
		Use:                   "tasks",
		Short:                 "show tasks for beacon",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range task.TaskGroups.Get() {
				timestamp := v.GetCreatedAt().Format("2006-01-02 15:04:05")
				fmt.Printf("[%s] (%4d) %s: %s\n",
					timestamp,
					v.GetId(),
					v.GetAuthor(),
					v.GetCmd(),
				)
			}
		},
	}

	return tasksCmd
}
