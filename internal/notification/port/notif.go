package port

import (
	"context"

	"github.com/babyhando/order-service/internal/common"
	"github.com/babyhando/order-service/internal/notification/domain"
	userDomain "github.com/babyhando/order-service/internal/user/domain"
)

type Service interface {
	Send(ctx context.Context, notif *domain.Notification) error
	CheckUserNotifValue(ctx context.Context, userID userDomain.UserID, val string) (bool, error)
	common.OutboxHandler[domain.NotificationOutbox]
}
