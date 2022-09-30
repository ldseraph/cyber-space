package main

import (
	"net/http"
	_ "net/http/pprof"
	"video-intelligence/cmd"
	_ "video-intelligence/docs"
	_ "video-intelligence/migrations"
)

// @title           Video Intelligence
// @version         0.1

// @description     Video Intelligence
// @termsOfService  http://swagger.io/terms/

// @contact.name   beijingyingnuoweixun
// @contact.url    https://e.gitee.com/beijingyingnuoweixun

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8090
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	//pprof
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	cmd.Execute()
}
