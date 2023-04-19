package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saurfang/vote/appV2/model"
	"saurfang/vote/appV2/tools"
)

type User struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
		})
	}
	//TODO: 入参校验 和 SQL注入问题
	fmt.Printf("data:%+v\n", user)
	dbUser := model.GetUser(user.Name, user.Pwd)
	fmt.Printf("user:%+v\n", dbUser)
	if dbUser.Id > 0 {
		err := model.SetSession(c, dbUser.Name, dbUser.Id)
		if err != nil {
			c.JSON(http.StatusOK, tools.HttpCode{
				Code:    tools.UserInfoErr,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "登录成功，整在跳转~",
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "用户信息错误",
	})
	return
}

func Logout(c *gin.Context) {
	_ = model.FlushSession(c)
	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
	})
	return
}
