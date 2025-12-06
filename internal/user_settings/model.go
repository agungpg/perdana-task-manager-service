package task

import (
	"github.com/uptrace/bun"
)

type Theme string

const (
	Light Theme = "light"
	Dark  Theme = "dark"
)

type UserSettings struct {
	bun.BaseModel   `bun:"table:user_settings"`
	ID              string `bun:"id,pk,unique,notnull"`
	ActiveProjectId string `bun:"active_project_id,notnull"`
	UserId          string `bun:"user_id,notnull"`
	Theme           Theme  `bun:"theme,notnull"`
	Language        string `bun:"language,scanonly"`
	CreatedAt       string `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt       string `bun:"updated_at,notnull,default:current_timestamp"`
}
