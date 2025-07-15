package project

import (
	"context"

	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateProject(ctx context.Context, project *Project) error {
	_, err := r.db.NewInsert().Model(project).Exec(ctx)
	return err
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Project, error) {
	project := new(Project)
	err := r.db.NewSelect().Model(project).Where("id = ?", id).Scan(ctx)
	return project, err
}

// add pagination support
func (r *Repository) GetProjectList(ctx context.Context, page, pageSize int, name string) ([]*Project, error) {
	var projects []*Project
	query := r.db.NewSelect().Model(&projects).
		Offset((page - 1) * pageSize).
		Limit(pageSize)
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	err := query.Scan(ctx)
	return projects, err
}
