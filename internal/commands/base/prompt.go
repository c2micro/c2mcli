package base

import (
	"fmt"

	"github.com/fatih/color"
)

func GetPrompt() string {
	return fmt.Sprintf("[%s] > ", color.MagentaString("c2m"))
}
