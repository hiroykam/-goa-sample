package design

import (
	"time"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"

)

var MediaSamples = MediaType("application/vnd.samples+json", func() {
	Description("sample list")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "名前", func() {
		Example("サンプル1")
	})
	Attribute("created_at", DateTime, "作成日", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("updated_at", DateTime, "更新日", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(2019, 01, 31, 12, 30, 50, 0, loc).Format(time.RFC3339))
	})
	Required("id", "name", "created_at", "updated_at")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var MediaSample = MediaType("application/vnd.sample+json", func() {
	Description("sample detail")
	Attribute("id", Integer, "sample id", func() {
		Example(1)
	})
	Attribute("user_id", Integer, "user id", func() {
		Example(1)
	})
	Attribute("name", String, "名前", func() {
		Example("サンプル1")
	})
	Attribute("detail", String, "詳細", func() {
		Example("サンプル1の詳細")
	})
	Attribute("created_at", DateTime, "作成日", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("updated_at", DateTime, "更新日", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(2019, 01, 31, 12, 30, 50, 0, loc).Format(time.RFC3339))
	})
	Required("id", "user_id", "name", "detail", "created_at", "updated_at")
	View("default", func() {
		Attribute("id")
		Attribute("user_id")
		Attribute("name")
		Attribute("detail")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
