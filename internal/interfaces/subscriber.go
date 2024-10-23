package interfaces

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/subscriber"
)

// ISubscriber interface.
type ISubscriber interface {
	Subscribe(context.Context, string, ...subscriber.HandlerFunction) error
	Close() error
}
