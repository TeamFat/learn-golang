package controllers

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	//方式1
	userId := this.Ctx.Input.Param(":id")
	_ = userId
	//方式2 - 直接获取的是int类型 , 如果获取其他类型： GetString(), GetStrings(), GetInt(), GetBool, GetFloat()
	userId2 := this.Input().Get("id")
	_ = userId2
	//方式3
	//支持数据绑定 - 将数据直接绑定到id上
	var id int
	this.Ctx.Input.Bind(&id, ":id")
	this.Data["json"] = map[string]interface{}{"user_id": id, "username": "jsser"}
	this.ServeJSON()
}
