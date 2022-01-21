package echolog

import (
	"github.com/labstack/echo/v4"
	"fmt"
)

func makeLog(c echo.Context, msg ...interface{}) log.JSON {
	req := c.Request()
	res := c.Response()

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}

	output := log.JSON{
		"id": id,
		"message": fmt.Sprint(msg),
	}

	return output
}

// Log logs json encoded information
func Log(c echo.Context, msg ...interface{}) {
	output := makeLog(c, msg...)
	c.Logger().Infoj(output)
}

// Err logs json encoded error
func Err(c echo.Context, msg ...interface{}) {
	output := makeLog(c, msg...)
	c.Logger().Errorj(output)
}

// War logs json encoded wanring
func War(c echo.Context, msg ...interface{}) {
	output := makeLog(c, msg...)
	c.Logger().Debugj(output)
}
