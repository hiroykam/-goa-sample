package services

import (
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/models"
	"github.com/hiroykam/goa-sample/models/entities"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type SampleService struct {
	model *models.SampleModel
}

func NewSampleService(db *gorm.DB) *SampleService {
	return &SampleService{
		model: models.NewSampleModel(db),
	}
}

func convertSample(e *entities.Sample) *app.Sample {
	return &app.Sample{
		ID: e.ID,
		UserID: e.UserID,
		Name: e.Name,
		Detail: e.Detail,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func convertSamples(e *entities.Sample) *app.Samples {
	return &app.Samples{
		ID: e.ID,
		Name: e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func convertSamplesCollection(es []*entities.Sample) app.SamplesCollection {
	objs := app.SamplesCollection{}

	for _, e := range es {
		objs = append(objs, convertSamples(e))
	}
	return objs
}

func (s *SampleService) GetSamples(userId int) (app.SamplesCollection, *sample_error.SampleError) {
	r, err := s.model.List(userId)
	if err != nil {
		return nil, err
	}
	obj := convertSamplesCollection(r)

	return obj, nil
}

func (s *SampleService) Add(userId int, name string, detail string) (*app.Sample, *sample_error.SampleError) {
	e := &entities.Sample{}
	e.UserID = userId
	e.Name = name
	e.Detail = detail

	err := s.model.Add(e)
	if err != nil {
		return nil, err
	}
	return convertSample(e), nil
}

func (s *SampleService) Show(id int) (*app.Sample, *sample_error.SampleError) {
	e, err := s.model.Get(id)
	if err != nil {
		return nil, err
	}
	return convertSample(e), nil
}

func (s *SampleService) Update(id int, userId int, name string, detail string) *sample_error.SampleError {
	e := &entities.Sample{}
	e.ID = id
	e.UserID = userId
	e.Name = name
	e.Detail = detail
	err := s.model.Update(e)
	if err != nil {
		return err
	}
	return nil
}

func (s *SampleService) Delete(id int) *sample_error.SampleError {
	err := s.model.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
