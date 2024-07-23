package domain

type Priority string
type Status string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

const (
	New        Status = "new"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Priority    Priority `json:"priority"`
	Status      Status   `json:"status"`
	AssigneeID  string   `json:"assignee_id"`
	ProjectID   string   `json:"project_id"`
	CreatedAt   string   `json:"created_at"`
	CompletedAt string   `json:"completed_at"`
}
