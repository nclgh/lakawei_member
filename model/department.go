package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
)

const DepartmentTableName = "department"

type Department struct {
	Id         int64     `gorm:"primary_key;not null;auto_increment"`
	Code       string    `gorm:"type:varchar(255);not null;default:''"`
	Name       string    `gorm:"type:varchar(255);not null;default:''"`
	CreateTime time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func InsertDepartment(db *gorm.DB, code, name string) error {
	dept := &Department{
		Code:       code,
		Name:       name,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := db.Table(DepartmentTableName).Create(dept).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDepartment(db *gorm.DB, code string) error {
	// 查询部门人员是否存在
	_, cnt, err := QueryMember(db, &Member{
		DepartmentCode: code,
	}, 0, 1)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return fmt.Errorf("still exist member in this department")
	}
	return db.Table(DepartmentTableName).Delete(&Department{}, "code = ?", code).Error
}

func QueryDepartment(db *gorm.DB, dept *Department, page, pageSize int64) ([]*Department, int64, error) {
	retDept := make([]*Department, 0)
	totalCnt := int64(0)
	err := db.Table(DepartmentTableName).Where(dept).Offset(page * pageSize).Limit(pageSize).Find(&retDept).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Table(DepartmentTableName).Where(dept).Count(&totalCnt).Error
	if err != nil {
		return nil, 0, err
	}
	return retDept, totalCnt, nil
}

func GetDepartmentByCode(db *gorm.DB, codes []string) ([]*Department, error) {
	retDept := make([]*Department, 0)
	err := db.Where("code in (?)", codes).Find(&retDept).Error
	if err != nil {
		return nil, err
	}
	return retDept, err
}
