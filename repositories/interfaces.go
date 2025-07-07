package repositories

import "github.com/alirezasaharkhiz/FlagChain/models"

type FlagRepository interface {
	Create(flag *models.Flag) error
	FindByName(name string) (*models.Flag, error)
	FindByID(id uint) (*models.Flag, error)
	Update(flag *models.Flag) error
	ListAll() ([]models.Flag, error)
}

type DependencyRepository interface {
	Add(dep *models.Dependency) error
	ListWhere(condition string, args ...interface{}) ([]models.Dependency, error)
	RemoveAllFor(flagID uint) error
	Exists(flagID, dependsOnID uint) (bool, error)
}

type AuditRepository interface {
	Log(entry *models.AuditLog) error
	ListFor(flagID uint) ([]models.AuditLog, error)
}
