package models

import (
	"fmt"
	// "github.com/go-xorm/xorm"
	"time"
)

func AddInfomation(info *Infomation) bool {

	fmt.Println(info)
	connectDB()
	info.Time = time.Now()
	_, err := engine.Insert(info)
	if err == nil {
		return true
	}
	defer engine.Close()

	return false

}

func ModifyInfomation() {

}

func SearchInfomation() {

}

func ShowInfomation() (Info *Infomation) {
	connectDB()
	var userid int64
	userid = 4141168
	// Infoi := &Infomation{UserId: userid}
	Info = &Infomation{UserId: userid}
	engine.Id(userid).Get(Info)
	fmt.Println(Info)
	defer engine.Close()

	return
}

func ShowAllInfomation() (AllInfo []*Infomation) {
	connectDB()
	// Infoi := &Infomation{UserId: userid}

	AllInfo = make([]*Infomation, 0)
	engine.Find(&AllInfo)
	fmt.Println(AllInfo)
	defer engine.Close()

	return
}

func ExportInfomation() {

}
