package main

import (
	"github.com/jinzhu/gorm"
)

// Work is Model
type Work struct {
	gorm.Model
	Name        string
	ImageURL    string
	Description string
	URL         string
	UserID      uint
}

// WorkItem is Model
type WorkItem struct {
	gorm.Model
	ImageURL string
	Body     string
	WorkID   uint
}

// User is Model
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"type:varchar(100);unique_index;not null"`
	ImageURL     string
	Introduction string
	Password     string
	URL          string
}
