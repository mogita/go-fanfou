package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// SearchService handles communication with the search related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#search
type SearchService struct {
	client *Client
}

type SearchUsersResult struct {
	TotalNumber int64        `json:"total_number,omitempty"`
	Users       []UserResult `json:"users,omitempty"`
}

// SearchOptParams specifies the optional params for search API
type SearchOptParams struct {
	ID      string
	SinceID string
	MaxID   string
	Page    int64
	Count   int64
	Mode    string
	Format  string
}

// PublicTimeline shall search for statuses of the whole platform
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/search.public-timeline
func (s *SearchService) PublicTimeline(q string, opt *SearchOptParams) ([]Status, error) {
	u := fmt.Sprintf("search/public_timeline.json")
	params := url.Values{
		"q": []string{q},
	}

	if opt != nil {
		if opt.SinceID != "" {
			params.Add("since_id", opt.SinceID)
		}
		if opt.MaxID != "" {
			params.Add("max_id", opt.MaxID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
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
		return nil, err
	}

	newStatuses := new([]Status)
	_, err = s.client.Do(req, newStatuses)
	if err != nil {
		return nil, err
	}

	return *newStatuses, nil
}

// UserTimeline shall search for statuses of the specified user, or of the current user
// if no ID specified
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/search.user-timeline
func (s *SearchService) UserTimeline(q string, opt *SearchOptParams) ([]Status, error) {
	u := fmt.Sprintf("search/user_timeline.json")
	params := url.Values{
		"q": []string{q},
	}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
		if opt.SinceID != "" {
			params.Add("since_id", opt.SinceID)
		}
		if opt.MaxID != "" {
			params.Add("max_id", opt.MaxID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
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
		return nil, err
	}

	newStatuses := new([]Status)
	_, err = s.client.Do(req, newStatuses)
	if err != nil {
		return nil, err
	}

	return *newStatuses, nil
}

// Users shall search for users of the whole platform
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/search.users
func (s *SearchService) Users(q string, opt *SearchOptParams) (*SearchUsersResult, error) {
	u := fmt.Sprintf("search/users.json")
	params := url.Values{
		"q": []string{q},
	}

	if opt != nil {
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
		return nil, err
	}

	newSearchUsersResult := new(SearchUsersResult)
	_, err = s.client.Do(req, newSearchUsersResult)
	if err != nil {
		return nil, err
	}

	return newSearchUsersResult, nil
}
