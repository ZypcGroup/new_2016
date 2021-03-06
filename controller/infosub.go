package controller

import (
	"fmt"
	"new_2016/models"
	"strconv"
	"time"

	"github.com/go-macaron/captcha"
	"gopkg.in/macaron.v1"
)

type User struct {
	UserId     int64 `xorm:"index"`
	UserName   string
	Password   string
	Grade      string
	Choose     string
	Department string
	Email      string
	Telnumber  string
	Time       time.Time `xorm:"index"`
	Flag       int
}

var user *models.User

//var user *User

func InfoSubHandler(ctx *macaron.Context, cpt *captcha.Captcha)  {
	if !cpt.VerifyReq(ctx.Req) {
		ErrorInfo = "验证码错误！"
		ctx.Redirect("/errorinfo", 301)
	}
	user = new(models.User)
	//	user = new(User)
	user.UserName = ctx.Req.FormValue("name")
	user.Choose = ctx.Req.FormValue("choose")
	user.Department = ctx.Req.FormValue("department")
	user.Email = ctx.Req.FormValue("email")
	user.Flag = 0
	user.Grade = ctx.Req.FormValue("grade")
	user.Telnumber = ctx.Req.FormValue("phone")
	user.Time = time.Now()
	userid := ctx.Req.FormValue("number")
	t, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	user.UserId = t

	ok, err := models.CheckUser(user.UserId)
	if err != nil {
		fmt.Println(err)
	}
	if !ok {
		err = models.AddUser(user)
		if err != nil {
			fmt.Println(err)
		}
		ctx.HTML(200,"success")
	}

	ctx.HTML(200,"return")

}
