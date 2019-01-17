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

type SavedSearchResult struct {
	ID        int64  `json:"id,omitempty"`
	Query     string `json:"query,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// Show shall get a saved searches detail
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.show
func (s *SavedSearchesService) Show(ID string) (*SavedSearchResult, error) {
	u := fmt.Sprintf("saved_searches/show.json")
	params := url.Values{
		"id": []string{ID},
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newSavedSearch := new(SavedSearchResult)
	_, err = s.client.Do(req, newSavedSearch)
	if err != nil {
		return nil, err
	}

	return newSavedSearch, nil
}

// List shall get the list of the current user's saved searches
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.list
func (s *SavedSearchesService) List() ([]SavedSearchResult, error) {
	u := fmt.Sprintf("saved_searches/list.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newSavedSearches := new([]SavedSearchResult)
	_, err = s.client.Do(req, newSavedSearches)
	if err != nil {
		return nil, err
	}

	return *newSavedSearches, nil
}

// Create shall create a saved search
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/saved-searches.create
func (s *SavedSearchesService) Create(query string) (*SavedSearchResult, error) {
	u := fmt.Sprintf("saved_searches/create.json")
	params := url.Values{
		"query": []string{query},
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, err
	}

	newSavedSearch := new(SavedSearchResult)
	_, err = s.client.Do(req, newSavedSearch)
	if err != nil {
		return nil, err
	}

	return newSavedSearch, nil
}
