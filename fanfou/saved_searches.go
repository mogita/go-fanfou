package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
)

// SavedSearchesService handles communication with the saved searches related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#saved-searches
type SavedSearchesService struct {
	client *Client
}

// SavedSearchResult is the structure of saved search
type SavedSearchResult struct {
	ID        int64  `json:"id,omitempty"`
	Query     string `json:"query,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// Show shall get a saved searches detail
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.show
func (s *SavedSearchesService) Show(ID string) (*SavedSearchResult, *string, error) {
	u := fmt.Sprintf("saved_searches/show.json")
	params := url.Values{
		"id": []string{ID},
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newSavedSearch := new(SavedSearchResult)
	resp, err := s.client.Do(req, newSavedSearch)
	if err != nil {
		return nil, nil, err
	}

	return newSavedSearch, resp.BodyStrPtr, nil
}

// List shall get the list of the current user's saved searches
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.list
func (s *SavedSearchesService) List() ([]SavedSearchResult, *string, error) {
	u := fmt.Sprintf("saved_searches/list.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newSavedSearches := new([]SavedSearchResult)
	resp, err := s.client.Do(req, newSavedSearches)
	if err != nil {
		return nil, nil, err
	}

	return *newSavedSearches, resp.BodyStrPtr, nil
}

// Create shall create a saved search
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.create
func (s *SavedSearchesService) Create(query string) (*SavedSearchResult, *string, error) {
	u := fmt.Sprintf("saved_searches/create.json")
	params := url.Values{
		"query": []string{query},
	}

	req, err := s.client.NewRequest(http.MethodPost, u, params.Encode())
	if err != nil {
		return nil, nil, err
	}

	newSavedSearch := new(SavedSearchResult)
	resp, err := s.client.Do(req, newSavedSearch)
	if err != nil {
		return nil, nil, err
	}

	return newSavedSearch, resp.BodyStrPtr, nil
}

// Destroy shall delete a saved search
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.destroy
func (s *SavedSearchesService) Destroy(ID string) (*SavedSearchResult, *string, error) {
	u := fmt.Sprintf("saved_searches/destroy.json")
	params := url.Values{
		"id": []string{ID},
	}

	req, err := s.client.NewRequest(http.MethodPost, u, params.Encode())
	if err != nil {
		return nil, nil, err
	}

	newSavedSearch := new(SavedSearchResult)
	resp, err := s.client.Do(req, newSavedSearch)
	if err != nil {
		return nil, nil, err
	}

	return newSavedSearch, resp.BodyStrPtr, nil
}
