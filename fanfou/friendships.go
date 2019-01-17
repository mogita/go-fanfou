package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// FriendshipsService handles communication with the friendships related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#friendships
type FriendshipsService struct {
	client *Client
}

// FriendsOptParams specifies the optional params for friendships API
type FriendshipsOptParams struct {
	ID     string
	Page   int64
	Count  int64
	Mode   string
	Format string
}

// Create shall add the specified user as a friend (follow)
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.create
func (s *FriendshipsService) Create(ID string, opt *FriendshipsOptParams) (*UserResult, error) {
	u := fmt.Sprintf("friendships/create.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, err
	}

	newUser := new(UserResult)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Destroy shall unfriend the specified user (unfollow)
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.destroy
func (s *FriendshipsService) Destroy(ID string, opt *FriendshipsOptParams) (*UserResult, error) {
	u := fmt.Sprintf("friendships/destroy.json")
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
		return nil, err
	}

	newUser := new(UserResult)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Requests shall get the list of friendship requests (other users'
// requests to follow the current user)
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.requests
func (s *FriendshipsService) Requests(opt *FriendshipsOptParams) ([]UserResult, error) {
	u := fmt.Sprintf("friendships/requests.json")
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
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, err
	}

	newUsers := new([]UserResult)
	_, err = s.client.Do(req, newUsers)
	if err != nil {
		return nil, err
	}

	return *newUsers, nil
}
