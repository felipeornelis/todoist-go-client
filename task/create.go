package task

import (
	"errors"
	"time"

	"github.com/felipeornelis/todoist-go-client/pkg"
)

type NewParams struct {
	ID           string
	ProjectID    string
	SectionID    string
	Content      string
	Description  string
	IsCompleted  bool
	Labels       []string
	ParentID     string
	Order        uint8
	Priority     uint8
	Due          taskDue
	URL          string
	CommentCount int
	CreatedAt    string
	AssigneeID   string
	AssignerID   string
	Duration     taskDuration
}

func (t Task) New(params NewParams) (Task, error) {
	if params.Content == "" {
		return Task{}, errors.New("`Content` field is required")
	}

	return Task{
		ID:           pkg.NewUUID(),
		ProjectID:    params.ProjectID,
		SectionID:    params.SectionID,
		Content:      params.Content,
		Description:  params.Description,
		IsCompleted:  params.IsCompleted,
		Labels:       params.Labels,
		ParentID:     params.ParentID,
		Order:        params.Order,
		Priority:     params.Priority,
		Due:          params.Due,
		URL:          params.URL,
		CommentCount: params.CommentCount,
		CreatedAt:    time.Now().String(),
		AssigneeID:   params.AssigneeID,
		AssignerID:   params.AssignerID,
		Duration:     params.Duration,
	}, nil
}
