package rtsp

import (
	"video-intelligence/cmd"

	"github.com/pocketbase/pocketbase/core"
)

func init() {
	cmd.AddInit(func(e *core.ServeEvent) error {
		dao = e.App.Dao()
		return nil
	})
}

// func initZlmReg() {

// }
