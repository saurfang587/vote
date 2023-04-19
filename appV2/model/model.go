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
	Id          int64      `gorm:"id" json:"id"`
	Title       string     `json:"title"`
	During      int        `json:"during"`
	Type        int        `json:"type"`
	Count       int        `json:"count"`
	UserId      int        `json:"user_id"`
	StartTime   int64      `json:"start_time"`
	CreatedTime time.Time  `json:"created_time"`
	VoteOpt     []*VoteOpt `gorm:"foreignKey:VoteId" json:"vote_opt"`
}

type VoteOpt struct {
	Id          int64     `gorm:"id"`
	VoteId      int64     `json:"vote_id"`
	Count       int       `json:"count"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_time"`
	Percent     int64     `json:"percent"`
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
