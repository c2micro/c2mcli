package beacon

import (
	"github.com/c2micro/c2mcli/internal/storage/task"
	"github.com/fatih/color"
	"github.com/reeflective/console"
)

// печать результатов таск группы
func printTaskGroupData(c *console.Console, v task.TaskData) {
	switch data := v.(type) {
	case *task.Message:
		c.Printf("%s\n", data.String())
	case *task.Task:
		preambule := data.StringStatus()
		if preambule != "" {
			c.Printf("%s\n", preambule)
		}
		if data.GetOutputLen() == 0 {
			return
		}
		if data.GetIsOutputBig() {
			c.Printf("[%s] %s %d %s\n", color.YellowString("!"), "output too big, use: task download", data.GetId(), "<path to save>")
			return
		}
		if data.GetIsBinary() {
			c.Printf("[%s] %s %d %s\n", color.YellowString("!"), "output is possible binary, use: task download", data.GetId(), "<path to save>")
			return
		}
		output := data.GetOutputString()
		if output != "" {
			c.Printf("%s\n", output)
		}
	}
}
