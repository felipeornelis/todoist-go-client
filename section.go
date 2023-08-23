package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const SECTION_URL = BASE_URL + "/sections"

type Section struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Order     uint8  `json:"order"`
	Name      string `json:"name"`
}

func (t Todoist) GetSections(id string) ([]Section, error) {
	url := fmt.Sprintf("%s?project_id=%s", PROJECT_URL, id)

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

	var section []Section
	if err := json.Unmarshal(data, &section); err != nil {
		return nil, err
	}

	return section, nil
}

func (t Todoist) GetSection(id string) (Section, error) {
	if id == "" {
		return Section{}, errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", PROJECT_URL, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Section{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Section{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Section{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Section{}, err
	}

	var section Section
	if err := json.Unmarshal(data, &section); err != nil {
		return Section{}, err
	}

	return section, nil
}

type AddSectionArgs struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	Order     uint   `json:"order,omitempty"`
}

func (t Todoist) AddSection(args AddSectionArgs) (Section, error) {
	requestBody, err := json.Marshal(args)
	if err != nil {
		return Section{}, err
	}

	request, err := http.NewRequest(http.MethodPost, SECTION_URL, bytes.NewReader(requestBody))
	if err != nil {
		return Section{}, err
	}

	request.Header.Set("Auhtorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Section{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Section{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Section{}, err
	}

	var section Section
	if err := json.Unmarshal(data, &section); err != nil {
		return Section{}, err
	}

	return section, nil
}

type UpdateSectionArgs struct {
	Name string `json:"name"`
}

func (t Todoist) UpdateSection(args UpdateSectionArgs, id string) (Section, error) {
	if args.Name == "" {
		return Section{}, errors.New("Name field is required")
	}

	if id == "" {
		return Section{}, errors.New("ID is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Section{}, err
	}

	url := fmt.Sprintf("%s/%s", SECTION_URL, id)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return Section{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Section{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Section{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Section{}, err
	}

	var section Section
	if err := json.Unmarshal(data, &section); err != nil {
		return Section{}, err
	}

	return section, nil
}

func (t Todoist) DeleteSection(id string) error {
	if id == "" {
		return errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", PROJECT_URL, id)

	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

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
