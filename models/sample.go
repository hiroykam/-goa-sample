package models

import (
	"github.com/hiroykam/goa-sample/models/entities"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type SampleModel struct {
	db *gorm.DB
}

func (m SampleModel) TableName() string {
	return "samples"
}

func NewSampleModel(db *gorm.DB) *SampleModel {
	return &SampleModel{
		db: db,
	}
}

func (m *SampleModel) Get(id int) (*entities.Sample, *sample_error.SampleError) {
	var native entities.Sample

	db := m.db.Table(m.TableName()).Where("id = ?", id).First(&native)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, sample_error.NewSampleError(sample_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}

	return &native, nil
}

func (m *SampleModel) List(userId int) ([]*entities.Sample, *sample_error.SampleError) {
	var objs []*entities.Sample

	db := m.db.Table(m.TableName()).Where("user_id = ?", userId).Find(&objs)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, sample_error.NewSampleError(sample_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}

	return objs, nil
}

func (m *SampleModel) Add(s *entities.Sample) *sample_error.SampleError {
	db := m.db.Create(s)
	if db.Error != nil {
		return sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}
	return nil
}

func (m *SampleModel) Update(s *entities.Sample) *sample_error.SampleError {
	db := m.db.Table(m.TableName()).Where("id = ?", s.ID).UpdateColumn(s)
	if db.Error != nil {
		return sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}
	return nil
}

func (m *SampleModel) Delete(id int) *sample_error.SampleError {
	var obj entities.Sample
	db := m.db.Delete(&obj, id)
	if db.Error != nil {
		return sample_error.NewSampleError(sample_error.InternalError, db.Error.Error())
	}
	return nil
}
