package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	_ "saurfang/vote/appV2/docs"
	"saurfang/vote/appV2/logic"
	"saurfang/vote/appV2/tools"
)

func New() *gin.Engine {
	r := gin.Default()

	{
		r.GET("/votes", logic.GetVotes)
	}

	vote := r.Group("/vote")
	vote.Use(AuthCheck()) //后续添加 登录和限流中间件
	{
		//执行投票接口
		vote.POST("/do", logic.DoVote)

		//新增投票RestFul接口
		vote.GET("/:id", logic.GetVote)
		vote.POST("/basic", logic.AddVoteBasic)
		vote.PUT("/basic", logic.PutVoteBasic)
		vote.DELETE("/:id", logic.DeleteVoteBasic)

		//选项
		vote.GET("/opts/:vote_id", logic.GetOpts)
		vote.POST("/opt", logic.AddOpt)
		vote.PUT("/opt", logic.PutOpt)
		vote.DELETE("/opt/:id", logic.DeleteOpt)

	}

	login := r.Group("")
	{
		login.POST("/login", logic.Login)
		login.GET("/logout", logic.Logout)
	}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//测试模式不需要验签,需要自己在请求的头部伪造一个Debug数据
		if c.GetHeader("debug") != "" {
			c.Next()
			return
		}
		auth := c.GetHeader("Authorization")
		fmt.Printf("auth:%+v\n", auth)
		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "验签失败！",
			})
		}
		fmt.Printf("data:%+v\n", data)
		if data.ID <= 0 || data.Name == "" {
			//如果用户没有登录，中间件直接返回，不再向后继续
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "用户信息获取错误",
			})
			return
		}

		//将用户信息塞到Context中
		c.Set("name", data.Name)
		c.Set("userId", data.ID)

		c.Next()
	}
}
