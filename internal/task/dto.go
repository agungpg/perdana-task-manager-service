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
