package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/babyhando/order-service/internal/common"
	"github.com/babyhando/order-service/internal/notification/domain"
	"github.com/babyhando/order-service/internal/notification/port"
	userDomain "github.com/babyhando/order-service/internal/user/domain"
	"github.com/babyhando/order-service/pkg/adapters/storage/mapper"
	"github.com/babyhando/order-service/pkg/adapters/storage/types"
	"github.com/babyhando/order-service/pkg/cache"
	"github.com/babyhando/order-service/pkg/conv"
	"gorm.io/gorm"
)

type notifRepo struct {
	db            *gorm.DB
	cacheProvider cache.Provider
}

func NewNotificationRepo(db *gorm.DB, cacheProvider cache.Provider) port.Repo {
	return &notifRepo{
		db:            db,
		cacheProvider: cacheProvider,
	}
}

func (r *notifRepo) Create(ctx context.Context, notif *domain.Notification) (domain.NotifID, error) {
	no := mapper.Notification2Storage(notif)
	if err := r.db.WithContext(ctx).Table("notifications").Create(no).Error; err != nil {
		return 0, err
	}

	if notif.ForValidation {
		if err := r.cacheProvider.Set(ctx, fmt.Sprintf("notif.%d", notif.UserID), notif.TTL, conv.ToBytes(notif.Content)); err != nil {
			return 0, err
		}
	}

	return domain.NotifID(no.ID), nil
}

func (r *notifRepo) CreateOutbox(ctx context.Context, no *domain.NotificationOutbox) error {
	outbox, err := mapper.NotifOutbox2Storage(no)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Table("outboxes").Create(outbox).Error
}

func (r *notifRepo) QueryOutboxes(ctx context.Context, limit uint, status common.OutboxStatus) ([]domain.NotificationOutbox, error) {
	var outboxes []types.Outbox

	err := r.db.WithContext(ctx).Table("outboxes").
		Where(`"type" = ?`, common.OutboxTypeNotif).
		Where("status = ?", status).
		Limit(int(limit)).Scan(&outboxes).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	result := make([]domain.NotificationOutbox, len(outboxes))

	for i := range outboxes {
		v, err := mapper.OutboxStorage2Notif(outboxes[i])
		if err != nil {
			return nil, err
		}
		result[i] = v
	}

	return result, nil
}

func (r *notifRepo) GetUserNotifValue(ctx context.Context, userID userDomain.UserID) (string, error) {
	v, err := r.cacheProvider.Get(ctx, fmt.Sprintf("notif.%d", userID))
	return conv.ToStr(v), err
}
