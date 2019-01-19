package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// FriendsService handles communication with the friends related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#friends
type FriendsService struct {
	client *Client
}

// FriendsOptParams specifies the optional params for friends API
type FriendsOptParams struct {
	ID    string
	Page  int64
	Count int64
}

// IDs shall get friends IDs of the specified user, or of the current user
// if no ID specified
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friends.ids
func (s *FriendsService) IDs(opt *FriendsOptParams) (*UserIDs, *string, error) {
	u := fmt.Sprintf("friends/ids.json")
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
	}

	u += "?" + params.Encode()

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
