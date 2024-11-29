package service

import (
	"context"
	"order-service/api/pb"
	"order-service/internal/user"
	"order-service/internal/user/domain"
	userPort "order-service/internal/user/port"
	"order-service/pkg/jwt"
	"order-service/pkg/time"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	svc                   userPort.Service
	authSecret            string
	expMin, refreshExpMin uint
}

func NewUserService(svc userPort.Service, authSecret string, expMin, refreshExpMin uint) *UserService {
	return &UserService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
	}
}

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserNotFound           = user.ErrUserNotFound
)

func (s *UserService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	userID, err := s.svc.CreateUser(ctx, domain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Phone:     domain.Phone(req.GetPhone()),
	})

	if err != nil {
		return nil, err
	}

	accessToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.expMin, true)),
		},
		UserID: uint(userID),
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.refreshExpMin, true)),
		},
		UserID: uint(userID),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserSignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*pb.User, error) {
	user, err := s.svc.GetUserByID(ctx, domain.UserID(id))
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        uint64(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     string(user.Phone),
	}, nil
}
