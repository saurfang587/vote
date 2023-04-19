package router

import (
	"github.com/gin-gonic/gin"
	"saurfang/vote/appV2/logic"
	"saurfang/vote/appV2/model"
)

func New() *gin.Engine {
	r := gin.Default()

	{
		//首页方法，这两个同指向根目录
		r.GET("/", logic.Votes)
		r.GET("/index", logic.Votes)
		r.GET("/votes", logic.GetVotes)
	}

	vote := r.Group("")
	vote.Use(AuthCheck()) //后续添加 登录和限流中间件
	{
		vote.GET("/info", logic.GetVote)
		vote.POST("/vote", logic.DoVote)
	}

	login := r.Group("")
	{
		login.POST("/login", logic.Login)
		login.GET("/logout", logic.Logout)
	}

	return r
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := model.GetSession(c)
		id, ok1 := data["id"]
		name, ok2 := data["name"]
		idInt64, _ := id.(int64)
		if !ok1 || !ok2 || idInt64 <= 0 || name == "" {
			//c.Redirect(http.StatusFound, "/login")
			c.Abort() //如果用户没有登录，中间件直接返回，不再向后继续
		}
		c.Set("name", name)
		c.Set("userId", idInt64)
		c.Next()
	}
}
