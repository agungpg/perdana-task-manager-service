package task

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetUserSettings(ctx context.Context, userId string) (*UserSettingsDto, error) {
	userSettings, err := s.repo.GetUserSettings(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &UserSettingsDto{
		ID:              userSettings.ID,
		UserId:          userSettings.UserId,
		ActiveProjectId: userSettings.ActiveProjectId,
		Theme:           userSettings.Theme,
		Language:        userSettings.Language,
		CreatedAt:       userSettings.CreatedAt,
		UpdatedAt:       userSettings.UpdatedAt,
	}, nil
}
