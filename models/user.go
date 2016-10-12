package models

import (
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	//	 "github.com/go-xorm/xorm"
	// "time"
	// "zypc_submit/modules"
	"strconv"
)

func AddUser(user *User) (err error) {
	connectDB()
	//	checkUserInSchool(user.UserId)

	_, err = engine.Insert(user)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func CheckUser(userid int64) (has bool, err error) {
	connectDB()

	schooluser := &YdStudent{Sno: strconv.FormatInt(userid, 10)}
	has, err = engine.Get(schooluser)
	if err != nil {
		return true, err
	}
	fmt.Println(has, err)
	if has {
		user := &User{UserId: userid}
		has, err = engine.Get(user)
		if err != nil {
			return false, err
		}
		return has, nil
	}
	fmt.Println(has, err)
	defer engine.Close()
	return true, nil
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
