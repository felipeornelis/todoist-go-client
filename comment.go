package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/felipeornelis/todoist-go-client/pkg"
)

const COMMENT_URL = BASE_URL + "/comments"

type Comment struct {
	ID         string            `json:"id"`
	TaskID     string            `json:"task_id"`
	ProjectID  string            `json:"project_id"`
	PostedAt   string            `json:"posted_at"`
	Content    string            `json:"content"`
	Attachment CommentAttachment `json:"attachment,omitempty"`
}

type CommentAttachment struct {
	FileName     string `json:"file_name,omitempty"`
	FileType     string `json:"file_type,omitempty"`
	FileURL      string `json:"file_url,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

type GetCommentsArgs struct {
	TaskID    string `json:"task_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

func (t Todoist) GetComments(args GetCommentsArgs) ([]Comment, error) {
	if args.ProjectID == "" && args.TaskID == "" {
		return nil, errors.New("task_id or project_id is required")
	}

	var url string

	if args.ProjectID == "" {
		url = fmt.Sprintf("%s?task_id=%s", COMMENT_URL, args.TaskID)
	} else {
		url = fmt.Sprintf("%s?project_id=%s", COMMENT_URL, args.ProjectID)
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	if err := json.Unmarshal(data, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (t Todoist) GetComment(id string) (Comment, error) {
	if id == "" {
		return Comment{}, errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", COMMENT_URL, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Comment{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Comment{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Comment{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Comment{}, err
	}

	var comment Comment
	if err := json.Unmarshal(data, &comment); err != nil {
		return Comment{}, err
	}

	return comment, nil
}

type AddCommentArgs struct {
	TaskID     string            `json:"task_id,omitempty"`
	ProjectID  string            `json:"project_id,omitempty"`
	Content    string            `json:"content"`
	Attachment CommentAttachment `json:"attachment,omitempty"`
}

func (t Todoist) AddComment(args AddCommentArgs) (Comment, error) {
	if args.TaskID == "" && args.ProjectID == "" {
		return Comment{}, errors.New("task_id or project_id is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Comment{}, err
	}

	request, err := http.NewRequest(http.MethodPost, COMMENT_URL, bytes.NewReader(bodyRequest))
	if err != nil {
		return Comment{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Comment{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Comment{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Comment{}, err
	}

	var comment Comment
	if err := json.Unmarshal(data, &comment); err != nil {
		return Comment{}, err
	}

	return comment, nil
}

type UpdateCommentArgs struct {
	Content string `json:"content"`
}

func (t Todoist) UpdateComment(id string, args UpdateCommentArgs) (Comment, error) {
	if id == "" {
		return Comment{}, errors.New("ID is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Comment{}, err
	}

	url := fmt.Sprintf("%s/%s", COMMENT_URL, id)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return Comment{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Comment{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Comment{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Comment{}, err
	}

	var comment Comment
	if err := json.Unmarshal(data, &comment); err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func (t Todoist) DeleteComment(id string) error {
	url := fmt.Sprintf("%s/%s", COMMENT_URL, id)

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
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
