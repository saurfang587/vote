package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saurfang/vote/appV0/model"
	"saurfang/vote/appV0/tools"
	"strconv"
)

type User struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
		})
		return
	}
	//TODO: 入参校验
	//SQL 注入
	fmt.Printf("data:%+v\n", user)
	dbUser := model.GetUser(user.Name, user.Pwd)
	fmt.Printf("user:%+v", dbUser)
	if dbUser.Id > 0 {
		//设置cookie
		c.SetCookie("name", dbUser.Name, 3600, "/", "", false, true) //domain 写域名的话 会导致IP访问无效
		c.SetCookie("id", strconv.FormatInt(dbUser.Id, 10), 3600, "/", "", false, true)
		c.Redirect(http.StatusMovedPermanently, "/") //新问题
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "用户信息错误",
	})
	return
}

func Logout(c *gin.Context) {
	//设置cookie
	c.SetCookie("name", "", 3600, "/", "", false, true) //domain 写域名的话 会导致IP访问无效
	c.SetCookie("id", "", 3600, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/login")
	return

}
