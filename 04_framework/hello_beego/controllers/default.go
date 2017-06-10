package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	// c.Ctx.WriteString("hello")
	c.Data["title"] = "这是第一个beego文件标题"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
