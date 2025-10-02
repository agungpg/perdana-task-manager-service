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

func (r *Repository) GetTaskByID(ctx context.Context, id string) (*Task, error) {
	var task Task
	err := r.db.NewSelect().Model(&task).Where("id = ?", id).Where("is_deleted = ?", false).Scan(ctx)
	return &task, err
}

func (r *Repository) GetTaskList(ctx context.Context, projectId string) ([]*Task, error) {
	var tasks []*Task
	err := r.db.NewSelect().Model(&tasks).Where("project_id = ?", projectId).Where("is_deleted = ?", false).Scan(ctx)
	return tasks, err
}
