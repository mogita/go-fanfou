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
func (s *BlocksService) IDs() (*UserIDs, error) {
	u := fmt.Sprintf("blocks/ids.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUserIDs := new(UserIDs)
	_, err = s.client.Do(req, newUserIDs)
	if err != nil {
		return nil, err
	}

	return newUserIDs, nil
}

// Users shall get the list of blocked user details
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/blocks.blocking
func (s *BlocksService) Blocking(opt *BlocksOptParams) ([]User, error) {
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
		return nil, err
	}

	newUsers := new([]User)
	_, err = s.client.Do(req, newUsers)
	if err != nil {
		return nil, err
	}

	return *newUsers, nil
}
