package repositories

import (
	"github.com/alirezasaharkhiz/FlagChain/models"
	"gorm.io/gorm"
)

type GormFlagRepository struct{ db *gorm.DB }

func NewFlagRepository(db *gorm.DB) FlagRepository { return &GormFlagRepository{db} }
func (r *GormFlagRepository) Create(flag *models.Flag) error {
	return r.db.Create(flag).Error
}
func (r *GormFlagRepository) FindByName(name string) (*models.Flag, error) {
	var f models.Flag
	err := r.db.Preload("Dependencies").Where("name = ?", name).First(&f).Error
	return &f, err
}
func (r *GormFlagRepository) FindByID(id uint) (*models.Flag, error) {
	var f models.Flag
	err := r.db.Preload("Dependencies").First(&f, id).Error
	return &f, err
}
func (r *GormFlagRepository) Update(flag *models.Flag) error {
	return r.db.Save(flag).Error
}
func (r *GormFlagRepository) ListAll() ([]models.Flag, error) {
	var flags []models.Flag
	err := r.db.Preload("Dependencies").Find(&flags).Error
	return flags, err
}
