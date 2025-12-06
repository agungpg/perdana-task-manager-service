package task

type UserSettingsDto struct {
	ID              string `json:"id"`
	UserId          string `json:"user_id"`
	ActiveProjectId string `json:"active_project_id"`
	Theme           Theme  `json:"theme"`
	Language        string `json:"language"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
