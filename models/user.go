package models

import (
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/go-xorm/xorm"
	// "time"
	// "zypc_submit/modules"
)

func AddUser(user *User) (err error) {
	connectDB()
	_, err = engine.Insert(user)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func CheckUser(userid int64) (has bool, err error) {
	connectDB()
	user := &User{UserId: userid}
	has, err = engine.Get(user)
	if err != nil {
		return false, err
	}
	defer engine.Close()

	return has, nil
}

func ModifyUser() {

}

func JudgeUser(user *User) (has bool, err error) {
	connectDB()
	has, err = engine.Get(user)
	// mt.Println(user, "judge user")
	if err != nil {
		return false, err
	}
	fmt.Println(user, "judge user")
	defer engine.Close()
	return has, nil
}
