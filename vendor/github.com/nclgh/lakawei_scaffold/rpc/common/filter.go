package common

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
)

type Filter struct {
	TF []*TimeFilter
	IF []*InFilter
}

func (f *Filter) Filter(db *gorm.DB) *gorm.DB {
	for _, v := range f.TF {
		db = v.Filter(db)
	}
	for _, v := range f.IF {
		db = v.Filter(db)
	}
	return db
}

const TfFormat = "%s >= ? AND %s <= ?"

// timeFilter
type TimeFilter struct {
	Field string
	Start time.Time
	End   time.Time
}

func (tf *TimeFilter) Filter(db *gorm.DB) *gorm.DB {
	return db.Where(fmt.Sprintf(TfFormat, tf.Field, tf.Field), tf.Start, tf.End)
}

// InFilter
type InFilter struct {
	Field     string
	Condition interface{}
}

const IfFormat = "%s in (?)"

func (i *InFilter) Filter(db *gorm.DB) *gorm.DB {
	return db.Where(fmt.Sprintf(IfFormat, i.Field), i.Condition)
}
