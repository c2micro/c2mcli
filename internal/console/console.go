package console

import (
	"context"
	"io"
	"os"

	baseCommands "github.com/c2micro/c2mcli/internal/commands/base"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/reeflective/console"
)

func Run(ctx context.Context) error {
	app := console.New("c2mcli")
	base := app.ActiveMenu()
	base.Short = "base operator cli"
	// кастомный промпт
	base.Prompt().Primary = baseCommands.GetPrompt
	// хендлер на обработку Ctrl+D
	base.AddInterrupt(io.EOF, func(c *console.Console) {
		if utils.ExitConsole(c) {
			service.Close()
			os.Exit(0)
		}
	})
	// добавление базовых команд
	base.SetCommands(baseCommands.Commands(app))
	return app.StartContext(ctx)
}