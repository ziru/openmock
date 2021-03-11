package openmock

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func (om *OpenMock) startHTTP() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logrus.WithFields(logrus.Fields{
			"http_path":   c.Path(),
			"http_method": c.Request().Method,
			"http_host":   c.Request().Host,
			"http_req":    string(reqBody),
			"http_res":    string(resBody),
		}).Info()
	}))
	if om.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowCredentials: true,
			AllowHeaders:     []string{"*"},
			AllowMethods:     []string{"*"},
		}))
	}
	mocks := om.repo.HTTPMocks
	for h, ms := range mocks {
		func(h ExpectHTTP, ms MocksArray) {
			e.Match(
				[]string{h.Method},
				h.Path,
				func(ec echo.Context) error {
					body, _ := ioutil.ReadAll(ec.Request().Body)
					c := Context{
						HTTPContext:     ec,
						HTTPHeader:      ec.Request().Header,
						HTTPBody:        string(body),
						HTTPPath:        ec.Request().URL.String(),
						HTTPQueryString: ec.QueryString(),
						om:              om,
					}
					if !ms.DoActions(c) {
						return errors.New("no mock matched")
					}
					return nil
				},
			)
		}(h, ms)
	}

	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%d", om.HTTPHost, om.HTTPPort)),
	)
}
