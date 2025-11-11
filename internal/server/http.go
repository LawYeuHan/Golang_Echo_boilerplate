package server

import (
	"context"
	"ecpos/internal/handler"
	"ecpos/pkg/log"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServerHTTP(logger *log.Log, userHandler handler.UserHandler) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Debug = false
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP:  true,
		LogHost:      true,
		LogUserAgent: true,
		LogLatency:   true,
		LogStatus:    true,
		LogRequestID: true,
		LogURI:       true,
		LogError:     true,
		HandleError:  true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l := HTTPLogger(logger, c, v)
			if v.Error == nil {
				l.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST")
			} else {
				l.Log(context.Background(), slog.LevelError, "REQUEST_ERROR", "err", v.Error)
			}
			return nil
		},
	}))

	e.GET("/user/:id", userHandler.GetUserByID)
	e.GET("/userWithError/:id", userHandler.GetUserByIDWithError)

	return e
}

func HTTPLogger(logger *log.Log, c echo.Context, v middleware.RequestLoggerValues) *log.Log {
	slogger := logger.With(
		"request_id", c.Response().Header().Get(echo.HeaderXRequestID),
		"uri", v.URI,
		"status", v.Status,
		"latency", v.Latency.Nanoseconds(),
	)
	return &log.Log{Logger: slogger}
}
