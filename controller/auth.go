package controller

import (
	"github.com/goadesign/goa"
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/hiroykam/goa-sample/sample_logger"
	"github.com/hiroykam/goa-sample/services"
	"github.com/jinzhu/gorm"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db *gorm.DB
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, db *gorm.DB) *AuthController {
	return &AuthController{
		service.NewController("AuthController"),
		db,
	}
}

// Login runs the login action.
func (c *AuthController) Login(ctx *app.LoginAuthContext) error {
	// AuthController_Login: start_implement

	// Put your logic here

	l, err := sample_logger.NewSampleLooger(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}

	user, err := services.NewUserService(c.db)
	if err != nil {
		return ctx.BadRequest(err)
	}

	res, err := user.AuthWithEmailAndPassword(ctx.Payload.Email, ctx.Payload.Password)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}

	return ctx.OK(res)
	// AuthController_Login: end_implement
}

func (c *AuthController) Reauthenticate(ctx *app.ReauthenticateAuthContext) error {
	l, err := sample_logger.NewSampleLooger(ctx)
	if err != nil {
		return ctx.BadRequest(err)
	}

	if err != nil {
		return ctx.BadRequest(err)
	}

	h, err := services.NewHashedRefreshTokenService(c.db)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}

	a, err := h.Update(ctx.Payload.RefreshToken)
	if err != nil {
		l.SampleError(err)
		if err.Code == sample_error.NotFoundError {
			return ctx.NotFound()
		}
		return ctx.BadRequest(err)
	}

	return ctx.OK(a)
}
