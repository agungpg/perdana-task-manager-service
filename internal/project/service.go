package project

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

// make thumbnail and description optional
func (s *Service) CreateProject(ctx context.Context, userID, name, thumbnail, description string) error {

	project := &Project{
		ID:          uuid.New().String(),
		Name:        name,
		Thumbnail:   thumbnail,
		Description: description,
		CreatedAt:   time.Now(),
		CreatedBy:   userID,
		UpdatedAt:   time.Now(),
	}

	return s.repo.CreateProject(ctx, project)
}

func (s *Service) GetProjectByID(ctx context.Context, id string) (*Project, error) {
	project, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) GetProjectList(ctx context.Context, page, pageSize int, name string) ([]*Project, error) {
	projects, err := s.repo.GetProjectList(ctx, page, pageSize, name)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
