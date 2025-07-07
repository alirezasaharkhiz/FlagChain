package models

import (
	"gorm.io/gorm"
)

type Dependency struct {
	gorm.Model
	FlagID      uint `gorm:"index;not null"`
	DependsOnID uint `gorm:"index;not null"`
}
