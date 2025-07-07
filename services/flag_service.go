package services

import (
	"fmt"
	"github.com/alirezasaharkhiz/FlagChain/models"
	"github.com/alirezasaharkhiz/FlagChain/repositories"
)

type FeatureFlagService struct {
	FlagRepo  repositories.FlagRepository
	DepRepo   repositories.DependencyRepository
	AuditRepo repositories.AuditRepository
}

func NewFeatureFlagService(fr repositories.FlagRepository, dr repositories.DependencyRepository, ar repositories.AuditRepository) *FeatureFlagService {
	return &FeatureFlagService{FlagRepo: fr, DepRepo: dr, AuditRepo: ar}
}

func (s *FeatureFlagService) CreateFlag(name string, deps []string, actor string) (*models.Flag, error) {
	//check dependencies
	var depFlags []*models.Flag
	for _, dn := range deps {
		depFlag, err := s.FlagRepo.FindByName(dn)
		if err != nil {
			return nil, fmt.Errorf("dependency %q does not exist: %w", dn, err)
		}
		depFlags = append(depFlags, depFlag)
	}

	flag := &models.Flag{Name: name}
	if err := s.FlagRepo.Create(flag); err != nil {
		return nil, err
	}

	//assign dependencies to flag
	for _, depFlag := range depFlags {
		dep := &models.Dependency{FlagID: flag.ID, DependsOnID: depFlag.ID}
		if err := s.DepRepo.Add(dep); err != nil {
			return nil, err
		}
	}

	//log audit
	if err := s.AuditRepo.Log(&models.AuditLog{
		FlagID: flag.ID,
		Action: "create",
		Actor:  actor,
		Reason: "initial creation",
	}); err != nil {
		return flag, err
	}

	return flag, nil
}

func (s *FeatureFlagService) ToggleFlag(id uint, actor string) (*models.Flag, error) {
	flag, err := s.FlagRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if flag.Enabled {
		return flag, s.disable(flag, actor)
	}
	return flag, s.enable(flag, actor)
}

func (s *FeatureFlagService) enable(flag *models.Flag, actor string) error {
	// چک کردن فعال بودن همه‌ی dependency ها
	deps, _ := s.DepRepo.ListWhere("flag_id = ?", 5)
	missing := []string{}
	for _, d := range deps {
		depFlag, _ := s.FlagRepo.FindByID(d.DependsOnID)
		if !depFlag.Enabled {
			missing = append(missing, fmt.Sprintf("name:%s id:%d", depFlag.Name, depFlag.ID))
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("Missing active dependencies: %v", missing)
	}

	flag.Enabled = true
	if err := s.FlagRepo.Update(flag); err != nil {
		return err
	}

	_ = s.AuditRepo.Log(&models.AuditLog{
		FlagID: flag.ID, Action: "toggle_on", Actor: actor, Reason: "enabled",
	})

	return nil
}

func (s *FeatureFlagService) disable(flag *models.Flag, actor string) error {
	// غیرفعال کردن خودش
	flag.Enabled = false
	if err := s.FlagRepo.Update(flag); err != nil {
		return err
	}

	_ = s.AuditRepo.Log(&models.AuditLog{
		FlagID: flag.ID, Action: "toggle_off", Actor: actor, Reason: "disabled",
	})

	dependents, _ := s.DepRepo.ListWhere("depends_on_id = ?", flag.ID)
	for _, d := range dependents {
		depFlag, _ := s.FlagRepo.FindByID(d.FlagID)
		if depFlag.Enabled {
			if err := s.disable(depFlag, actor); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *FeatureFlagService) AddDependency(flagID, dependsOnID uint) error {
	if s.hasCircularDependency(dependsOnID, flagID, map[uint]bool{}) {
		return fmt.Errorf("Circular dependency detected between %d and %d", flagID, dependsOnID)
	}

	exists, err := s.DepRepo.Exists(flagID, dependsOnID)
	if err != nil {
		return fmt.Errorf("failed to check existing dependency: %w", err)
	}
	if exists {
		return fmt.Errorf("dependency from %d to %d already exists", flagID, dependsOnID)
	}

	dep := &models.Dependency{
		FlagID:      flagID,
		DependsOnID: dependsOnID,
	}
	return s.DepRepo.Add(dep)
}

func (s *FeatureFlagService) hasCircularDependency(currentID, targetID uint, visited map[uint]bool) bool {
	if currentID == targetID {
		return true
	}

	if visited[currentID] {
		return false
	}
	visited[currentID] = true

	deps, _ := s.DepRepo.ListWhere("flag_id = ?", currentID)
	for _, d := range deps {
		if s.hasCircularDependency(d.DependsOnID, targetID, visited) {
			return true
		}
	}
	return false
}

func (s *FeatureFlagService) ListFlags() ([]models.Flag, error) {
	return s.FlagRepo.ListAll()
}

func (s *FeatureFlagService) GetHistory(flagID uint) ([]models.AuditLog, error) {
	return s.AuditRepo.ListFor(flagID)
}
