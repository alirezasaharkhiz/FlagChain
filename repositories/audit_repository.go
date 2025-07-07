package repositories

import (
	"github.com/alirezasaharkhiz/FlagChain/models"
	"gorm.io/gorm"
)

type GormAuditRepository struct{ db *gorm.DB }

func NewAuditRepository(db *gorm.DB) AuditRepository { return &GormAuditRepository{db} }
func (r *GormAuditRepository) Log(entry *models.AuditLog) error {
	return r.db.Create(entry).Error
}
func (r *GormAuditRepository) ListFor(flagID uint) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Where("flag_id = ?", flagID).Find(&logs).Error
	return logs, err
}
