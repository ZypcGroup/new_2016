package controller

import (
	"gopkg.in/macaron.v1"
	"time"
	"zypc_submit/models"
)

func Showhandler(ctx *macaron.Context) {
	sess, _ := Sess.Start(ctx)

	createtime := sess.Get("CreateTime")
	// fmt.Println(createtime, "\n-------------\n")
	if createtime == nil {
		// fmt.Println(createtime, "\n-------------\n")
		ctx.Redirect("/login", 301)
	} else if (createtime.(int64) + 360) > time.Now().Unix() {
		// fmt.Println(createtime, "\n-------------\n")

		ctx.Data["IsLogin"] = true
		// ctx.Redirect("/show", 301)
	}

	exit := ctx.Req.FormValue("exit")

	if exit == "true" {
		sess.Delete(createtime)
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
