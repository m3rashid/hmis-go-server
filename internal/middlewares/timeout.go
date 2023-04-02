package middlewares

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

var errorTimeOut = errors.New("server taking too long to respond")

func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return errorTimeOut
	case <-timer.C:
	}
	return nil
}

func serverTimeout(c *fiber.Ctx) error {
	sleepTime, _ := time.ParseDuration(c.Params("sleepTime") + "ms")
	if err := sleepWithContext(c.UserContext(), sleepTime); err != nil {
		return fmt.Errorf("%w: execution error", err)
	}
	return nil
}

// throws error if it takes more than 5 seconds
// to be used in the individual routes
var ServerTimeout = timeout.New(serverTimeout, 5*time.Second, errorTimeOut)
