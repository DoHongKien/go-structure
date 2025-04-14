package service

import (
	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/model/dto"
	"github.com/DoHongKien/go-structure/internal/model/dto/login"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/utils/auth"
	"go.uber.org/zap"
)

type IAuthService interface {
	Login(userRequest dto.UserRequest) (bool, error)
	LoginGenToken(userRequest dto.UserRequest) (login.LoginResponse, error)
}

type authService struct {
	repo repo.IUserRepo
}

func NewAuthService(repo repo.IUserRepo) IAuthService {
	return &authService{
		repo: repo,
	}
}

// Login implements IAuthService.
func (a *authService) Login(userRequest dto.UserRequest) (bool, error) {
	isLoginSuccess, err := a.repo.Login(userRequest)

	if err != nil {
		global.Logger.Error("Login error", zap.Error(err))
		return false, err
	}

	return isLoginSuccess, nil
}

// LoginGenToken implements IAuthService.
func (a *authService) LoginGenToken(userRequest dto.UserRequest) (login.LoginResponse, error) {
	isLoginSuccess, err := a.repo.Login(userRequest)

	if err != nil {
		global.Logger.Error("Login error", zap.Error(err))
		return login.LoginResponse{}, err
	}

	if isLoginSuccess {
		accessToken, err := auth.CreateAccessToken(userRequest.Username)
		if err != nil {
			global.Logger.Error("Create access token error", zap.Error(err))
			return login.LoginResponse{}, err
		}

		refreshToken, err := auth.CreateRefreshToken(userRequest.Username)
		if err != nil {
			global.Logger.Error("Create refresh token error", zap.Error(err))
			return login.LoginResponse{}, err
		}

		return login.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}

	return login.LoginResponse{}, nil
}
