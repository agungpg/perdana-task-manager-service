package project

import (
	"context"
	"time"

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

func (r *Repository) AcceptProjectInvitation(ctx context.Context, projectMemberId string) error {
	_, err := r.db.NewUpdate().
		Model(&ProjectMembers{}).
		Set("invitation_status = ?", "accepted").
		Set("updated_at = ?", time.Now()).
		Where("id = ?", projectMemberId).
		Exec(ctx)

	return err
}

func (r *Repository) RemoveProjectMember(ctx context.Context, member_id string) error {
	_, err := r.db.NewUpdate().
		Model(&ProjectMembers{}).
		Set("is_deleted = ?", true).
		Where("user_id = ?", member_id).
		Exec(ctx)
	return err
}

func (r *Repository) GetProjectMember(ctx context.Context, projectId, userId string, tx *bun.Tx) ([]*ProjectMembers, error) {
	runner := utils.GetQueryRunner(tx, r.db)
	var members []*ProjectMembers

	query := runner.NewSelect().
		Model(&members).
		Where("is_deleted = FALSE").
		Where("project_id = ?", projectId)

	if userId != "" {
		query.Where("user_id = ?", userId)
	}

	err := query.Scan(ctx)

	return members, err
}

func (r *Repository) SetProjectStatusDefault(ctx context.Context, projectId, statusId string, tx *bun.Tx) error {
	runner := utils.GetQueryRunner(tx, r.db)
	_, err := runner.NewUpdate().
		Model(&Project{}).
		Set("status_default = ?", statusId).
		Where("id = ?", projectId).
		Exec(ctx)
	return err
}

func (r *Repository) SetProjectDoneStatus(ctx context.Context, projectId, statusId string, tx *bun.Tx) error {
	runner := utils.GetQueryRunner(tx, r.db)
	_, err := runner.NewUpdate().
		Model(&Project{}).
		Set("done_status_id = ?", statusId).
		Where("id = ?", projectId).
		Exec(ctx)
	return err
}

func (r *Repository) GetProjectStatuses(ctx context.Context, projectId string) ([]*ProjectStatuses, error) {
	var statuses []*ProjectStatuses
	err := r.db.NewSelect().Model(&statuses).Where("project_id = ?", projectId).Scan(ctx)
	return statuses, err
}

func (r *Repository) AddMultipleProjectStatuses(ctx context.Context, statuses []*ProjectStatuses, tx *bun.Tx) error {
	runner := utils.GetQueryRunner(tx, r.db)
	_, err := runner.NewInsert().Model(&statuses).Exec(ctx)
	return err
}

func (r *Repository) RemoveProjectStatusesByProjectId(ctx context.Context, projectId string) error {
	_, err := r.db.NewUpdate().
		Model(&ProjectStatuses{}).
		Set("is_deleted = ?", true).
		Where("project_id = ?", projectId).
		Exec(ctx)
	return err
}
