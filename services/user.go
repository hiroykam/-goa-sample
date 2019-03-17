package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/models"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	model *models.UserModel
	Auth  *AuthSharedService
}

func NewUserService(db *gorm.DB) (*UserService, *sample_error.SampleError) {
	s, err := NewAuthSharedService()
	if err != nil {
		return nil, err
	}

	return &UserService{
		model: models.NewUserModel(db),
		Auth:  s,
	}, nil
}

func (s *UserService) AuthWithEmailAndPassword(email, password string) (*app.Auth, *sample_error.SampleError) {
	tx := s.model.Db.Begin()

	h, err := NewHashedRefreshTokenService(tx)
	if err != nil {
		return nil, err
	}

	u, err := s.model.GetWithEmail(email, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = Confirm(u.HashedPassword, password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	a, jti, err := s.Auth.IssueTokens(u.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = h.AddOrUpdate(u.ID, jti)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	fmt.Println("test")

	return a, nil
}

func (s *UserService) AuthWithToken(token *jwt.Token) (*int, *sample_error.SampleError) {
	id, err := s.Auth.GetId(token)
	if err != nil {
		return nil, err
	}

	_, err = s.model.Get(id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
