package order

import (
	"context"
	"order-service/internal/order/domain"
	"order-service/internal/order/port"
	userPort "order-service/internal/user/port"
	"time"
)

type service struct {
	userService userPort.Service
	repo        port.Repo
}

func NewService(userService userPort.Service, repo port.Repo) port.Service {
	return &service{
		userService: userService,
		repo:        repo,
	}
}

func (s *service) CreateOrder(ctx context.Context, order domain.Order) (domain.OrderID, error) {
	panic("not implemented")
}

func (s *service) SubmitOrder(ctx context.Context, orderID domain.OrderID, submittedAt time.Time) error {
	panic("not implemented")
}
