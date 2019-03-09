package models

import (
	"github.com/hiroykam/goa-sample/models/entities"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func (m UserModel) TableName() string {
	return "users"
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) GetWithEmail(email string,) (*entities.User, *sample_error.SampleError) {
	var native entities.User

	db := m.db.Table(m.TableName()).Where("email = ?", email).First(&native)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, sample_error.NewSampleError(sample_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}

	return &native, nil
}

func (m *UserModel) Get(id int) (*entities.User, *sample_error.SampleError) {
	var native entities.User

	db := m.db.Table(m.TableName()).Where("id = ?", id).First(&native)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, sample_error.NewSampleError(sample_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}

	return &native, nil
}
