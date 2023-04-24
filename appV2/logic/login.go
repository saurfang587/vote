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

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
//	@Summary		用户登录
//	@Description	会执行用户登录操作
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/login [POST]
func Login(c *gin.Context) {
	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	//TODO: 入参校验 和 SQL注入问题

	dbUser := model.GetUser(user.Name, user.Pwd)
	if dbUser.Id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	//下发Token
	a, r, err := tools.Token.GetToken(dbUser.Id, dbUser.Name)
	fmt.Printf("atoken:%s\n", a)
	fmt.Printf("rtoken:%s\n", r)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "Token生成失败！错误信息：" + err.Error(),
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "登录成功，整在跳转~",
		Data: Token{
			AccessToken:  a,
			RefreshToken: r,
		},
	})
	return
}

// Logout godoc
//
//	@Summary		用户退出
//	@Description	会执行用户退出操作
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@response		500,401	{object}	tools.HttpCode
//	@Router			/logout [get]
func Logout(c *gin.Context) {
	_ = model.FlushSession(c)
	//这里暂时先不改为 401，有些接口确实不需要登录态
	c.JSON(http.StatusUnauthorized, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}
