package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/babyhando/order-service/api/pb"
	notifDomain "github.com/babyhando/order-service/internal/notification/domain"
	notifPort "github.com/babyhando/order-service/internal/notification/port"
	"github.com/babyhando/order-service/internal/user"
	"github.com/babyhando/order-service/internal/user/domain"
	userPort "github.com/babyhando/order-service/internal/user/port"
	"github.com/babyhando/order-service/pkg/jwt"
	timeutils "github.com/babyhando/order-service/pkg/time"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	svc                   userPort.Service
	notifSvc              notifPort.Service
	authSecret            string
	expMin, refreshExpMin uint
}

func NewUserService(svc userPort.Service, authSecret string, expMin, refreshExpMin uint, notifSvc notifPort.Service) *UserService {
	return &UserService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
		notifSvc:      notifSvc,
	}
}

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserNotFound           = user.ErrUserNotFound
	ErrInvalidUserPassword    = errors.New("invalid password")
	ErrWrongOTP               = errors.New("wrong otp")
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

func (s *UserService) SendSignInOTP(ctx context.Context, phone string) error {
	user, err := s.svc.GetUserByFilter(ctx, &domain.UserFilter{
		Phone: phone,
	})

	if err != nil {
		return err
	}

	if user == nil {
		return nil
	}

	code := rand.IntN(999999) + 100000

	return s.notifSvc.Send(ctx, notifDomain.NewNotification(user.ID, fmt.Sprint(code), notifDomain.NotifTypeSMS, true, time.Minute*2))
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

	ok, err := s.notifSvc.CheckUserNotifValue(ctx, user.ID, req.GetOtp())
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, ErrWrongOTP
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
			ExpiresAt: jwt2.NewNumericDate(timeutils.AddMinutes(s.expMin, true)),
		},
		UserID: uint(userID),
	})
	if err != nil {
		return
	}

	refresh, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(timeutils.AddMinutes(s.refreshExpMin, true)),
		},
		UserID: uint(userID),
	})

	if err != nil {
		return
	}

	return
}
