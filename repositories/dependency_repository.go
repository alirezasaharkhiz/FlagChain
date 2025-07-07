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
func (r *GormDependencyRepository) ListWhere(condition string, args ...interface{}) ([]models.Dependency, error) {
	var deps []models.Dependency
	err := r.db.Where(condition, args...).Find(&deps).Error
	return deps, err
}
func (r *GormDependencyRepository) RemoveAllFor(flagID uint) error {
	return r.db.Where("flag_id = ?", flagID).Delete(&models.Dependency{}).Error
}

func (r *GormDependencyRepository) Exists(flagID, dependsOnID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Dependency{}).
		Where("flag_id = ? AND depends_on_id = ?", flagID, dependsOnID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
