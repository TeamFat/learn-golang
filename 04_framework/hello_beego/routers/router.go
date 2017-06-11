package routers

import (
	"encoding/json"

	"time"

	"github.com/TeamFat/learn-golang/04_framework/hello_beego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//使用controller -router
	beego.Router("/user/:id", &controllers.UserController{})

	//使用匿名函数 - router
	beego.Get("/payment/:user_id", func(ctx *context.Context) {
		userID := ctx.Input.Param(":user_id")
		ctx.Output.Header("Content-Type", "application/json")
		data := make(map[string]interface{})
		data["code"] = 200
		data["timestamp"] = time.Now().Unix()
		_data := map[string]interface{}{"user_id": userID, "username": "jsser", "sn": "20170301023789172398123", "pay_type": "qrcode", "pay_method": "wechat"}
		data["data"] = _data
		json, _ := json.Marshal(data)
		ctx.Output.Body([]byte(json))
	})

	//admin - router

}
