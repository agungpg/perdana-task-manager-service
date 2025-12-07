package project

type CreateProjectRequest struct {
	Name        string                 `json:"name"`
	Thumbnail   string                 `json:"thumbnail,omitempty"`
	Description string                 `json:"description,omitempty"`
	Statuses    []ProjectStatusRequest `json:"statuses,omitempty"`
}

type ProjectStatusRequest struct {
	Name       string `json:"name"`
	IsDefault  bool   `json:"is_default"`
	IsDone     bool   `json:"is_done"`
	OrderIndex int    `json:"order_index"`
}

type InviteProjectMemberRequest struct {
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
	IsAdmin   bool   `json:"is_admin"`
}

type AcceptMemberRequest struct {
	ProjectID string `json:"project_id"`
}
type RemoveMemberRequest struct {
	ProjectID string `json:"project_id"`
	MemberId  string `json:"member_id"`
}

type ProjectSummary struct {
	ProjectID   string         `json:"project_id"`
	ProjectName string         `json:"project_name"`
	Statuses    map[string]int `json:"statuses"` // key = status name, value = count
}
