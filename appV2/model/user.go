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

func GetUser(name, pwd string) *User {
	ret := &User{}
	if err := GlobalConn.Table("users").Where("name = ? and pwd = ?", name, pwd).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}
