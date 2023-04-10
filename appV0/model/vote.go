package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetVote(id int64) *Vote {
	ret := &Vote{}
	if err := GlobalConn.Preload("VoteOpt").Table("vote_title").Where("id = ?", id).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

func GetVotes() {

}

func DoVote(userid, voteId int64, opt []int64) bool {
	var err error
	tx := GlobalConn.Begin()
	//TODO：查询是否投过票

	//投票总数增加
	err = tx.Table("vote_title").Where("id = ?", voteId).
		Update("count", gorm.Expr("count + ?", len(opt))).
		Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		tx.Rollback()
		return false
	}

	for _, va := range opt {
		//投票选项+1
		err = tx.Table("vote_opt").Where("id = ?", va).
			Update("count", gorm.Expr("count + ?", 1)).
			Error
		//记录用户投票记录
		user := VoteOptUser{

			VoteId:    voteId,
			UserId:    userid,
			VoteOptId: va,
		}
		err = tx.Create(&user).Error // 通过数据的指针来创建
	}

	if err != nil {
		fmt.Printf("err:%s", err.Error())
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}

func GetVoteByUser() {

}
