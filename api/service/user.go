package service

import (
	"context"
	"errors"
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
	ErrInvalidUserPassword    = errors.New("invalid password")
)

func (s *UserService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	userID, err := s.svc.CreateUser(ctx, domain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Phone:     domain.Phone(req.GetPhone()),
		Password:  domain.NewPassword(req.GetPassword()),
	})

	if err != nil {
		return nil, err
	}

	access, refresh, err := s.createTokens(uint(userID))
	if err != nil {
		return nil, err
	}

	return &pb.UserSignUpResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *UserService) SignIn(ctx context.Context, req *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	user, err := s.svc.GetUserByFilter(ctx, &domain.UserFilter{
		Phone: req.GetPhone(),
	})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if !user.PasswordIsCorrect(req.GetPassword()) {
		return nil, ErrInvalidUserPassword
	}

	access, refresh, err := s.createTokens(uint(user.ID))
	if err != nil {
		return nil, err
	}

	return &pb.UserSignInResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *UserService) createTokens(userID uint) (access, refresh string, err error) {
	access, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.expMin, true)),
		},
		UserID: uint(userID),
	})
	if err != nil {
		return
	}

	refresh, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.refreshExpMin, true)),
		},
		UserID: uint(userID),
	})

	if err != nil {
		return
	}

	return
}
