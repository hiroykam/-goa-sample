package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa-sample", func() {
	Title("The Sample API")
	Description("A simple goa service")
	Version("v1")
	Scheme("http", "https")
	BasePath("/api/v1")
	Consumes("application/json")
	Produces("application/json")
	Host("localhost:8080")

	Origin("http://localhost:8080/swagger", func() {
		Expose("X-Time")
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})
})
