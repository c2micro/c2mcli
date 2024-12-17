package base

import (
	"fmt"

	"github.com/fatih/color"
)

func GetPrompt() string {
	return fmt.Sprintf("[%s] > ", color.CyanString("c2m"))
}
