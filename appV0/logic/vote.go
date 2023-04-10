package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saurfang/vote/appV0/model"
	"strconv"
)

func GetVote(c *gin.Context) {
	ret := model.GetVote(1)

	if ret.Id > 0 {
		fmt.Printf("vote:%+v\n", ret)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"vote": ret,
		})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"err_title": "数据查询失败",
	})
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
		re := fmt.Sprintf("/vote?id=%d", voteId)
		c.Redirect(http.StatusMovedPermanently, re)
	}
}

func DoVoteResult(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)

	if ret.Id > 0 {
		for _, opt := range ret.VoteOpt {
			s, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(opt.Count)/float64(ret.Count)), 64)
			v, _ := strconv.ParseInt(fmt.Sprintf("%.f", s*100), 10, 64)
			opt.Percent = v
		}
		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"vote": ret,
		})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"err_title": "数据查询失败",
	})
}
