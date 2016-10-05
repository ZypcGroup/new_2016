package controller

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	// "github.com/go-macaron/session"
	// "macaron/controller"
	// "zypc_submit/models"
	"time"
)

var conf *goconfig.ConfigFile

const (
	WebSiteTitle = "智邮普创工作室信息提交系统"
	WebSiteHome  = "智邮普创工作室"
	WebSiteLink  = "http://www.smartxupt.com"
	WebSiteIcon  = "images/zypc.ico"
)

var (
	websitetitle = WebSiteTitle
	websitehome  = WebSiteHome
	websitelink  = WebSiteLink
	websiteicon  = WebSiteIcon
)

func init() {
	conf, err := goconfig.LoadConfigFile("conf/app.conf")
	if err != nil {
		fmt.Println("Load Config File Error! \t", err)
	}

	if ok, err := conf.GetValue("AppInfo", "WebSiteTitle"); err == nil {
		websitetitle = ok
	}
	if ok, err := conf.GetValue("AppInfo", "WebSiteHome"); err == nil {
		websitehome = ok
	}
	if ok, err := conf.GetValue("AppInfo", "WebSiteLink"); err == nil {
		websitelink = ok
	}
	if ok, err := conf.GetValue("AppInfo", "WebSiteIcon"); err == nil {
		websiteicon = ok
	}
	// Debug info
	if ok, err := conf.GetValue("Server", "Debug"); err == nil {
		if ok == "true" {
			Debug = true
		}
	}
}

func Homehandler(ctx *macaron.Context) {

	sess, _ := Sess.Start(ctx)

	createtime := sess.Get("CreateTime")
	// fmt.Println(createtime, "\n-------------\n")
	if createtime == nil {
		// fmt.Println(createtime, "\n-------------\n")
		ctx.Redirect("/login", 301)
	} else if (createtime.(int64) + 360) > time.Now().Unix() {
		// fmt.Println(createtime, "\n-------------\n")

		ctx.Data["IsLogin"] = true
	} else {
		ctx.Redirect("/login", 301)
	}

	exit := ctx.Req.FormValue("exit")

	if exit == "true" {
		sess.Delete(createtime)
	}

	ctx.Data["WebSiteTitle"] = websitetitle
	ctx.Data["WebSiteHome"] = websitehome
	ctx.Data["WebSiteLink"] = websitelink
	ctx.Data["WebSiteIcon"] = websiteicon
	ctx.Data["IsHome"] = true

	ctx.Data["Info1"] = "学号"
	ctx.Data["Info2"] = "班级"
	ctx.Data["Info3"] = "姓名"
	ctx.Data["Info4"] = "博客地址"
	ctx.Data["Info5"] = "备注"

	ctx.HTML(200, "home")
}
