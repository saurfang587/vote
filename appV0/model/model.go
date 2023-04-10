package model

import "time"

type User struct {
	Id          int64 `gorm:"id"`
	Name        string
	Pwd         string
	Tel         string
	CreatedTime time.Time
}

type Vote struct {
	Id          int64 `gorm:"id"`
	Title       string
	During      int
	Type        int
	Count       int
	UserId      int
	StartTime   int64
	CreatedTime time.Time
	VoteOpt     []*VoteOpt `gorm:"foreignKey:VoteId"`
}

type VoteOpt struct {
	Id          int64 `gorm:"id"`
	VoteId      int64
	Count       int
	Name        string
	CreatedTime time.Time
	Percent     int64
}

type VoteOptUser struct {
	Id        int64 `gorm:"id"`
	VoteId    int64
	UserId    int64
	VoteOptId int64
}

func (Vote) TableName() string {
	return "vote_title"
}

func (VoteOpt) TableName() string {
	return "vote_opt"
}

func (VoteOptUser) TableName() string {
	return "vote_opt_user"
}
