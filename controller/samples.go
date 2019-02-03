package controller

import (
	"github.com/goadesign/goa"
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/hiroykam/goa-sample/sample_logger"
	"github.com/hiroykam/goa-sample/services"
	"github.com/jinzhu/gorm"
)

// SamplesController implements the samples resource.
type SamplesController struct {
	*goa.Controller
	db *gorm.DB
}

// NewSamplesController creates a samples controller.
func NewSamplesController(service *goa.Service, db *gorm.DB) *SamplesController {
	return &SamplesController{
		service.NewController("SamplesController"),
		db,
	}
}

// Add runs the add action.
func (c *SamplesController) Add(ctx *app.AddSamplesContext) error {
	// SamplesController_Add: start_implement

	// Put your logic here
	l, err := sample_logger.New(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}
	l.Info("Sample::Add")

	s := services.NewSampleService(c.db)
	res, err := s.Add(ctx.Payload.UserID, ctx.Payload.Name, ctx.Payload.Detail)
	if err != nil {
		l.SampleError(err)
		return ctx.BadRequest(err)
	}

	return ctx.OK(res)
	// SamplesController_Add: end_implement
}

// Delete runs the delete action.
func (c *SamplesController) Delete(ctx *app.DeleteSamplesContext) error {
	// SamplesController_Delete: start_implement

	// Put your logic here
	l, err := sample_logger.New(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}
	l.Info("Sample::Delete")

	s := services.NewSampleService(c.db)
	err = s.Delete(ctx.ID)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}

	return ctx.NoContent()
	// SamplesController_Delete: end_implement
}

// List runs the list action.
func (c *SamplesController) List(ctx *app.ListSamplesContext) error {
	// SamplesController_List: start_implement

	// Put your logic here
	l, err := sample_logger.New(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}
	l.Info("Sample::List")

	s := services.NewSampleService(c.db)
	res, err := s.GetSamples(ctx.UserID)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}
	return ctx.OK(res)
	// SamplesController_List: end_implement
}

// Show runs the show action.
func (c *SamplesController) Show(ctx *app.ShowSamplesContext) error {
	// SamplesController_Show: start_implement

	// Put your logic here
	l, err := sample_logger.New(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}
	l.Info("Sample::Show")

	s := services.NewSampleService(c.db)
	res, err := s.Show(ctx.ID)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}

	return ctx.OK(res)
	// SamplesController_Show: end_implement
}

// Update runs the update action.
func (c *SamplesController) Update(ctx *app.UpdateSamplesContext) error {
	// SamplesController_Update: start_implement

	// Put your logic here
	l, err := sample_logger.New(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}
	l.Info("Sample::Update")

	s := services.NewSampleService(c.db)
	err = s.Update(ctx.ID, ctx.Payload.UserID, ctx.Payload.Name, ctx.Payload.Detail)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}

	return ctx.NoContent()
	// SamplesController_Update: end_implement
}
