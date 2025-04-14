package repo

import (
	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/model"
	"github.com/DoHongKien/go-structure/internal/model/dto"
	"github.com/DoHongKien/go-structure/internal/model/dto/login"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Login(userRequest dto.UserRequest) (bool, error)
	LoginWithEmail(userRequest dto.UserRequest) (login.LoginResponse, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewAuthRepository() IUserRepo {
	return &userRepo{
		db: global.Mdb,
	}
}

// Login implements UserRepo.
func (u *userRepo) Login(userRequest dto.UserRequest) (bool, error) {
	var user model.User

	err := u.db.Where("username = ?", userRequest.Username).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// LoginWithEmail implements UserRepo.
func (u *userRepo) LoginWithEmail(userRequest dto.UserRequest) (login.LoginResponse, error) {
	panic("unimplemented")
}
