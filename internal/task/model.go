package task

import (
	"time"

	"github.com/uptrace/bun"
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks"`
	ID            string       `bun:"id,pk,unique,notnull"`
	ProjectID     string       `bun:"project_id,notnull"`
	StatusID      string       `bun:"status_id,notnull"`
	StatusName    string       `bun:"status_name,scanonly"`
	AssigneeID    string       `bun:"assigned_to,notnull"`
	AssigneeName  string       `bun:"assignee_name,scanonly"`
	ReporterID    string       `bun:"reported_to,notnull"`
	Priority      TaskPriority `bun:"priority,notnull"`
	Name          string       `bun:"name,notnull"`
	Description   string       `bun:"description"`
	Thumbnail     string       `bun:"thumbnail"`
	DueDate       time.Time    `bun:"due_date,nullzero"`
	CompletedDate time.Time    `bun:"completed_date,nullzero"`
	CreatedAt     time.Time    `bun:"created_at,notnull,default:current_timestamp"`
	CreatedBy     string       `bun:"created_by"`
	UpdatedAt     time.Time    `bun:"updated_at,notnull,default:current_timestamp"`
	UpdatedBy     string       `bun:"updated_by"`
	IsDeleted     bool         `bun:"is_deleted,default:false"`
}
