package repositories

import "github.com/alirezasaharkhiz/FlagChain/models"

// FlagRepository ...
type FlagRepository interface {
	Create(flag *models.Flag) error
	FindByName(name string) (*models.Flag, error)
	FindByID(id uint) (*models.Flag, error)
	Update(flag *models.Flag) error
	ListAll() ([]models.Flag, error)
}

// DependencyRepository ...
type DependencyRepository interface {
	Add(dep *models.Dependency) error
	ListFor(flagID uint) ([]models.Dependency, error)
	RemoveAllFor(flagID uint) error
}

// AuditRepository ...
type AuditRepository interface {
	Log(entry *models.AuditLog) error
	ListFor(flagID uint) ([]models.AuditLog, error)
}
