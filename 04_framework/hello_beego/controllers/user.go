package controllers

import (
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	userId := this.Ctx.Input.Param(":id")
	intUserId, _ := strconv.ParseInt(userId, 10, 64)
	this.Data["name"] = "jsser"
	this.Data["user_id"] = intUserId
	this.ServeJSON()
}
