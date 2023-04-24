package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetVote(id int64) *Vote {
	ret := &Vote{}
	if err := GlobalConn.Preload("VoteOpt").Table("vote_title").Where("id = ?", id).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

func GetVotes() []*Vote {
	ret := make([]*Vote, 0)
	if err := GlobalConn.Table("vote_title").Where("id > 0").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

func GetVotesV1() []*Vote {
	ret := make([]*Vote, 0)
	sql := "select * from `vote_title` where id > 0"
	err := GlobalConn.Raw(sql).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return ret
	}

	//为切片字段设置空值
	for _, vote := range ret {
		vote.VoteOpt = make([]*VoteOpt, 0)
	}
	return ret
}

func DoVote(userId, voteId int64, opt []int64) bool {
	var err error
	tx := GlobalConn.Begin()
	var count int64
	err = tx.Table("vote_opt_user").Where("user_id = ? and vote_id = ?", userId, voteId).Count(&count).Error
	if err != nil || count > 0 {
		fmt.Printf("err:%s", err.Error())
		tx.Rollback()
		return false
	}

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
			UserId:    userId,
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

func DoVoteV1(userId, voteId int64, opt []int64) bool {
	var err error
	tx := GlobalConn.Exec("begin")
	var count int64
	sql := "select count(id) from vote_opt_user where user_id = ? and vote_id = ?"
	err = tx.Raw(sql, userId, voteId).Scan(&count).Error
	if err != nil || count > 0 {
		fmt.Printf("err:%s", err.Error())
		tx.Exec("rollback")
		return false
	}

	//投票总数增加
	sql = "update vote_title set count=count+? where id = ? limit 1"
	err = tx.Exec(sql, len(opt), voteId).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		tx.Exec("rollback")
		return false
	}

	for _, va := range opt {
		//投票选项+1
		sql = "update vote_opt set count = count +1 where id = ? limit 1"
		if err = tx.Exec(sql, va).Error; err != nil {
			break
		}

		//记录用户投票记录
		sql = "insert into vote_opt_user (`vote_id`,`user_id`,`vote_opt_id`) values(?,?,?)"
		err = tx.Exec(sql, voteId, userId, va).Error
	}

	if err != nil {
		fmt.Printf("err:%s", err.Error())
		tx.Exec("rollback")
		return false
	}

	tx.Exec("commit")
	return true
}

// DeleteVote 伪删除
func DeleteVote(voteId int64) error {
	err := GlobalConn.Table("vote_title").Where("id = ?", voteId).Update("deleted", 1).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return err
	}
	return nil
}

// DeleteVoteV1 真删除
func DeleteVoteV1(voteId int64) error {
	err := GlobalConn.Table("vote_title").Delete(&Vote{}, voteId).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return err
	}
	return nil
}

// CreateVote 创建一个投票信息
func CreateVote(vote *Vote) error {
	return GlobalConn.Create(vote).Error
}

// UpdateVote 更新数据
func UpdateVote(vote *Vote) error {
	oldVote := Vote{}
	err := GlobalConn.Table("vote_title").Where("id = ?", vote.Id).First(&oldVote).Error
	if err != nil || (oldVote.UserId != 0 && oldVote.UserId != vote.UserId) {
		return errors.New("数据查询失败，或者用户ID错误。")
	}

	return GlobalConn.Where("id = ?", vote.Id).Updates(vote).Error
}
