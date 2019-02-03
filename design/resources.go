package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("samples", func() {
	BasePath("/samples")
	Action("list", func() {
		Description("複数")
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("user_id", Integer, "user id", func() {
				Example(12345)
			})
			Required("user_id")
		})
		Response(OK, CollectionOf(MediaSamples))
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
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
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("追加")
		Routing(
			POST("/"),
		)
		Payload(func() {
			Attribute("user_id", Integer, "user id", func() {
				Example(12345)
			})
			Attribute("name", String, "name of sample", func() {
				Example("sample1のタイトル")
			})
			Attribute("detail", String, "detail of sample", func() {
				Example("sample1の詳細")
			})
			Required("user_id", "name", "detail")
		})
		MultipartForm()
		Response(OK, MediaSample)
		Response(NotFound)
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
			Param("user_id", Integer, "name of sample", func() {
				Example(12345)
			})
			Param("name", String, "name of sample", func() {
				Example("sample1のタイトル")
			})
			Param("detail", String, "detail of sample", func() {
				Example("sample1の詳細")
			})
			Required("user_id", "name", "detail")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
