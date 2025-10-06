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
func (s *Service) CreateProject(ctx context.Context, userID, name, thumbnail, description string, statuses []ProjectStatusRequest) error {
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

	defaultStatusId := ""
	doneStatusId := ""
	projectStatuses := make([]*ProjectStatuses, len(statuses))
	for i, status := range statuses {
		projectStatuses[i] = &ProjectStatuses{
			ID:         uuid.New().String(),
			Name:       status.Name,
			ProjectID:  project.ID,
			CreatedAt:  time.Now(),
			CreatedBy:  userID,
			UpdatedAt:  time.Now(),
			OrderIndex: status.OrderIndex,
		}
		if status.IsDefault {
			defaultStatusId = projectStatuses[i].ID
		}
		if status.IsDone {
			doneStatusId = projectStatuses[i].ID
		}
	}

	err = s.repo.CreateProject(ctx, project, &tx)
	if err != nil {
		return err
	}
	err = s.repo.AddMultipleProjectStatuses(ctx, projectStatuses, &tx)
	if err != nil {
		return err
	}
	err = s.repo.SetProjectStatusDefault(ctx, project.ID, defaultStatusId, &tx)
	if err != nil {
		return err
	}
	if doneStatusId != "" {
		err = s.repo.SetProjectDoneStatus(ctx, project.ID, doneStatusId, &tx)
		if err != nil {
			return err
		}
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

func (s *Service) AcceptProjectInvitation(ctx context.Context, projectId, memberId string) error {
	members, err := s.repo.GetProjectMember(ctx, projectId, memberId, nil)
	if err != nil || len(members) == 0 {
		return errors.New("failed to accept project invitation: member not found")
	}

	procjectMember := &ProjectMembers{}
	found := false
	for _, member := range members {
		fmt.Println(member.UserID)
		fmt.Println(member.InvitationStatus)
		if member.UserID == memberId && member.InvitationStatus == "pending" {
			found = true
			procjectMember = member
			break
		}
	}

	if !found {
		return errors.New("failed to accept project invitation: member not found")
	}
	return s.repo.AcceptProjectInvitation(ctx, procjectMember.ID)
}

func (s *Service) RemoveProjectMember(ctx context.Context, projectId, memberId, userId string) error {
	members, err := s.repo.GetProjectMember(ctx, projectId, userId, nil)
	if err != nil || len(members) == 0 {
		return errors.New("failed to rempve project member: you don't have authority to do it")
	}
	procjectMember := &ProjectMembers{}
	for _, member := range members {
		if member.UserID == userId {
			procjectMember = member
			break
		}
	}

	if !procjectMember.IsAdmin {
		return errors.New("failed to rempve project member: you don't have authority to do it")
	}

	err = s.repo.RemoveProjectMember(ctx, memberId)

	return err
}
