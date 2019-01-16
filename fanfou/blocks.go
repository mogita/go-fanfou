package fanfou

import (
	"fmt"
	"net/http"
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
