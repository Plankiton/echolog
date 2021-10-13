package echolog

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func makeLog(c echo.Context, msg ...interface{}) log.JSON {
	req := c.Request()
	res := c.Response()

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}

	output := log.JSON{"id": id}
	output["message"] = msg

	return output
}

// Log logs json encoded information
func Log(c echo.Context, msg ...interface{}) {
	output := makeLog(c, msg...)
	output["level"] = "info"
	c.Logger().Printj(output)
}

// Err logs json encoded error
func Err(c echo.Context, msg ...interface{}) {
	output := makeLog(c, msg...)
	output["level"] = "error"
	c.Logger().Printj(output)
}

// War logs json encoded wanring
func War(c echo.Context, msg ...interface{}) {
	output := makeLog(c, msg...)
	output["level"] = "error"
	c.Logger().Printj(output)
}
