package project

import (
	"time"

	"github.com/uptrace/bun"
)

type InvitationStatus string

const (
	InvitationPending  InvitationStatus = "pending"
	InvitationAccepted InvitationStatus = "accepted"
	InvitationRejected InvitationStatus = "rejected"
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
	StatusDefault string    `bun:"status_default,nullzero"`
	DoneStatusID  string    `bun:"done_status_id,nullzero"`
}

type ProjectMembers struct {
	bun.BaseModel    `bun:"table:project_members"`
	ID               string           `bun:"id,pk,unique,notnull"`
	ProjectID        string           `bun:"project_id,notnull"`
	UserID           string           `bun:"user_id,notnull"`
	CreatedAt        time.Time        `bun:"created_at,notnull,default:current_timestamp"`
	CreatedBy        string           `bun:"created_by"`
	UpdatedAt        time.Time        `bun:"updated_at,notnull,default:current_timestamp"`
	UpdatedBy        string           `bun:"updated_by"`
	IsAdmin          bool             `bun:"is_admin"`
	InvitationStatus InvitationStatus `bun:"invitation_status" json:"invitation_status"`
}

type ProjectStatuses struct {
	bun.BaseModel `bun:"table:project_statuses"`
	ID            string    `bun:"id,pk,unique,notnull"`
	Name          string    `bun:"name,notnull"`
	ProjectID     string    `bun:"project_id,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	CreatedBy     string    `bun:"created_by"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	UpdatedBy     string    `bun:"updated_by"`
}
