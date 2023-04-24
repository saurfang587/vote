package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saurfang/vote/appV2/model"
	"saurfang/vote/appV2/tools"
	"strconv"
	"time"
)

// GetOpts godoc
//
//	@Summary		获取投票列表
//	@Description	获取投票列表
//	@Tags			opt
//	@Accept			json
//	@Produce		json
//	@Param			vote_id	path		int	false	"int valid"	minimum(1)
//	@response		200,500	{object}	tools.HttpCode{data=[]model.VoteOpt}
//	@Router			/opts/{vote_id} [get]
func GetOpts(c *gin.Context) {
	idStr := c.Param("vote_id")
	fmt.Printf("idstr:%s\n", idStr)
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code: tools.NotFound,
			Data: struct{}{},
		})
		return
	}

	ret := model.GetOpts(id)
	c.JSON(http.StatusOK, tools.HttpCode{ //注意此处返回的为HTTP的错误代码
		Code: tools.OK, //此处返回的是业务自定义错误码
		Data: ret,
	})
}

// AddOpt godoc
//
//	@Summary		新增投票选项
//	@Description	新增加投票选项
//	@Tags			opt
//	@Accept			json
//	@Produce		json
//	@Param			vote_id		body		int		false	"投票主ID"
//	@Param			name		body		string	false	"选项名称"
//	@response		200,400,500	{object}	tools.HttpCode
//	@Router			/opt [post]
func AddOpt(c *gin.Context) {
	opt := &model.VoteOpt{}
	if err := c.ShouldBind(&opt); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "数据解析失败",
			Data:    struct{}{},
		})
		return
	}

	fmt.Printf("opt:%+v", opt)
	opt.CreatedTime = time.Now()
	err := model.CreateOpt(opt)
	if err != nil {
		//这里没有返回4XX或者5XX的错误码
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "创建失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}

// PutOpt godoc
//
//	@Summary		更新投票选项
//	@Description	更新投票的选项
//	@Tags			opt
//	@Accept			json
//	@Produce		json
//	@Param			id			body		int		false	"投票主题ID"
//	@Param			name		body		string	false	"投票主题"
//	@response		200,400,500	{object}	tools.HttpCode
//	@Router			/opt [put]
func PutOpt(c *gin.Context) {
	opt := &model.VoteOpt{}
	if err := c.ShouldBind(&opt); err != nil || opt.Id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "数据解析失败,Id不能为空",
			Data:    struct{}{},
		})
		return
	}

	fmt.Printf("opt:%+v", opt)

	err := model.UpdateOpt(opt)
	if err != nil {
		//这里没有返回4XX或者5XX的错误码
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "更新失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}

// DeleteOpt godoc
//
//	@Summary		删除投票选项
//	@Description	根据ID删除选项
//	@Tags			opt
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int	false	"int valid"	minimum(1)
//	@response		200,404,500	{object}	tools.HttpCode
//	@Router			/opt/:id [delete]
func DeleteOpt(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	fmt.Printf("id : %s\n", idStr)
	if id <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code: tools.NotFound,
			Data: struct{}{},
		})
		return
	}

	if err := model.DeleteOpt(id); err == nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.DoErr,
		Message: "删除失败！",
		Data:    struct{}{},
	})
	return
}
