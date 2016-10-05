package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	"github.com/go-macaron/session"
	"zypc_submit/controller"
	"zypc_submit/models"
	// "macaron/modules/initConf"
	// "encoding/base64"
)

const (
	Port int = 8080
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
		Provider:       "file",
		ProviderConfig: "data/sessions",
	}))
	m.Get("/", controller.Homehandler)
	m.Post("/submit", controller.Submithandler)
	m.Get("/login", controller.Loginhandler)
	m.Post("/login", controller.LoginJudgehandler)
	m.Post("/register", controller.Registerhandler)
	m.Get("/errorinfo", controller.ErrorInfohandler)
	// m.Get("/test", controller.Testhandler)

	m.Get("/ok", controller.Okhandler)
	m.Get("/show", controller.Showhandler)
	m.Run(port)
}
