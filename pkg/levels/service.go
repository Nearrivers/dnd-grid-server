package levels

import (
	"context"

	"github.com/Nearrivers/dnd-grid-server/pkg/models/repository"
)

type Service interface {
	DeleteLevel(int64) error
	GetLevel(int64) (repository.Levels, error)
	GetLevelWithEntities(int64) ([]repository.GetLevelWithEntitiesRow, error)
	GetLevels() ([]repository.Levels, error)
	NewLevel(repository.NewLevelParams) error
	UpdateLevel(repository.UpdateLevelParams) error
}

type service struct {
	levelRepository *repository.Queries
}

func NewService(lr *repository.Queries) Service {
	return &service{
		levelRepository: lr,
	}
}

func (s *service) DeleteLevel(id int64) error {
	ctx := context.Background()
	return s.levelRepository.DeleteLevel(ctx, id)
}

func (s *service) GetLevel(id int64) (repository.Levels, error) {
	ctx := context.Background()
	return s.levelRepository.GetLevel(ctx, id)
}

func (s *service) GetLevelWithEntities(id int64) ([]repository.GetLevelWithEntitiesRow, error) {
	ctx := context.Background()
	return s.levelRepository.GetLevelWithEntities(ctx, id)
}

func (s *service) GetLevels() ([]repository.Levels, error) {
	ctx := context.Background()
	return s.levelRepository.GetLevels(ctx)
}

func (s *service) NewLevel(newLevel repository.NewLevelParams) error {
	ctx := context.Background()
	return s.levelRepository.NewLevel(ctx, newLevel)
}

func (s *service) UpdateLevel(newLevelValues repository.UpdateLevelParams) error {
	ctx := context.Background()
	return s.levelRepository.UpdateLevel(ctx, newLevelValues)
}
