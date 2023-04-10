package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saurfang/vote/appV0/logic"
	"saurfang/vote/appV0/model"
)

func New() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("appV0/view/*")
	r.Static("/kit", "./kit")
	basic := r.Group("")
	basic.Use(AuthCheck()) //后续添加 登录和限流中间件
	//basic.GET("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "这是index信息",
	//	})
	//})

	basic.GET("/", logic.GetVote)
	basic.POST("/vote", logic.DoVote)
	basic.GET("/vote", logic.DoVoteResult)
	login := r.Group("")
	{
		login.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"title": "投票系统",
			})
		})
		login.POST("/login", logic.Login)
		login.GET("/logout", logic.Logout)
	}

	return r
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, _ := c.Cookie("name")
		id, _ := c.Cookie("id")
		fmt.Printf("name:%s tel:%s\n ", name, id)
		if name == "" || id == "" || !model.Check(id, name) {
			//c.JSON(http.StatusOK, tools.HttpCode{ //注意此处返回的为HTTP的错误代码
			//	Code:    tools.NotLogin, //此处返回的是业务自定义错误码
			//	Message: "您尚未登录",
			//})
			c.Redirect(http.StatusMovedPermanently, "/login")

			c.Abort() //如果用户没有登录，中间件直接返回，不再向后继续
		}
		c.Next()
	}
}
