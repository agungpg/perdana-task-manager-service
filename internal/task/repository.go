package task

import (
	"context"

	"github.com/agungpg/perdana-task-manager/pkg/utils"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateTask(ctx context.Context, task *Task, tx *bun.Tx) error {
	runner := utils.GetQueryRunner(tx, r.db)
	_, err := runner.NewInsert().Model(task).Exec(ctx)
	return err
}

func (r *Repository) GetTaskByID(ctx context.Context, id string) (*TaskDetail, error) {
	var task TaskDetail
	err := r.db.NewSelect().
		TableExpr("tasks AS task").
		Column("task.id", "task.name", "task.assigned_to", "task.status_id", "task.due_date", "task.thumbnail", "task.priority", "task.description", "task.completed_date").
		Column("task.created_at", "task.created_by", "task.updated_at", "task.updated_by", "task.reported_to").
		ColumnExpr("p.name AS project_name").
		ColumnExpr("ps.name AS status_name").
		ColumnExpr("u_assignee.username AS assignee_name").
		ColumnExpr("u_reporter.username AS reporter_name").
		Join("LEFT JOIN projects AS p ON p.id = task.project_id").
		Join("LEFT JOIN project_statuses AS ps ON ps.id = task.status_id").
		Join("LEFT JOIN project_members AS pm_assignee ON pm_assignee.id = task.assigned_to").
		Join("LEFT JOIN users AS u_assignee ON u_assignee.id = pm_assignee.user_id").
		Join("LEFT JOIN project_members AS pm_reporter ON pm_reporter.id = task.reported_to").
		Join("LEFT JOIN users AS u_reporter ON u_reporter.id = pm_reporter.user_id").
		Where("task.id = ?", id).
		Where("task.is_deleted = ?", false).
		Scan(ctx, &task)
	return &task, err
}

func (r *Repository) GetTaskList(ctx context.Context, projectId string) ([]*Task, error) {
	var tasks []*Task
	err := r.db.NewSelect().Model(&tasks).
		Column("task.id", "task.name", "task.assigned_to", "task.status_id", "task.due_date", "task.thumbnail", "task.priority").
		ColumnExpr("ps.name AS status_name").
		ColumnExpr("u.username AS assignee_name").
		Join("LEFT JOIN project_statuses AS ps ON ps.id = task.status_id").
		Join("LEFT JOIN project_members AS pm ON pm.id = task.assigned_to").
		Join("LEFT JOIN users AS u ON u.id = pm.user_id").
		Where("task.project_id = ?", projectId).
		Where("task.is_deleted = ?", false).
		Scan(ctx)
	return tasks, err
}
