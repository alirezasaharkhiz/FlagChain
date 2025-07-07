package repositories

import (
	"github.com/alirezasaharkhiz/FlagChain/models"
	"gorm.io/gorm"
)

type GormDependencyRepository struct{ db *gorm.DB }

func NewDependencyRepository(db *gorm.DB) DependencyRepository { return &GormDependencyRepository{db} }
func (r *GormDependencyRepository) Add(dep *models.Dependency) error {
	return r.db.Create(dep).Error
}
func (r *GormDependencyRepository) ListFor(flagID uint) ([]models.Dependency, error) {
	var deps []models.Dependency
	err := r.db.Where("flag_id = ?", flagID).Find(&deps).Error
	return deps, err
}
func (r *GormDependencyRepository) RemoveAllFor(flagID uint) error {
	return r.db.Where("flag_id = ?", flagID).Delete(&models.Dependency{}).Error
}
