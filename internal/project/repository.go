package project

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

func (r *Repository) BeginTx(ctx context.Context) (*bun.Tx, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *Repository) CreateProject(ctx context.Context, project *Project, tx *bun.Tx) error {
	runner := utils.GetQueryRunner(tx, r.db)
	// if tx != nil {
	// 	_, err := runner.NewInsert().Model(project).Exec(ctx)
	// 	return err
	// }
	_, err := runner.NewInsert().Model(project).Exec(ctx)
	return err
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Project, error) {
	project := new(Project)
	err := r.db.NewSelect().Model(project).Where("id = ?", id).Where("is_deleted = FALSE").Scan(ctx)
	return project, err
}

// add pagination support
func (r *Repository) GetProjectList(ctx context.Context, page, pageSize int, name string) ([]*Project, int, error) {
	var projects []*Project
	query := r.db.NewSelect().Model(&projects).Where("is_deleted = FALSE")
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	err = query.Scan(ctx)
	return projects, total, err
}

func (r *Repository) AddProjectMember(ctx context.Context, member *ProjectMembers, tx *bun.Tx) error {
	runner := utils.GetQueryRunner(tx, r.db)
	_, err := runner.NewInsert().Model(member).Exec(ctx)

	return err
}
