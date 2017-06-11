package backend

import (
	"strings"

	"strconv"

	"github.com/TeamFat/learn-golang/04_framework/hello_beego/models"
	"github.com/TeamFat/learn-golang/04_framework/hello_beego/utils"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	isBackendLogin bool
	userId         int
	username       string
	moduleName     string
	controllerName string
	actionName     string
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")

	return s[0]

}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "backend"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	//验证
	this.auth()
}

func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arr) == 2 {
		idstr, token := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			var user models.User
			user.Id = userId
			if user.Read() == nil && token == utils.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
				this.userId = user.Id
				this.username = user.Username
			}
		}
	}

	if this.userId == 0 {
		this.Redirect("/admin/login", 302)
	}

}
