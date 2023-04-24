package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalConn *gorm.DB

func New() {
	//parseTime=True&loc=Local MySQL 默认时间是格林尼治时间，与我们差八小时，需要定位到我们当地时间
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "127.0.0.1:3306", "vote")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	GlobalConn = conn
}

func Close() {
	db, _ := GlobalConn.DB()
	_ = db.Close()
}

//建表语句
//CREATE TABLE `users` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
//`pwd` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
//`tel` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//PRIMARY KEY (`id`),
//KEY `tel` (`tel`) USING BTREE
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

//CREATE TABLE `vote_opt` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`vote_id` bigint DEFAULT NULL,
//`count` int DEFAULT NULL,
//`name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`percent` int DEFAULT NULL,
//`deleted` tinyint DEFAULT '0' COMMENT '删除标志位',
//PRIMARY KEY (`id`),
//KEY `vote_id` (`vote_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

//CREATE TABLE `vote_opt_user` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`vote_id` bigint DEFAULT NULL,
//`vote_opt_id` bigint DEFAULT NULL,
//`user_id` bigint DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

//CREATE TABLE `vote_title` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`title` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
//`start_time` bigint DEFAULT NULL,
//`type` int DEFAULT NULL COMMENT '0 单选 1 多选',
//`user_id` bigint DEFAULT NULL,
//`count` int DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`during` int DEFAULT NULL,
//`deleted` tinyint DEFAULT '0' COMMENT '1为删除',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
