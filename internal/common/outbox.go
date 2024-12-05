package common

import (
	"context"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type OutboxStatus uint8

const (
	OutboxStatusCreated OutboxStatus = iota + 1
	OutboxStatusPicked
	OutboxStatusDone
	OutboxStatusFailed
)

type OutboxType uint8

const (
	OutboxTypeNotif OutboxType = iota + 1
)

type OutboxHandler[T any] interface {
	Handle(ctx context.Context, outboxes []T) error
	Query(ctx context.Context) ([]T, error)
	Interval() time.Duration
}

type OutboxRunner[T any] struct {
	handler   OutboxHandler[T]
	scheduler gocron.Scheduler
}

func RegisterOutboxRunner[T any](handler OutboxHandler[T], scheduler gocron.Scheduler) {
	runner := &OutboxRunner[T]{
		handler:   handler,
		scheduler: scheduler,
	}

	runner.register()
}

func (o *OutboxRunner[T]) register() {
	o.scheduler.NewJob(
		gocron.DurationJob(o.handler.Interval()),
		gocron.NewTask(func() { // poller logic
			ctx := context.Background() // todo : can be configurable
			// todo : logger should be injected, *zap.Logger
			outboxes, err := o.handler.Query(context.Background())
			if err != nil {
				log.Println("failed to fetch outboxes, err ", err.Error())
				return
			}

			if err := o.handler.Handle(ctx, outboxes); err != nil {
				log.Println("failed to handle outboxes, err", err.Error())
			}
		}),
	)
}

type OutboxID uint

type OutboxRepo interface {
	UpdateStatus(ctx context.Context, status OutboxStatus, id OutboxID) error
	UpdateBulkStatuses(ctx context.Context, status OutboxStatus, ids ...OutboxID) error
}
