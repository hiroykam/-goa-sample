//go:generate goagen bootstrap -d github.com/hiroykam/goa-sample/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/controller"
	"github.com/hiroykam/goa-sample/db"
)

func main() {
	// Create service
	service := goa.New("goa-sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	gdb, err := db.Open()
	if err != nil {
		service.LogError("open db", "err", err)
		return
	}

	// Mount "samples" controller
	sample := controller.NewSamplesController(service, gdb)
	app.MountSamplesController(service, sample)

	// Mount "swagger" controller
	s := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, s)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
