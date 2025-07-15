package project

import (
	"time"

	"github.com/uptrace/bun"
)

type Project struct {
	bun.BaseModel `bun:"table:projects"`
	ID            string    `bun:"id,pk,unique,notnull"`
	Name          string    `bun:"name,unique,notnull"`
	Thumbnail     string    `bun:"thumbnail,notnull"`
	Description   string    `bun:"description"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	CreatedBy     string    `bun:"created_by"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	UpdatedBy     string    `bun:"updated_by"`
}
