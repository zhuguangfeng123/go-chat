package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        int64                 `gorm:"primarykey"         json:"id"`
	CreatedAt uint                  `gorm:"column:created_at"  json:"createdAt"`
	UpdatedAt uint                  `gorm:"column:updated_at;autoUpdateTime;"  json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at"  json:"deletedAt"`
}

type Strings []string

func (s *Strings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, s)
}

func (s Strings) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type Int64s []int64

func (i *Int64s) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, i)
}

func (i Int64s) Value() (driver.Value, error) {
	return json.Marshal(i)
}
