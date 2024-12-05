package port

import (
	"context"

	"github.com/babyhando/order-service/internal/common"
	"github.com/babyhando/order-service/internal/notification/domain"
	userDomain "github.com/babyhando/order-service/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, notif *domain.Notification) (domain.NotifID, error)
	CreateOutbox(ctx context.Context, outbox *domain.NotificationOutbox) error
	QueryOutboxes(ctx context.Context, limit uint, status common.OutboxStatus) ([]domain.NotificationOutbox, error)
	GetUserNotifValue(ctx context.Context, userID userDomain.UserID) (string, error)
}
