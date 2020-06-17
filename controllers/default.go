package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["head"] = "edit_head.tpl"
	c.LayoutSections["body"] = "edit_body.tpl"
}
