package task

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

func (r *Repository) GetUserSettings(ctx context.Context, userId string) (*UserSettings, error) {
	var userSettings UserSettings
	err := r.db.NewSelect().Model(&userSettings).
		Where("user_id = ?", userId).
		Scan(ctx)

	return &userSettings, err
}
