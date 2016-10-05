package controller

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	// "github.com/go-macaron/session"
	// "macaron/controller"
	// "strconv"
	"zypc_submit/models"
)

func Submithandler(ctx *macaron.Context) {

	sess, _ := Sess.Start(ctx)
	userid := sess.Get("UserID")
	// fmt.Println(userid)
	info := new(models.Infomation)
	// info.UserId, _ = strconv.ParseInt(string(userid), 10, 64)
	info.UserId = userid.(int64)
	info.Info1 = ctx.Req.FormValue("Info1")
	info.Info2 = ctx.Req.FormValue("Info2")
	info.Info3 = ctx.Req.FormValue("Info3")
	info.Info4 = ctx.Req.FormValue("Info4")
	info.Info5 = ctx.Req.FormValue("Info5")
	info.Info6 = ctx.Req.FormValue("info6")
	info.Info7 = ctx.Req.FormValue("info7")
	info.Info8 = ctx.Req.FormValue("info8")
	info.Info9 = ctx.Req.FormValue("info9")
	info.Info10 = ctx.Req.FormValue("info10")
	info.Info11 = ctx.Req.FormValue("info11")
	info.Info12 = ctx.Req.FormValue("info12")
	info.Info13 = ctx.Req.FormValue("info13")
	info.Info14 = ctx.Req.FormValue("info14")
	info.Info15 = ctx.Req.FormValue("info15")
	info.Info16 = ctx.Req.FormValue("info16")
	info.Info17 = ctx.Req.FormValue("info17")
	info.Info18 = ctx.Req.FormValue("info18")
	info.Info19 = ctx.Req.FormValue("info19")
	info.Info20 = ctx.Req.FormValue("info20")

	rel := models.AddInfomation(info)
	if rel {
		ctx.Redirect("/ok", 301)
	} else {
		ErrorInfo = "信息提交失败！"
		ctx.Redirect("/errorinfo", 301)
	}

}
