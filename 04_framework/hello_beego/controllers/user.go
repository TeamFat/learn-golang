package controllers

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	userId := this.Ctx.Input.Param(":id")
	_ = userId
	//支持数据绑定 - 将数据直接绑定到id上
	var id int
	this.Ctx.Input.Bind(&id, ":id")
	this.Data["json"] = map[string]interface{}{"user_id": id, "username": "jsser"}
	this.ServeJSON()
}
