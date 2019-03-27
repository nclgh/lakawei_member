package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

const UserTableName = "user"

type User struct {
	Id         int64     `gorm:"primary_key;not null;auto_increment"`
	Name       string    `gorm:"type:varchar(255);not null;default:''"`
	Username   string    `gorm:"type:varchar(255);not null;default:''"`
	Password   string    `gorm:"type:varchar(255);not null;default:''"`
	CreateTime time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func QueryUserByPwd(db *gorm.DB, username string) (*User, error) {
	userObj := &User{}
	err := db.Table(UserTableName).Where("username = ?", username).First(userObj).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return userObj, nil
}
