package models

import (
	"github.com/hiroykam/goa-sample/models/entities"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type RefreshTokenModel struct {
	Db *gorm.DB
}

func (m RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}

func NewRefreshTokenModel(db *gorm.DB) *RefreshTokenModel {
	return &RefreshTokenModel{
		Db: db,
	}
}

func (m *RefreshTokenModel) Add(s *entities.RefreshToken) *sample_error.SampleError {
	db := m.Db.Create(s)
	if db.Error != nil {
		return sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}
	return nil
}

func (m *RefreshTokenModel) GetByJti(jti string) (*entities.RefreshToken, *sample_error.SampleError) {
	var native entities.RefreshToken

	db := m.Db.Table(m.TableName()).Where("jti = ?", jti).First(&native)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, sample_error.NewSampleError(sample_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}

	return &native, nil
}

func (m *RefreshTokenModel) GetByUserId(userId int) (*entities.RefreshToken, *sample_error.SampleError) {
	var native entities.RefreshToken

	db := m.Db.Table(m.TableName()).Where("user_id = ?", userId).First(&native)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, sample_error.NewSampleError(sample_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}

	return &native, nil
}

func (m *RefreshTokenModel) Update(userId int, jti string) *sample_error.SampleError {
	db := m.Db.Table(m.TableName()).Where("user_id = ?", userId).Update("jti", jti)
	if db.Error != nil {
		return sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}
	return nil
}
