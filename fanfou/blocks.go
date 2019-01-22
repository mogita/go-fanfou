package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// BlocksService handles communication with the blocks related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#blocks
type BlocksService struct {
	client *Client
}

type UserIDs []string

// BlocksOptParams specifies the optional params for blocks API
type BlocksOptParams struct {
	ID      string
	SinceID string
	MaxID   string
	Page    int64
	Count   int64
	Mode    string
	Format  string
}

// IDs shall get the list of blocked user IDs
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/blocks.ids
func (s *BlocksService) IDs() (*UserIDs, *string, error) {
	u := fmt.Sprintf("blocks/ids.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newUserIDs := new(UserIDs)
	resp, err := s.client.Do(req, newUserIDs)
	if err != nil {
		return nil, nil, err
	}

	return newUserIDs, resp.BodyStrPtr, nil
}

// Blocking shall get the list of blocked user details
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/blocks.blocking
func (s *BlocksService) Blocking(opt *BlocksOptParams) ([]UserResult, *string, error) {
	u := fmt.Sprintf("blocks/blocking.json")
	params := url.Values{}

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
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newUsers := new([]UserResult)
	resp, err := s.client.Do(req, newUsers)
	if err != nil {
		return nil, nil, err
	}

	return *newUsers, resp.BodyStrPtr, nil
}

// Exists shall check whether the specified user is blocked by the
// current user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/blocks.exists
func (s *BlocksService) Exists(ID string, opt *BlocksOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("blocks/exists.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newUser := new(UserResult)
	resp, err := s.client.Do(req, newUser)
	if err != nil {
		return nil, nil, err
	}

	return newUser, resp.BodyStrPtr, nil
}

// Create shall block a specified user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/blocks.create
func (s *BlocksService) Create(ID string, opt *BlocksOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("blocks/create.json")
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

	req, err := s.client.NewRequest(http.MethodPost, u, params.Encode())
	if err != nil {
		return nil, nil, err
	}

	newUser := new(UserResult)
	resp, err := s.client.Do(req, newUser)
	if err != nil {
		return nil, nil, err
	}

	return newUser, resp.BodyStrPtr, nil
}

// Destroy shall unblock a specified user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/blocks.destroy
func (s *BlocksService) Destroy(ID string, opt *BlocksOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("blocks/destroy.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
	}

	req, err := s.client.NewRequest(http.MethodPost, u, params.Encode())
	if err != nil {
		return nil, nil, err
	}

	newUser := new(UserResult)
	resp, err := s.client.Do(req, newUser)
	if err != nil {
		return nil, nil, err
	}

	return newUser, resp.BodyStrPtr, nil
}
