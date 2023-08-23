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

const PROJECT_URL = BASE_URL + "/projects"

type Project struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Color          string `json:"color"`
	ParentID       string `json:"parent_id"`
	Order          uint8  `json:"order"`
	CommentCount   int    `json:"comment_count"`
	IsShared       bool   `json:"is_shared"`
	IsFavorite     bool   `json:"is_favorite"`
	IsInboxProject bool   `json:"is_inbox_project"`
	IsTeamInbox    bool   `json:"is_team_inbox"`
	ViewStyle      string `json:"view_style"`
	URL            string `json:"url"`
}

func (t Todoist) GetProjects() ([]Project, error) {
	request, err := http.NewRequest(http.MethodGet, PROJECT_URL, nil)
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

	var projects []Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (t Todoist) GetProject(id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", PROJECT_URL, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Project{}, err
	}

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	response, err := client.Do(request)
	if err != nil {
		return Project{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Project{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Project{}, err
	}

	var project Project
	if err := json.Unmarshal(data, &project); err != nil {
		return Project{}, err
	}

	return project, nil
}

type UpdateProjectArgs struct {
	Name       string `json:"name,omitempty"`
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
	ViewStyle  string `json:"view_style,omitempty"`
}

func (t Todoist) UpdateProject(args UpdateProjectArgs, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", PROJECT_URL, id)

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Project{}, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return Project{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Request-Id", pkg.NewUUID())

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Project{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Project{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Project{}, err
	}

	var project Project
	if err := json.Unmarshal(data, &project); err != nil {
		return Project{}, err
	}

	return project, nil
}

func (t Todoist) DeleteProject(id string) error {
	if id == "" {
		return errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", PROJECT_URL, id)

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

type GetAllCollaboratorsOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (t Todoist) GetAllCollaborators(id string) ([]GetAllCollaboratorsOutput, error) {
	if id == "" {
		return nil, errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s/collaborators", PROJECT_URL, id)

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

	var collaborators []GetAllCollaboratorsOutput

	if err := json.Unmarshal(data, &collaborators); err != nil {
		return nil, err
	}

	return collaborators, nil
}
