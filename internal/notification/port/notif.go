package port

import (
	"context"

	"github.com/babyhando/order-service/internal/common"
	"github.com/babyhando/order-service/internal/notification/domain"
)

type Service interface {
	Send(ctx context.Context, notif *domain.Notification) error
	common.OutboxHandler[domain.NotificationOutbox]
}
