package model

import (
    "gorm.io/gorm"
    "time"
)

type V2 struct {
    ID        uint   `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *V2) TableName() string {
    return "v2"
 }
