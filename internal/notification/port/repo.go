package port

import (
	"context"

	"github.com/babyhando/order-service/internal/notification/domain"
)

type Repo interface {
	Create(ctx context.Context, notif *domain.Notification) (domain.NotifID, error)
	CreateOutbox(ctx context.Context, outbox *domain.NotificationOutbox) error
}
