package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-go-client/pkg"
)

type Task struct {
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

type taskDue struct {
	String      string
	Date        string
	IsRecurring bool
	Datetime    string
	Timezone    string
}

type taskDuration struct {
	Amount uint
	Unit   string
}

type AddTaskArgs struct {
	Content      string   `json:"content"`
	Description  string   `json:"description,omitempty"`
	ProjectID    string   `json:"project_id,omitempty"`
	SectionID    string   `json:"section_id,omitempty"`
	ParentID     string   `json:"parent_id,omitempty"`
	Order        uint8    `json:"order,omitempty"`
	Labels       []string `json:"labels,omitempty"`
	Priority     uint8    `json:"priority,omitempty"`
	DueString    string   `json:"due_string,omitempty"`
	DueDate      string   `json:"due_date,omitempty"`
	DueDatetime  string   `json:"due_datetime,omitempty"`
	DueLang      string   `json:"due_lang,omitempty"`
	AssigneeID   string   `json:"assignee_id,omitempty"`
	Duration     uint     `json:"duration,omitempty"`
	DurationUnit string   `json:"duration_unit,omitempty"`
}

func (t Todoist) AddTask(args AddTaskArgs) (Task, error) {
	if args.Content == "" {
		return Task{}, errors.New("`Content` field is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Task{}, err
	}

	url := "https://api.todoist.com/rest/v2/tasks"

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return Task{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return Task{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Task{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Task{}, nil
	}

	var task Task
	if err := json.Unmarshal(data, &task); err != nil {
		return Task{}, err
	}

	return Task{
		ID:           task.ID,
		ProjectID:    task.ProjectID,
		SectionID:    task.SectionID,
		Content:      task.Content,
		Description:  task.Description,
		IsCompleted:  task.IsCompleted,
		Labels:       task.Labels,
		ParentID:     task.ParentID,
		Order:        task.Order,
		Priority:     task.Priority,
		Due:          task.Due,
		URL:          task.URL,
		CommentCount: task.CommentCount,
		CreatedAt:    task.CreatedAt,
		AssigneeID:   task.AssigneeID,
		AssignerID:   task.AssignerID,
		Duration:     task.Duration,
	}, nil
}

func (t Todoist) GetTask(id string) (Task, error) {
	return Task{}, nil
}

func (t Todoist) GetTasks() ([]Task, error) {
	return []Task{}, nil
}

type UpdateTaskArgs struct {
	Content      string   `json:"content,omitempty"`
	Description  string   `json:"description,omitempty"`
	Labels       []string `json:"labels,omitempty"`
	Priority     uint8    `json:"priority,omitempty"`
	DueString    string   `json:"due_string,omitempty"`
	DueDate      string   `json:"due_date,omitempty"`
	DueDatetime  string   `json:"due_datetime,omitempty"`
	DueLang      string   `json:"due_lang,omitempty"`
	AssigneeID   string   `json:"assignee_id,omitempty"`
	Duration     uint     `json:"duration,omitempty"`
	DurationUnit string   `json:"duration_unit,omitempty"`
}

func (t Todoist) UpdateTask(args UpdateTaskArgs, id string) (Task, error) {
	if id == "" {
		return Task{}, errors.New("ID is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Task{}, err
	}

	url := fmt.Sprintf("https://api.todoist.com/rest/v2/tasks/%s", id)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return Task{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return Task{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Task{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Task{}, err
	}

	var task Task
	if err := json.Unmarshal(data, &task); err != nil {
		return Task{}, err
	}

	return Task{
		ID:           task.ID,
		ProjectID:    task.ProjectID,
		SectionID:    task.SectionID,
		Content:      task.Content,
		Description:  task.Description,
		IsCompleted:  task.IsCompleted,
		Labels:       task.Labels,
		ParentID:     task.ParentID,
		Order:        task.Order,
		Priority:     task.Priority,
		Due:          task.Due,
		URL:          task.URL,
		CommentCount: task.CommentCount,
		CreatedAt:    task.CreatedAt,
		AssigneeID:   task.AssigneeID,
		AssignerID:   task.AssignerID,
		Duration:     task.Duration,
	}, nil
}

func (t Todoist) CloseTask(id string) error {
	url := fmt.Sprintf("https://api.todoist.com/rest/v2/tasks/%s/close", id)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	return nil
}

func (t Todoist) ReopenTask(id string) error {
	url := fmt.Sprintf("https://api.todoist.com/rest/v2/tasks/%s/reopen", id)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	return nil
}

func (t Todoist) DeleteTask(id string) error {
	url := fmt.Sprintf("https://api.todoist.com/rest/v2/tasks/%s", id)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	return nil
}
