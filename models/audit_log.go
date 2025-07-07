package models

import (
	"gorm.io/gorm"
)

type AuditLog struct {
	gorm.Model
	FlagID uint   `gorm:"index;not null"`
	Action string `gorm:"size:50"`
	Actor  string `gorm:"size:100"`
	Reason string `gorm:"type:text"`
}
