package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

var _ = Resource("auth", func() {
	Action("login", func() {
		Routing(POST("/login"))
		Payload(func() {
			Attribute("email", String, "name of sample", func() {
				Example("sample@goa-sample.test.com")
			})
			Attribute("password", String, "detail of sample", func() {
				Example("test1234")
			})
			Required("email", "password")
		})
		Response(OK, AuthSamples)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("samples", func() {
	BasePath("/samples")
	Description("sample APIs with JWT Authorization")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("list", func() {
		Description("複数")
		Routing(
			GET("/"),
		)
		Response(OK, CollectionOf(MediaSamples))
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})
	Action("show", func() {
		Description("単数")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "sample id", func() {
				Example(678)
			})
			Required("id")
		})
		Response(OK, MediaSample)
		Response(NotFound)
		Response(Unauthorized, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("追加")
		Routing(
			POST("/"),
		)
		Payload(func() {
			Attribute("name", String, "name of sample", func() {
				Example("sample1のタイトル")
			})
			Attribute("detail", String, "detail of sample", func() {
				Example("sample1の詳細")
			})
			Required("name", "detail")
		})
		Response(OK, MediaSample)
		Response(NotFound)
		Response(Unauthorized, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("削除")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "sample id", func() {
				Example(678)
			})
			Required("id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("更新")
		Routing(
			PATCH("/:id"),
		)
		Params(func() {
			Param("id", Integer, "sample id")
			Required("id")
		})
		Payload(func() {
			Param("name", String, "name of sample", func() {
				Example("sample1のタイトル")
			})
			Param("detail", String, "detail of sample", func() {
				Example("sample1の詳細")
			})
			Required("name", "detail")
		})
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})
