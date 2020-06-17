package routers

import (
	"github.com/astaxie/beego"
)

func init() {
    beego.SetStaticPath("/api", "static/api")
    beego.SetStaticPath("/css", "static/css")
    beego.SetStaticPath("/images", "static/images")
    beego.SetStaticPath("/js", "static/js")
    beego.SetStaticPath("/lib", "static/lib")
    beego.SetStaticPath("/page", "static/page")
    beego.SetStaticPath("/", "static/index.html")
}
