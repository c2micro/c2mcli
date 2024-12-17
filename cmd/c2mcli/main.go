package main

import (
	"context"
	"os"
	"os/signal"
	"slices"

	"github.com/c2micro/c2mcli/cmd/c2mcli/internal/cmd"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/zapcfg"
	"github.com/fatih/color"
	"github.com/go-faster/sdk/zctx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	// создание логгера
	lg, err := zapcfg.New().Build()
	if err != nil {
		panic(err)
	}

	// flush функция для закрытия активных объектов
	flush := func() {
		// игнорируем ошибку: /dev/stderr: invalid argument
		_ = lg.Sync()
	}
	defer flush()

	// замена os.Exit
	exit := func(code int) {
		flush()
		os.Exit(code)
	}

	// выход из паники
	defer func() {
		if r := recover(); r != nil {
			lg.Fatal("recovered from panic", zap.Any("panic", r))
			exit(2)
		}
	}()

	// инициализация приложения
	app := cmd.App{}
	ctx, cancel := signal.NotifyContext(zctx.Base(context.Background(), lg), os.Interrupt)
	defer cancel()

	root := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,

		Use:   "c2mcli",
		Short: "c2m operator cli",
		Long:  "c2m operator cli",
		Args:  cobra.NoArgs,

		RunE: app.Run,

		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// если в имени команды не содержаться определенные подкоманды -> процессинг глобальных флагов
			if !slices.Contains([]string{
				"help",
			}, cmd.Name()) {
				// валидация глобальных флагов
				if err = app.Validate(); err != nil {
					return err
				}
				// создание GRPC клиента
				if err = service.Init(cmd.Context(), app.Host, app.Token); err != nil {
					return err
				}
			}
			return nil
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			flush()
		},
	}

	root.CompletionOptions.DisableDefaultCmd = true
	app.RegisterFlags(root.PersistentFlags())

	if err = root.ExecuteContext(ctx); err != nil {
		color.Red("%s", err.Error())
		exit(2)
	}
}
