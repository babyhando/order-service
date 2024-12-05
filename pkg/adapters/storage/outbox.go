package storage

import (
	"context"

	"github.com/babyhando/order-service/internal/common"
	"gorm.io/gorm"
)

type outboxRepo struct {
	db *gorm.DB
}

func (o *outboxRepo) UpdateBulkStatuses(ctx context.Context, status common.OutboxStatus, ids ...common.OutboxID) error {
	return o.db.Exec("update outboxes set status = ? where id in ?", status, ids).Error
}

func (o *outboxRepo) UpdateStatus(ctx context.Context, status common.OutboxStatus, id common.OutboxID) error {
	return o.db.Exec("update outboxes set status = ? where id = ?", status, id).Error
}

func NewOutboxRepo(db *gorm.DB) common.OutboxRepo {
	return &outboxRepo{db}
}
