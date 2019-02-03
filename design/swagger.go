package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// Swagger routing
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger.yaml", "swagger/swagger.yaml")
	Files("/swagger-ui/*filepath", "public/swagger-ui/dist")
})
