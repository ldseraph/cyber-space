package cmd

import (
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	"video-intelligence/api"
	"video-intelligence/pkg/swagger"
	"video-intelligence/ui"

	"github.com/pocketbase/pocketbase"
	pocketbaseCmd "github.com/pocketbase/pocketbase/cmd"
	"github.com/pocketbase/pocketbase/core"

	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var pb *pocketbase.PocketBase

func init() {
	pb = pocketbase.New()

	pb.RootCmd.AddCommand(pocketbaseCmd.NewServeCommand(pb, true))
	pb.RootCmd.AddCommand(pocketbaseCmd.NewMigrateCommand(pb))
	pb.RootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := pb.Bootstrap(); err != nil {
			return err
		}
		go func() {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit
			if err := pb.ResetBootstrapState(); err != nil {
				log.Fatal().Err(err).Msg("bootstrap reset err")
			}
			os.Exit(0)
		}()
		return nil
	}
	pb.RootCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		return pb.ResetBootstrapState()
	}

	AddRoute(func(e *echo.Echo) error {
		e.GET("/swagger/*", swagger.WrapHandler)
		return nil
	})

	AddRoute(api.Init)
	AddStaticFS(ui.DistDirFS)

	AddCommand(pb.RootCmd)
}

// AddRoute Add OnBeforeServe Route Hook
func AddRoute(fn func(e *echo.Echo) error) {
	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		fn(e.Router)
		return nil
	})
}

// AddStaticFS Add Static FS Route
func AddStaticFS(f fs.FS) {
	AddRoute(func(e *echo.Echo) error {
		e.StaticFS("/public", f)
		e.FileFS("/", "index.html", f)
		return nil
	})
}
