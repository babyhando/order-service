package cache

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type Provider interface {
	Set(ctx context.Context, key string, ttl time.Duration, data []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
}

type SerializationType uint8

const (
	SerializationTypeUnknown = iota
	SerializationTypeJSON
	SerializationTypeGob
)

type ObjectCacher[T any] struct {
	provider          Provider
	serializationType SerializationType
}

const KEY_PREFIX = "ORDER_SERVICE"

func createKey(k string) string {
	return KEY_PREFIX + "." + k
}

func NewObjectCacher[T any](p Provider, st SerializationType) *ObjectCacher[T] {
	return &ObjectCacher[T]{
		provider:          p,
		serializationType: st,
	}
}

func NewJsonObjectCacher[T any](p Provider) *ObjectCacher[T] {
	return NewObjectCacher[T](p, SerializationTypeJSON)
}

func (c *ObjectCacher[T]) Get(ctx context.Context, key string) (T, error) {
	t := new(T)
	data, err := c.provider.Get(ctx, createKey(key))
	if err != nil {
		if errors.Is(err, ErrCacheMiss) {
			return *t, nil
		}
		return *t, err
	}

	return *t, c.unmarshal(data, &t)
}

func (c *ObjectCacher[T]) Del(ctx context.Context, key string) error {
	return c.provider.Del(ctx, createKey(key))
}

func (c *ObjectCacher[T]) Set(ctx context.Context, key string, ttl time.Duration, in T) error {
	data, err := c.Marshal(in)
	if err != nil {
		return err
	}

	return c.provider.Set(ctx, createKey(key), ttl, data)
}
