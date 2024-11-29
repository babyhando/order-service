package storage

import (
	"context"
	"log"
	"order-service/internal/user/domain"
	userPort "order-service/internal/user/port"
	"order-service/pkg/cache"
	"strconv"
)

type userCachedRepo struct {
	repo     userPort.Repo
	provider cache.Provider
}

func (r *userCachedRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	uId, err := r.repo.Create(ctx, userDomain)
	if err != nil {
		return 0, err
	}
	userDomain.ID = uId

	oc := cache.NewJsonObjectCacher[*domain.User](r.provider)
	if err := oc.Set(ctx, r.userIDKey(uId), 0, &userDomain); err != nil {
		log.Println("error on caching (SET) user with id :", uId)
	}

	return uId, nil
}

func (r *userCachedRepo) userIDKey(userID domain.UserID) string {
	return "users" + strconv.FormatUint(uint64(userID), 10)
}

func (r *userCachedRepo) GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	oc := cache.NewJsonObjectCacher[*domain.User](r.provider)

	key := r.userIDKey(userID)
	dUser, err := oc.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if dUser != nil && dUser.ID > 0 {
		log.Println("reading user from cache , ID : ", userID)
		return dUser, nil
	}

	dUser, err = r.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err := oc.Set(ctx, key, 0, dUser); err != nil {
		log.Println("error on caching (SET) user with id :", userID)
	}

	return dUser, nil
}
