// Package models provide all sql models for this project
// use gorm
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Role  bool
	Phone string
	avtar string
	// orgs ManyToMany
	// SelectedOrg oneToMany
}

type UserMixin struct {
	// 用户相关通用字段
}

type Org struct {
	gorm.Model
	Name  string
	OldID int
}
