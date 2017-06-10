package controllers

import (
	"github.com/TeamFat/learn-golang/04_framework/hello_beego/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	isLogin bool
	user    models.User
}
