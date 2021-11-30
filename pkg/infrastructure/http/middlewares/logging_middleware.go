package middlewares

import (
	"errors"
	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/gofiber/fiber/v2"
	lg "jarjarbinks/pkg/infrastructure/logging/interfaces"
	"strconv"
	"sync"
	"time"
)

func Logging(logger lg.Logger) fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		// Set variables
		var (
			once       sync.Once
			errHandler fiber.ErrorHandler
		)
		var errPadding = 15
		var errPaddingStr = strconv.Itoa(errPadding)
		once.Do(func() {
			stack := c.App().Stack()
			for m := range stack {
				for r := range stack[m] {
					if len(stack[m][r].Path) > errPadding {
						errPadding = len(stack[m][r].Path)
						errPaddingStr = strconv.Itoa(errPadding)
					}
				}
			}
			errHandler = c.App().ErrorHandler
		})

		var start, stop time.Time
		start = time.Now()
		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			logger.Error("an unexpected error occurred while processing request", chainErr, map[string]interface{}{
				"status":  c.Response().StatusCode(),
				"latency": stop.Sub(start).String(),
				"hostName": c.Hostname(),
				"path":c.Path(),
				"method": c.Method(),
				"url": c.OriginalURL(),
				"body": c.Body(),
				"bytesReceived": len(c.Request().Body()),
				"bytesSent": len(c.Response().Body()),
				"responseBody": c.Response().Body(),
				"queryString": c.Request().URI().QueryArgs().String(),
				"ip": c.IP(),
				"port": c.Port(),
				"referrer": c.Get(fiber.HeaderReferer),
				"userAgent": c.Get(fiber.HeaderUserAgent),
				"ips": c.Get(fiber.HeaderXForwardedFor),
			})
			if hub := sentryfiber.GetHubFromContext(c); hub != nil {
				hub.CaptureException(chainErr)
			}
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		stop = time.Now()
		if c.Response().StatusCode() >= 200 && c.Response().StatusCode() < 400{
			logger.Info("request processed successfully", map[string]interface{}{
				"status":  c.Response().StatusCode(),
				"latency": stop.Sub(start).String(),
				"hostName": c.Hostname(),
				"path":c.Path(),
				"method": c.Method(),
				"url": c.OriginalURL(),
				"body": c.Body(),
				"bytesReceived": len(c.Request().Body()),
				"bytesSent": len(c.Response().Body()),
				"responseBody": c.Response().Body(),
				"queryString": c.Request().URI().QueryArgs().String(),
				"ip": c.IP(),
				"port": c.Port(),
				"referrer": c.Get(fiber.HeaderReferer),
				"userAgent": c.Get(fiber.HeaderUserAgent),
				"ips": c.Get(fiber.HeaderXForwardedFor),
			})
		} else if c.Response().StatusCode() >= 400 && c.Response().StatusCode() < 500{
			logger.Warn("an unexpected error occurred while processing request", map[string]interface{}{
				"status":  c.Response().StatusCode(),
				"latency": stop.Sub(start).String(),
				"hostName": c.Hostname(),
				"path":c.Path(),
				"method": c.Method(),
				"url": c.OriginalURL(),
				"body": c.Body(),
				"bytesReceived": len(c.Request().Body()),
				"bytesSent": len(c.Response().Body()),
				"responseBody": c.Response().Body(),
				"queryString": c.Request().URI().QueryArgs().String(),
				"ip": c.IP(),
				"port": c.Port(),
				"referrer": c.Get(fiber.HeaderReferer),
				"userAgent": c.Get(fiber.HeaderUserAgent),
				"ips": c.Get(fiber.HeaderXForwardedFor),
			})
		}else{
			logger.Error("an unexpected error occurred while processing request", errors.New("an unexpected error occurred while processing request"), map[string]interface{}{
				"status":  c.Response().StatusCode(),
				"latency": stop.Sub(start).String(),
				"hostName": c.Hostname(),
				"path":c.Path(),
				"method": c.Method(),
				"url": c.OriginalURL(),
				"body": c.Body(),
				"bytesReceived": len(c.Request().Body()),
				"bytesSent": len(c.Response().Body()),
				"responseBody": c.Response().Body(),
				"queryString": c.Request().URI().QueryArgs().String(),
				"ip": c.IP(),
				"port": c.Port(),
				"referrer": c.Get(fiber.HeaderReferer),
				"userAgent": c.Get(fiber.HeaderUserAgent),
				"ips": c.Get(fiber.HeaderXForwardedFor),
			})
		}
		return nil
	}
}

