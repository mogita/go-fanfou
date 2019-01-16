package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
)

// AccountService handles communication with the account related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#account
type AccountService struct {
	client *Client
}

// AccountOptParams specifies the optional params for account API
type AccountOptParams struct {
	ID      string
	SinceID string
	MaxID   string
	Page    int64
	Count   int64
	Mode    string
	Format  string
}

// VerifyCredentials shall verify the current user's username and password
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.verify-credentials
func (s *AccountService) VerifyCredentials(opt *AccountOptParams) (*User, error) {
	u := fmt.Sprintf("account/verify_credentials.json")
	params := url.Values{}

	if opt != nil {
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

	newUser := new(User)
	_, err = s.client.Do(req, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
