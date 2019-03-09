package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/models"
	"github.com/hiroykam/goa-sample/models/entities"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

type SampleService struct {
	model *models.SampleModel
	User  *UserService
}

func NewSampleService(db *gorm.DB) (*SampleService, *sample_error.SampleError) {
	u, err := NewUserService(db)
	if err != nil {
		return nil, err
	}

	return &SampleService{
		model: models.NewSampleModel(db),
		User:  u,
	}, nil
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

func (s *SampleService) GetSamples(token *jwt.Token) (app.SamplesCollection, *sample_error.SampleError) {
	id, err := s.User.AuthWithToken(token)
	if err != nil {
		return nil, err
	}

	r, err := s.model.List(*id)
	if err != nil {
		return nil, err
	}
	obj := convertSamplesCollection(r)

	return obj, nil
}

func (s *SampleService) Add(token *jwt.Token, name string, detail string) (*app.Sample, *sample_error.SampleError) {
	id, err := s.User.AuthWithToken(token)
	if err != nil {
		return nil, err
	}

	e := &entities.Sample{}
	e.UserID = *id
	e.Name = name
	e.Detail = detail

	err = s.model.Add(e)
	if err != nil {
		return nil, err
	}
	return convertSample(e), nil
}

func (s *SampleService) Show(token *jwt.Token, id int) (*app.Sample, *sample_error.SampleError) {
	userId, err := s.User.AuthWithToken(token)
	if err != nil {
		return nil, err
	}

	e, err := s.model.Get(id, *userId)
	if err != nil {
		return nil, err
	}
	return convertSample(e), nil
}

func (s *SampleService) Update(token *jwt.Token, id int, name string, detail string) *sample_error.SampleError {
	userId, err := s.User.AuthWithToken(token)
	if err != nil {
		return err
	}

	e := &entities.Sample{}
	e.ID = id
	e.UserID = *userId
	e.Name = name
	e.Detail = detail
	err = s.model.Update(e)
	if err != nil {
		return err
	}
	return nil
}

func (s *SampleService) Delete(token *jwt.Token, id int) *sample_error.SampleError {
	userId, err := s.User.AuthWithToken(token)
	if err != nil {
		return err
	}

	err = s.model.Delete(id, *userId)
	if err != nil {
		return err
	}
	return nil
}
