package contexts

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	logger = "logger"
)

// SetLogger はロガーをコンテキストに追加します。
func SetLogger(c echo.Context, log *logrus.Entry) {
	c.Set(logger, log)
}

// Logger はロガーをコンテキストから取り出します。
func Logger(c echo.Context) *logrus.Entry {
	v := c.Get(logger)
	if v == nil {
		return nil
	}
	return v.(*logrus.Entry)
}
