package authService

import (
	"blog_api/app/auth/role"
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB *gorm.DB
}

func (authServ *AuthServiceImpl) Save(userReq *dto.UserRequest) (*model.User, error) {
	user := userReq.ToModel()
	user.Role = role.Member
	err := authServ.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	// if err := authServ.DB.Transaction(func(tx *gorm.DB) error {
	// 	if err := tx.Create(&user).Error; err != nil {
	// 		return echo.NewHTTPError(http.StatusBadGateway, err)
	// 	}
	// 	return nil
	// }); err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (authServ *AuthServiceImpl) GetByID(id uint) (*model.User, error) {
	return nil, nil
}

func (authServ *AuthServiceImpl) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := authServ.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("email = ?", email).First(&user).Error
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return nil, echo.NewHTTPError(http.StatusBadGateway, err)
		}
	}
	return &user, nil
}

func (authServ *AuthServiceImpl) UpdateByID(userReq *dto.UserRequest) (*model.User, error) {
	return nil, nil
}

func (authServ *AuthServiceImpl) SetDB(db *gorm.DB) {
	authServ.DB = db
}
