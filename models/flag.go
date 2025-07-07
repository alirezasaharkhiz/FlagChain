package models

import (
	"gorm.io/gorm"
)

type Flag struct {
	gorm.Model
	Name         string       `gorm:"uniqueIndex;not null"`
	Enabled      bool         `gorm:"default:false"`
	Dependencies []Dependency `gorm:"foreignKey:FlagID;constraint:OnDelete:CASCADE"`
}
