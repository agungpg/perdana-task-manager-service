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
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	project := &Project{
		ID:          uuid.New().String(),
		Name:        name,
		Thumbnail:   thumbnail,
		Description: description,
		CreatedAt:   time.Now(),
		CreatedBy:   userID,
		UpdatedAt:   time.Now(),
	}

	err = s.repo.CreateProject(ctx, project, &tx)
	if err != nil {
		return err
	}
	projectMember := &ProjectMembers{
		ID:               uuid.New().String(),
		ProjectID:        project.ID,
		UserID:           userID,
		IsAdmin:          true,
		InvitationStatus: "accepted",
	}
	err = s.repo.AddProjectMember(ctx, projectMember, &tx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *Service) GetProjectByID(ctx context.Context, id string) (*Project, error) {
	project, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Service) GetProjectList(ctx context.Context, page, pageSize int, name string) ([]*Project, int, error) {
	projects, total, err := s.repo.GetProjectList(ctx, page, pageSize, name)
	if err != nil {
		return nil, 0, err
	}
	return projects, total, nil
}

func (s *Service) InviteProjectMember(ctx context.Context, userId, projectId string, isAdmin bool) error {
	projectMember := &ProjectMembers{
		ID:               uuid.New().String(),
		ProjectID:        projectId,
		UserID:           userId,
		IsAdmin:          isAdmin,
		InvitationStatus: "pending",
	}

	return s.repo.AddProjectMember(ctx, projectMember, nil)
}
