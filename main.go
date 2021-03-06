package main

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	"new_2016/controller"
	"new_2016/models"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/session"
	// "macaron/modules/initConf"
	// "encoding/base64"
)

const (
	Port       int = 8080
	cookieName     = "ZypcCookie"
)

var (
	port int = Port
)

var conf *goconfig.ConfigFile

func init() {

	err := models.RegisterDB()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	conf, err = goconfig.LoadConfigFile("conf/app.conf")
	if err != nil {
		fmt.Println("Load Config File Error! \t", err)
	}

	if ok := conf.MustInt("Server", "ListenPort"); ok != 0 {
		port = ok
	}

}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner(session.Options{
		Provider: "memory",
		// 提供器的配置，根据提供器而不同
		// ProviderConfig: providerConfig,
		// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
		CookieName: cookieName,
	}))
	m.Use(cache.Cacher())
	m.Use(captcha.Captchaer())
	m.Get("/", controller.Homehandler)
	m.Get("/errorinfo", controller.ErrorInfohandler)

	m.Post("/infosub", controller.InfoSubHandler)
	m.Run(port)
}
