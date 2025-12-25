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

func (s *Service) GetTaskList(ctx context.Context, projectId, name, statusId string) ([]*TaskItem, error) {
	tasks, err := s.repo.GetTaskList(ctx, projectId, name, statusId)
	if err != nil {
		return nil, err
	}

	taskItems := make([]*TaskItem, len(tasks))
	for i, task := range tasks {
		remainingTime := time.Since(task.DueDate)
		remHours := remainingTime.Hours()
		days := remHours / 24
		hours := float64(int(remHours) % 24)
		remainingTimeStr := ""
		if days > 0 {
			remainingTimeStr = remainingTimeStr + fmt.Sprintf("%.0f", days) + "days"
		}
		if hours > 0 {
			remainingTimeStr = remainingTimeStr + " " + fmt.Sprintf("%.0f", hours) + "hours"
		}

		taskItems[i] = &TaskItem{
			ID:                    task.ID,
			Name:                  task.Name,
			AssigneeID:            task.AssigneeID,
			AssigneeName:          task.AssigneeName,
			StatusID:              task.StatusID,
			StatusName:            task.StatusName,
			DueDate:               task.DueDate.Format(time.RFC3339),
			RemainingTime:         remainingTimeStr,
			Thumbnail:             task.Thumbnail,
			Priority:              string(task.Priority),
			ProjectName:           string(task.ProjectName),
			TotalSubTask:          int(task.TotalSubTask),
			TotalCompletedSubTask: int(task.TotalCompletedSubTask),
		}
	}
	return taskItems, nil
}
func (s *Service) GetTaskByID(ctx context.Context, id string) (*TaskDetail, error) {
	task, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &TaskDetail{
		ID:           task.ID,
		Name:         task.Name,
		AssigneeID:   task.AssigneeID,
		AssigneeName: task.AssigneeName,
		StatusID:     task.StatusID,
		StatusName:   task.StatusName,
		DueDate:      task.DueDate,
		Thumbnail:    task.Thumbnail,
		Priority:     string(task.Priority),
		Description:  task.Description,
		CreatedAt:    task.CreatedAt,
		CreatedBy:    task.CreatedBy,
		UpdatedAt:    task.UpdatedAt,
		UpdatedBy:    task.UpdatedBy,
		ReporterID:   task.ReporterID,
		ReporterName: task.ReporterName,
		ProjectName:  task.ProjectName,
	}, nil
}
