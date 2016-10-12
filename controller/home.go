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
	//	"github.com/go-macaron/captcha"
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
}

func Homehandler(ctx *macaron.Context) {

	ctx.Data["WebSiteTitle"] = websitetitle
	ctx.Data["WebSiteHome"] = websitehome
	ctx.Data["WebSiteLink"] = websitelink
	ctx.Data["WebSiteIcon"] = websiteicon
	ctx.Data["IsHome"] = true
	ctx.HTML(200, "join")
}
