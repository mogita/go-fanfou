package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// FavoritesService handles communication with the followers related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#followers
type FavoritesService struct {
	client *Client
}

// FavoritesOptParams specifies the optional params for favorites API
type FavoritesOptParams struct {
	ID     string
	Page   int64
	Count  int64
	Mode   string
	Format string
}

// IDs shall get favorites of the specified user, or of the current user
// if no ID specified
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/favorites
func (s *FavoritesService) IDs(opt *FavoritesOptParams) ([]StatusResult, *string, error) {
	u := fmt.Sprintf("favorites/id.json")
	params := url.Values{}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newStatuses := new([]StatusResult)
	resp, err := s.client.Do(req, newStatuses)
	if err != nil {
		return nil, nil, err
	}

	return *newStatuses, resp.BodyStrPtr, nil
}

// Create shall create a favorite
// ID represents the status ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/favorites.create
func (s *FavoritesService) Create(ID string, opt *FavoritesOptParams) (*StatusResult, *string, error) {
	u := fmt.Sprintf("favorites/create.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, nil, err
	}

	newStatus := new(StatusResult)
	resp, err := s.client.Do(req, newStatus)
	if err != nil {
		return nil, nil, err
	}

	return newStatus, resp.BodyStrPtr, nil
}

// Destroy shall delete a favorite
// ID represents the status ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/favorites.destroy
func (s *FavoritesService) Destroy(ID string, opt *FavoritesOptParams) (*StatusResult, *string, error) {
	u := fmt.Sprintf("favorites/destroy.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, nil, err
	}

	newStatus := new(StatusResult)
	resp, err := s.client.Do(req, newStatus)
	if err != nil {
		return nil, nil, err
	}

	return newStatus, resp.BodyStrPtr, nil
}
