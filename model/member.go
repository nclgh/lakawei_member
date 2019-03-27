package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
)

const MemberTableName = "member"

type Member struct {
	Id           int64     `gorm:"primary_key;not null;auto_increment"`
	Code         string    `gorm:"type:varchar(255);not null;default:''"`
	Name         string    `gorm:"type:varchar(255);not null;default:''"`
	DepartmentId int64     `gorm:"type:bigint(20) unsigned;not null;default:0"`
	CreateTime   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func InsertMember(db *gorm.DB, code, name string, departmentId int64) error {
	// 查询部门是否存在
	ret, err := GetDepartmentById(db, []int64{departmentId})
	if err != nil {
		return err
	}
	if len(ret) <= 0 {
		return fmt.Errorf("department not exist")
	}

	mem := &Member{
		Code:         code,
		Name:         name,
		DepartmentId: departmentId,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	err = db.Table(MemberTableName).Create(mem).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteMember(db *gorm.DB, id int64) error {
	return db.Table(MemberTableName).Delete(&Member{}, "id = ?", id).Error
}

func QueryMember(db *gorm.DB, mem *Member, page, pageSize int64) ([]*Member, int64, error) {
	retDept := make([]*Member, 0)
	totalCnt := int64(0)
	err := db.Table(MemberTableName).Where(mem).Offset(page * pageSize).Limit(pageSize).Find(&retDept).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Table(MemberTableName).Where(mem).Count(&totalCnt).Error
	if err != nil {
		return nil, 0, err
	}
	return retDept, totalCnt, nil
}

func GetMemberById(db *gorm.DB, ids []int64) ([]*Member, error) {
	ret := make([]*Member, 0)
	err := db.Where("id in (?)", ids).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, err
}
