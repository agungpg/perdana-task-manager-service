package project

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Description string `json:"description,omitempty"`
}

type InviteProjectMemberRequest struct {
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
	IsAdmin   bool   `json:"is_admin"`
}
