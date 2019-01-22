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

// FriendshipsShowResult specifies Fanfou's friendship show structure
type FriendshipsShowResult struct {
	Relationship *RelationshipResult `json:"relationship,omitempty"`
}

type RelationshipResult struct {
	Source *RelationshipItem `json:"source,omitempty"`
	Target *RelationshipItem `json:"target,omitempty"`
}

type RelationshipItem struct {
	ID                   string `json:"id,omitempty"`
	ScreenName           string `json:"screen_name,omitempty"`
	Following            string `json:"following,omitempty"`
	FollowedBy           string `json:"followed_by,omitempty"`
	NotificationsEnabled string `json:"notifications_enabled,omitempty"`
	Blocking             string `json:"blocking,omitempty"`
}

// FriendshipsOptParams specifies the optional params for friendships API
type FriendshipsOptParams struct {
	ID     string
	Page   int64
	Count  int64
	Mode   string
	Format string
}

// FriendshipsShowOptParams specifies the params specifically for friendship
// show API
type FriendshipsShowOptParams struct {
	SourceLoginName string
	TargetLoginName string
	SourceID        string
	TargetID        string
}

// Create shall add the specified user as a friend (follow)
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.create
func (s *FriendshipsService) Create(ID string, opt *FriendshipsOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("friendships/create.json")
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

// Destroy shall unfriend the specified user (unfollow)
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.destroy
func (s *FriendshipsService) Destroy(ID string, opt *FriendshipsOptParams) (*UserResult, *string, error) {
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

// Requests shall get the list of friendship requests (other users'
// requests to follow the current user)
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.requests
func (s *FriendshipsService) Requests(opt *FriendshipsOptParams) ([]UserResult, *string, error) {
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
		return nil, nil, err
	}

	newUsers := new([]UserResult)
	resp, err := s.client.Do(req, newUsers)
	if err != nil {
		return nil, nil, err
	}

	return *newUsers, resp.BodyStrPtr, nil
}

// Deny shall reject the friendship request from the specified user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.deny
func (s *FriendshipsService) Deny(ID string, opt *FriendshipsOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("friendships/deny.json")
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

// Accept shall accept the friendship request from the specified user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.accept
func (s *FriendshipsService) Accept(ID string, opt *FriendshipsOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("friendships/accept.json")
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

// Exists shall check if user A follows user B
// "userA" and "userB" can be either the user ID or login name
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.exists
func (s *FriendshipsService) Exists(userA, userB string) (bool, *string, error) {
	u := fmt.Sprintf("friendships/exists.json")
	params := url.Values{
		"user_a": []string{userA},
		"user_b": []string{userB},
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return false, nil, err
	}

	// Caveat: Fanfou API returns "true" and "false" in string on this endpoint
	result := new(bool)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return false, nil, err
	}

	return *result, resp.BodyStrPtr, nil
}

// Show shall get the friendship between two users
// Note: Please either supply both a source and a target
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.show
func (s *FriendshipsService) Show(opt *FriendshipsShowOptParams) (*FriendshipsShowResult, *string, error) {
	u := fmt.Sprintf("friendships/show.json")
	params := url.Values{}

	if opt != nil {
		if opt.SourceLoginName != "" {
			params.Add("source_login_name", opt.SourceLoginName)
		}
		if opt.SourceID != "" {
			params.Add("source_id", opt.SourceID)
		}
		if opt.TargetLoginName != "" {
			params.Add("target_login_name", opt.TargetLoginName)
		}
		if opt.TargetID != "" {
			params.Add("target_id", opt.TargetID)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	result := new(FriendshipsShowResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp.BodyStrPtr, nil
}
