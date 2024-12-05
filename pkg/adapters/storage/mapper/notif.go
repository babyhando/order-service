package mapper

import (
	"encoding/json"

	"github.com/babyhando/order-service/internal/common"
	"github.com/babyhando/order-service/internal/notification/domain"
	"github.com/babyhando/order-service/pkg/adapters/storage/types"
)

func NotifOutbox2Storage(no *domain.NotificationOutbox) (*types.Outbox, error) {
	data, err := json.Marshal(&no.Data)
	if err != nil {
		return nil, err
	}

	return &types.Outbox{
		Data:   data,
		RefID:  uint(no.NotifID),
		Type:   uint8(no.Type),
		Status: uint8(no.Status),
	}, nil
}

func Notification2Storage(no *domain.Notification) *types.Notification {
	return &types.Notification{
		Content: no.Content,
		To:      uint(no.UserID),
		Type:    uint8(no.Type),
	}
}

func OutboxStorage2Notif(outbox types.Outbox) (domain.NotificationOutbox, error) {
	var outboxData domain.OutboxData
	err := json.Unmarshal([]byte(outbox.Data), &outboxData)
	if err != nil {
		return domain.NotificationOutbox{}, err
	}

	return domain.NotificationOutbox{
		OutboxID: common.OutboxID(outbox.ID),
		NotifID:  domain.NotifID(outbox.RefID),
		Data:     outboxData,
		Status:   common.OutboxStatus(outbox.Status),
		Type:     common.OutboxType(outbox.Type),
	}, nil
}
