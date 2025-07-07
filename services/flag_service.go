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
	flag := &models.Flag{Name: name}
	if err := s.FlagRepo.Create(flag); err != nil {
		return nil, err
	}
	for _, dn := range deps {
		depFlag, err := s.FlagRepo.FindByName(dn)
		if err != nil {
			return nil, err
		}
		dep := &models.Dependency{FlagID: flag.ID, DependsOnID: depFlag.ID}
		if err := s.DepRepo.Add(dep); err != nil {
			return nil, err
		}
	}
	err := s.AuditRepo.Log(&models.AuditLog{FlagID: flag.ID, Action: "create", Actor: actor, Reason: "initial creation"})
	if err != nil {
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
	deps, _ := s.DepRepo.ListFor(flag.ID)
	missing := []string{}
	for _, d := range deps {
		depFlag, _ := s.FlagRepo.FindByID(d.DependsOnID)
		if !depFlag.Enabled {
			missing = append(missing, depFlag.Name)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("Missing active dependencies: %v", missing)
	}
	flag.Enabled = true
	if err := s.FlagRepo.Update(flag); err != nil {
		return err
	}
	_ = s.AuditRepo.Log(&models.AuditLog{FlagID: flag.ID, Action: "toggle_on", Actor: actor, Reason: "enabled"})
	return nil
}

func (s *FeatureFlagService) disable(flag *models.Flag, actor string) error {
	flag.Enabled = false
	if err := s.FlagRepo.Update(flag); err != nil {
		return err
	}
	_ = s.AuditRepo.Log(&models.AuditLog{FlagID: flag.ID, Action: "toggle_off", Actor: actor, Reason: "disabled"})
	// TODO: Cascade disable
	return nil
}

func (s *FeatureFlagService) ListFlags() ([]models.Flag, error) {
	return s.FlagRepo.ListAll()
}

func (s *FeatureFlagService) GetHistory(flagID uint) ([]models.AuditLog, error) {
	return s.AuditRepo.ListFor(flagID)
}
