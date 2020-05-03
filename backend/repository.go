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
	DB, err = gorm.Open("postgres", "host=db user=pg password=pg dbname=pg sslmode=disable")
	// DB, err = gorm.Open("postgres", "host=on-board.cgzfjoukl84a.ap-northeast-1.rds.amazonaws.com user=pg password=on-board dbname=pg sslmode=disable")
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

// FetchWorks is
func FetchWorks(DB *gorm.DB, userID uint) ([]*Work, error) {
	var works []*Work
	if userID > 0 {
		err = DB.Find(&works, "user_id=?", userID).Error
	} else {
		err = DB.Find(&works).Error
	}
	return works, err
}

// FetchWorkByID is
func FetchWorkByID(DB *gorm.DB, id uint) (*Work, error) {
	var work Work
	err = DB.First(&work, id).Error
	return &work, err
}

// FetchWorkItemByID is
func FetchWorkItemByID(DB *gorm.DB, id uint) (*WorkItem, error) {
	var item WorkItem
	err = DB.First(&item, id).Error
	return &item, err
}

// FetchWorkItemsByWorkID is
func FetchWorkItemsByWorkID(DB *gorm.DB, workid uint) ([]*WorkItem, error) {
	var items []*WorkItem
	err = DB.Find(&items, "work_id=?", workid).Error
	return items, err
}

// FetchUsers is
func FetchUsers(DB *gorm.DB) ([]*User, error) {
	var users []*User
	err = DB.Find(&users).Error
	return users, err
}

// FetchUserByID is
func FetchUserByID(DB *gorm.DB, id uint) (*User, error) {
	var user User
	err = DB.First(&user, id).Error
	return &user, err
}

// FetchUserByEmail is
func FetchUserByEmail(DB *gorm.DB, email string) (*User, error) {
	var user User
	err = DB.First(&user, "email=?", email).Error
	return &user, err
}

// FetchInfoByUID is
func FetchInfoByUID(DB *gorm.DB, uid uint) ([]*Info, error) {
	var aLotOfInfo []*Info
	err = DB.Find(&aLotOfInfo, "user_id=? OR user_id=?", uid, 0).Error
	return aLotOfInfo, err
}
