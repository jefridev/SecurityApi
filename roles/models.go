package roles

import (
	"time"

	"github.com/jefridev/securityapi/common"
	"github.com/jefridev/securityapi/users"
)

// RoleModel basically describes a specific role for using this system per se.
type RoleModel struct {
	ID          uint64    `gorm:"auto_increment;primary_key"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description;size:1024"`
	Status      bool      `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at; not null"`
}

// UserRoleModel basically describes a relationship between users and roles.
type UserRoleModel struct {
	UserID    uint64          `gorm:"primary_key"`
	RoleID    uint64          `gorm:"primary_key"`
	Status    bool            `gorm:"column:status"`
	CreatedAt time.Time       `gorm:"column:created_at; not null"`
	User      users.UserModel `gorm:"foreignkey:UserID"`
	Role      RoleModel       `gorm:"foreignkey:RoleID"`
}

// Update properties for RoleModel.
func (r *RoleModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(r).Update(data).Error
	return err
}

// Update properties for UserRoleModel.
func (ur *UserRoleModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(ur).Update(data).Error
	return err
}

// FindOneRole input the condition and It will return RoleModel.
func FindOneRole(condition interface{}) (RoleModel, error) {
	db := common.GetDB()
	var model RoleModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// FindOneUserRole input the condition and it will return UserRoleModel.
func FindOneUserRole(condition interface{}) (UserRoleModel, error) {
	db := common.GetDB()
	var model UserRoleModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveOne RoleModel will be saved in database returning error if It couldnt save or
// nil if everything was done.
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
