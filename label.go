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

const LABEL_URL = BASE_URL + "/labels"

type Label struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Order      uint8  `json:"order"`
	IsFavorite bool   `json:"is_favorite"`
}

func (t Todoist) GetPersonalLabels() ([]Label, error) {
	request, err := http.NewRequest(http.MethodGet, LABEL_URL, nil)
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

	var labels []Label
	if err := json.Unmarshal(data, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

func (t Todoist) GetPersonalLabel(id string) (Label, error) {
	if id == "" {
		return Label{}, errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", LABEL_URL, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Label{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Label{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Label{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Label{}, err
	}

	var label Label
	if err := json.Unmarshal(data, &label); err != nil {
		return Label{}, err
	}

	return label, nil
}

type AddPersonalLabelArgs struct {
	Name       string `json:"name"`
	Order      uint8  `json:"order,omitempty"`
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
}

func (t Todoist) AddPersonalLabel(args AddPersonalLabelArgs) (Label, error) {
	if args.Name == "" {
		return Label{}, errors.New("`name` is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Label{}, err
	}

	request, err := http.NewRequest(http.MethodPost, LABEL_URL, bytes.NewReader(bodyRequest))
	if err != nil {
		return Label{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Label{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Label{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Label{}, err
	}

	var label Label
	if err := json.Unmarshal(data, &label); err != nil {
		return Label{}, err
	}

	return label, nil
}

type UpdatePersonalLabelArgs struct {
	Name       string `json:"name,omitempty"`
	Order      uint8  `json:"order,omitempty"`
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
}

func (t Todoist) UpdatePersonalLabel(id string, args UpdatePersonalLabelArgs) (Label, error) {
	if id == "" {
		return Label{}, errors.New("ID is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return Label{}, err
	}

	url := fmt.Sprintf("%s/%s", LABEL_URL, id)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return Label{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))

	client := &http.Client{
		Timeout: MAX_TIMEOUT,
	}

	response, err := client.Do(request)
	if err != nil {
		return Label{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Label{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Label{}, err
	}

	var label Label
	if err := json.Unmarshal(data, &label); err != nil {
		return Label{}, err
	}

	return label, nil
}

func (t Todoist) DeleteLabel(id string) error {
	if id == "" {
		return errors.New("ID is required")
	}

	url := fmt.Sprintf("%s/%s", LABEL_URL, id)

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

type GetSharedLabelsArgs struct {
	OmitPersonal bool `json:"omit_personal,omitempty"`
}

func (t Todoist) GetSharedLabels(args GetSharedLabelsArgs) ([]string, error) {
	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/shared", LABEL_URL)

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("Content-Type", "application/json")

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

	var labels []string
	if err := json.Unmarshal(data, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

type RenameSharedLabelsArgs struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

func (t Todoist) RenameSharedLabels(args RenameSharedLabelsArgs) error {
	if args.Name == "" || args.NewName == "" {
		return errors.New("`name` and `new_name` are required")
	}

	// This validation is not necessaraly required or supported by Todoist's REST API
	if args.Name == args.NewName {
		return errors.New("If you really mean to rename the label, previous and new names need to be different")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/shared/rename", LABEL_URL)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Content-Type", "application/json")

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

type RemoveSharedLabelsArgs struct {
	Name string `json:"name"`
}

func (t Todoist) RemoveSharedLabels(args RemoveSharedLabelsArgs) error {
	if args.Name == "" {
		return errors.New("`name` is required")
	}

	bodyRequest, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/shared/remove", LABEL_URL)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyRequest))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.authToken))
	request.Header.Set("X-Request-Id", pkg.NewUUID())
	request.Header.Set("Content-Type", "application/json")

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
