package main

import (
	"fmt"

	//	"github.com/Unknwon/goconfig"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/csrf"
	//	"github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	//	"github.com/go-macaron/binding"
)

func init() {

}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())

	//		m.Use(gzip)
	//	m.Use(macaron.Static())

	m.Use(macaron.Renderer())

	//		m.Use(i18n)

	m.Use(cache.Cacher())
	m.Use(captcha.Captchaer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())

	//	m.Use(toolbox)

	m.Get("/", hello)
	m.Run(4000)
}

func hello(ctx *macaron.Context) string {
	fmt.Println(ctx.Req)
	return "hello zypc."
}
