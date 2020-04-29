package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is var
var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("postgres", "host=on-board.cgzfjoukl84a.ap-northeast-1.rds.amazonaws.com user=pg password=on-board dbname=pg sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{}, &Work{}, &WorkItem{})
	fmt.Println("-------------------------------------\n",
		"-------------------------------------\n",
		"-----------Migration--Done-----------\n",
		"-------------------------------------\n",
		"-------------------------------------")
}

func FetchWorks(DB *gorm.DB, userID uint) ([]*Work, error) {
	var works []*Work
	if userID > 0 {
		err = DB.Find(&works, "user_id=?", userID).Error
	} else {
		err = DB.Find(&works).Error
	}
	return works, err
}

func FetchWorkByID(DB *gorm.DB, id uint) (*Work, error) {
	var work Work
	err = DB.First(&work, id).Error
	return &work, err
}

func FetchWorkItemByID(DB *gorm.DB, id uint) (*WorkItem, error) {
	var item WorkItem
	err = DB.First(&item, id).Error
	return &item, err
}

func FetchWorkItemsByWorkID(DB *gorm.DB, workid uint) ([]*WorkItem, error) {
	var items []*WorkItem
	err = DB.Find(&items, "work_id=?", workid).Error
	return items, err
}

func FetchUsers(DB *gorm.DB) ([]*User, error) {
	var users []*User
	err = DB.Find(&users).Error
	return users, err
}

func FetchUserByID(DB *gorm.DB, id uint) (*User, error) {
	var user User
	err = DB.First(&user, id).Error
	return &user, err
}

func FetchUserByEmail(DB *gorm.DB, email string) (*User, error) {
	var user User
	err = DB.First(&user, "email=?", email).Error
	return &user, err
}
