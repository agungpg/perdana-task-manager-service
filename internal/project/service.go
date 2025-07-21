package project

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	defer func() {
		if rErr := tx.Rollback(); rErr != nil && rErr != sql.ErrTxDone {
			fmt.Printf("rollback error: %v", rErr)
		}
	}()

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

func (s *Service) InviteProjectMember(ctx context.Context, userId, memberId, projectId string, isAdmin bool) error {
	members, err := s.repo.GetProjectMember(ctx, projectId, "", nil)
	if err != nil || len(members) == 0 {
		return errors.New("failed to invite project member: something went wrong")
	}
	if isAlreadyMember(memberId, members) {
		return errors.New("failed to invite project member: user is already a member")
	}

	_, err = s.validateInviterIsAdmin(userId, members)
	if err != nil {
		return fmt.Errorf("failed to invite project member: %w", err)
	}

	projectMember := &ProjectMembers{
		ID:               uuid.New().String(),
		ProjectID:        projectId,
		UserID:           memberId,
		IsAdmin:          isAdmin,
		InvitationStatus: "pending",
	}

	return s.repo.AddProjectMember(ctx, projectMember, nil)
}

func (s *Service) validateInviterIsAdmin(userId string, members []*ProjectMembers) (*ProjectMembers, error) {
	for _, member := range members {
		if member.UserID == userId {
			if member.IsAdmin {
				return member, nil
			}
			return nil, errors.New("unauthorized access: inviter is not admin")
		}
	}
	return nil, errors.New("unauthorized access: inviter not found")
}

func isAlreadyMember(memberId string, members []*ProjectMembers) bool {
	for _, member := range members {
		if member.UserID == memberId {
			return true
		}
	}
	return false
}
