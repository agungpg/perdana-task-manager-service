package task

import (
	"context"
	"database/sql"
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

func (s *Service) CreateTask(ctx context.Context, projectId, statusId, assigneeId, reporterId, priority, name, description, thumbnail string, dueDate time.Time, userId string) error {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if rErr := tx.Rollback(); rErr != nil && rErr != sql.ErrTxDone {
			fmt.Printf("rollback error: %v", rErr)
		}
	}()

	task := &Task{
		ID:          uuid.New().String(),
		ProjectID:   projectId,
		StatusID:    statusId,
		AssigneeID:  assigneeId,
		ReporterID:  reporterId,
		Priority:    TaskPriority(priority),
		Name:        name,
		Description: description,
		Thumbnail:   thumbnail,
		DueDate:     dueDate,
		CreatedAt:   time.Now(),
		CreatedBy:   userId,
		UpdatedAt:   time.Now(),
	}

	err = s.repo.CreateTask(ctx, task, &tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
