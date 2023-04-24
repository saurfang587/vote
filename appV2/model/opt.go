package model

import (
	"fmt"
)

func GetOpts(voteId int64) []*VoteOpt {
	ret := make([]*VoteOpt, 0)
	err := GlobalConn.Table("vote_opt").Where("vote_id = ?", voteId).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

// DeleteVote 伪删除
func DeleteOpt(id int64) error {
	err := GlobalConn.Table("vote_opt").Where("id = ?", id).Update("deleted", 1).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return err
	}
	return nil
}

// CreateOpt 创建一个投票选项
func CreateOpt(opt *VoteOpt) error {
	return GlobalConn.Create(opt).Error
}

// UpdateVote 更新数据
func UpdateOpt(opt *VoteOpt) error {
	return GlobalConn.Where("id = ?", opt.Id).Updates(opt).Error
}
