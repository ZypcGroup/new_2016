package controller

import (
	"new_2016/models"
	//	"time"

	"fmt"

	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func Showhandler(ctx *macaron.Context, sess session.Store) {
	if fmt.Sprintf("%v", sess.Get("status")) == "true" {
		ctx.Data["IsLogin"] = true

	}
	exit := ctx.Req.FormValue("exit")

	if exit == "true" {
		sess.Destory(ctx)
		sess.GC()

	}

	ctx.Data["WebSiteTitle"] = websitetitle
	ctx.Data["WebSiteHome"] = websitehome
	ctx.Data["WebSiteLink"] = websitelink
	ctx.Data["WebSiteIcon"] = websiteicon
	ctx.Data["IsInfo"] = true

	ctx.Data["Info"] = models.ShowInfomation()
	ctx.Data["AllInfo"] = models.ShowAllInfomation()

	ctx.HTML(200, "show")
}
