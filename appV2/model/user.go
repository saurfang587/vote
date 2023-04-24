package model

import "fmt"

// Check 校验用户信息
func Check(id, name string) bool {
	user := &User{}
	if GlobalConn.Table("users").Where("id = ?", id).First(user).Error != nil {
		return false
	}
	fmt.Printf("user:%+v\n", user)
	if user.Name != name {
		return false
	}

	return true
}

// CheckV1 增加一个原生SQL 的方法
func CheckV1(id, name string) bool {
	user := &User{}
	sql := "SELECT `id`,`name` from `users` where `id` = ? limit 1"
	err := GlobalConn.Raw(sql, id).Scan(user).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return false
	}

	if user.Name != name {
		return false
	}

	return true
}

func GetUser(name, pwd string) *User {
	ret := &User{}
	if err := GlobalConn.Table("users").Where("name = ? and pwd = ?", name, pwd).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

// GetUserV1 增加一个原生SQL 的方法
func GetUserV1(name, pwd string) *User {
	ret := &User{}
	sql := "select * from users where name = ? and pwd = ? limit 1"
	err := GlobalConn.Raw(sql, name, pwd).Scan(ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}
