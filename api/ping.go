package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// @Summary     ping
// @Description ping 接口
// @Tags        channels
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       id         path     string true  "设备id"
// @Param       memo       formData string false "通道备注"
// @Param       streamtype formData string false "播放类型，pull 媒体服务器拉流，push 摄像头推流,默认push"
// @Param       url        formData string false "静态拉流地址，streamtype=pull 时生效。"
// @Success     0          {object} string
// @Failure     1000    {object} string
// @Failure     1001    {object} string
// @Failure     1002    {object} string
// @Failure     1003    {object} string
// @Router      /devices/{id}/channels [post]
func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"ping": "pong",
	})
}
