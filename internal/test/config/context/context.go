package context

import (
	"context"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)

	return ctx, cancel
}
