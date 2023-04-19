package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saurfang/vote/appV1/model"
	"saurfang/vote/appV1/tools"
	"strconv"
)

func Votes(c *gin.Context) {
	n := "游客"
	name, ok := model.GetSession(c)["name"]
	if ok && name != "" {
		n = name.(string)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"name": n,
	})
	return
}

func GetVotes(c *gin.Context) {
	ret := model.GetVotes()
	c.JSON(http.StatusOK, tools.HttpCode{ //注意此处返回的为HTTP的错误代码
		Code: tools.OK, //此处返回的是业务自定义错误码
		Data: ret,
	})
}

func GetVote(c *gin.Context) {
	n := "游客"
	name, ok := model.GetSession(c)["name"]
	if ok && name != "" {
		n = name.(string)
	}

	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if id <= 0 {
		c.HTML(http.StatusOK, "info.tmpl", gin.H{
			"err_title": "数据错误",
			"name":      n,
		})
		return
	}

	ret := model.GetVote(id)
	if ret.Id > 0 {
		fmt.Printf("vote:%+v\n", ret)
		c.HTML(http.StatusOK, "info.tmpl", gin.H{
			"err":  "",
			"name": n,
			"vote": ret,
		})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"err_title": "数据查询失败",
		"name":      n,
	})
	return
}

func DoVote(c *gin.Context) {
	voteIdStr, _ := c.GetPostForm("vote_id")
	opt, _ := c.GetPostFormArray("opt[]")
	id, _ := c.Cookie("id")
	fmt.Printf("vote:%s", voteIdStr)
	fmt.Printf("opt:%v", opt)
	optIds := make([]int64, 0)
	for _, s := range opt {
		tmp, _ := strconv.ParseInt(s, 10, 64)
		optIds = append(optIds, tmp)
	}
	voteId, _ := strconv.ParseInt(voteIdStr, 10, 64)
	userId, _ := strconv.ParseInt(id, 10, 64)

	if model.DoVote(userId, voteId, optIds) {

		c.JSON(http.StatusOK, tools.HttpCode{
			Code: tools.OK,
			Data: VoteResult(voteId),
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.DoErr,
		Message: "投票失败！",
	})
	return
}

func VoteResult(id int64) *model.Vote {
	ret := model.GetVote(id)

	if ret.Id > 0 {
		for _, opt := range ret.VoteOpt {
			s, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(opt.Count)/float64(ret.Count)), 64)
			v, _ := strconv.ParseInt(fmt.Sprintf("%.f", s*100), 10, 64)
			opt.Percent = v
		}

	}

	return ret
}
