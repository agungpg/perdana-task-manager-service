package task

type CreateTaskRequest struct {
	ProjectID   string `json:"project_id"`
	StatusID    string `json:"status_id"`
	AssigneeID  string `json:"assignee_id"`
	ReporterID  string `json:"reporter_id"`
	Priority    string `json:"priority"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	DueDate     string `json:"due_date"`
}

type TaskItem struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	AssigneeID            string `json:"assigneeId"`
	AssigneeName          string `json:"assigneeName"`
	StatusID              string `json:"statusId"`
	StatusName            string `json:"statusName"`
	DueDate               string `json:"dueDate"`
	RemainingTime         string `json:"remainingTime"`
	Thumbnail             string `json:"thumbnail"`
	Priority              string `json:"priority"`
	ProjectId             string `json:"projectId"`
	ProjectName           string `json:"projectName"`
	TotalSubTask          int    `json:"totalSubtasks"`
	TotalCompletedSubTask int    `json:"completedSubtasks"`
}

type TaskDetail struct {
	ID            string `json:"id" bun:"id"`
	Name          string `json:"name" bun:"name"`
	AssigneeID    string `json:"assignee_id" bun:"assigned_to"`
	AssigneeName  string `json:"assignee_name" bun:"assignee_name,scanonly"`
	StatusID      string `json:"status_id" bun:"status_id"`
	StatusName    string `json:"status_name" bun:"status_name,scanonly"`
	DueDate       string `json:"due_date" bun:"due_date"`
	CompletedDate string `json:"completed_date" bun:"completed_date"`
	Thumbnail     string `json:"thumbnail" bun:"thumbnail"`
	Priority      string `json:"priority" bun:"priority"`
	Description   string `json:"description" bun:"description"`
	CreatedAt     string `json:"created_at" bun:"created_at"`
	CreatedBy     string `json:"created_by" bun:"created_by"`
	UpdatedAt     string `json:"updated_at" bun:"updated_at"`
	UpdatedBy     string `json:"updated_by" bun:"updated_by"`
	ReporterID    string `json:"reporter_id" bun:"reported_to"`
	ReporterName  string `json:"reporter_name" bun:"reporter_name,scanonly"`
	ProjectName   string `json:"project_name" bun:"project_name,scanonly"`
}
