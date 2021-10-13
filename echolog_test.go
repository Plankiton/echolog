package echolog

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Perm calls f with each permutation of a.
func Perm(a []interface{}, f func([]interface{})) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []interface{}, f func([]interface{}), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func MakeLogTest(t *testing.T) {
	testCases := []interface{}{
		log.JSON{
			"joao":  "maria",
			"maria": "joao",
		},
		"joao", "maria",
		struct {
			Joao  string
			Maria string
		}{
			Joao:  "maria",
			Maria: "joao",
		},
		[]interface{}{
			log.JSON{
				"joao":  "maria",
				"maria": "joao",
			},
			"joao", "maria",
			struct {
				Joao  string
				Maria string
			}{
				Joao:  "maria",
				Maria: "joao",
			},
		},
	}

	e := echo.New()
	e.Use(middleware.RequestID())

	const path = "/"
	e.GET(path, func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		req_id := req.Header.Get(echo.HeaderXRequestID)
		if req_id == "" {
			req_id = res.Header().Get(echo.HeaderXRequestID)
		}
		Perm(testCases, func(a []interface{}) {
			o := makeLog(c, a)
			if log_id, ok := o["id"]; ok {
				var err error
				if ok := !(log_id == req_id); ok {
					err = fmt.Errorf("request id is wrong, %s == %s : %v", log_id, req_id, ok)
				}

				if err != nil {
					t.Errorf("makeLog: %v", err)
				}
			} else {
				err := fmt.Errorf("missing request id in log maker return")
				t.Errorf("makeLog: %v", err)
			}
		})

		return nil
	})

	const url = ":6660"
	go e.Start(url)

	_, err := http.Get(url)
	if err != nil {
		t.Errorf("http: %v", err)
	}
}
