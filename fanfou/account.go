package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// AccountService handles communication with the account related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#account
type AccountService struct {
	client *Client
}

// RateLimitStatusResult is the structure of rate limit
type RateLimitStatusResult struct {
	ResetTime          string `json:"reset_time,omitempty"`
	RemainingHits      int64  `json:"remaining_hits,omitempty"`
	HourlyLimit        int64  `json:"hourly_limit,omitempty"`
	ResetTimeInSeconds int64  `json:"reset_time_in_seconds,omitempty"`
}

// NotificationResult is the structure of notification
type NotificationResult struct {
	Mentions       int64 `json:"mentions,omitempty"`
	DirectMessages int64 `json:"direct_messages,omitempty"`
	FriendRequests int64 `json:"friend_requests,omitempty"`
}

// NotifyNumResult is the structure of notify number
type NotifyNumResult struct {
	Result    string `json:"result,omitempty"`
	NotifyNum int64  `json:"notify_num,omitempty"`
}

// AccountOptParams specifies the optional params for account API
type AccountOptParams struct {
	URL         string
	Location    string
	Description string
	Name        string
	Email       string
	Mode        string
	Format      string
	NotifyNum   int64
}

// VerifyCredentials shall verify the current user's username and password
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.verify-credentials
func (s *AccountService) VerifyCredentials(opt *AccountOptParams) (*UserResult, *string, error) {
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
		return nil, nil, err
	}

	newUser := new(UserResult)
	resp, err := s.client.Do(req, newUser)
	if err != nil {
		return nil, nil, err
	}

	return newUser, resp.BodyStrPtr, nil
}

// RateLimitStatus shall get the API rate limit information of the current user
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.rate-limit-status
func (s *AccountService) RateLimitStatus() (*RateLimitStatusResult, *string, error) {
	u := fmt.Sprintf("account/rate_limit_status.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newRateLimitStatus := new(RateLimitStatusResult)
	resp, err := s.client.Do(req, newRateLimitStatus)
	if err != nil {
		return nil, nil, err
	}

	return newRateLimitStatus, resp.BodyStrPtr, nil
}

// UpdateProfile shall update the current user's profile
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.update-profile
func (s *AccountService) UpdateProfile(opt *AccountOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("account/update_profile.json")
	params := url.Values{}

	if opt != nil {
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.URL != "" {
			params.Add("url", opt.URL)
		}
		if opt.Location != "" {
			params.Add("location", opt.Location)
		}
		if opt.Description != "" {
			params.Add("description", opt.Description)
		}
		if opt.Name != "" {
			params.Add("name", opt.Name)
		}
		if opt.Email != "" {
			params.Add("email", opt.Email)
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

// UpdateProfileImage shall update the current user's profile image
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.update-profile-image
func (s *AccountService) UpdateProfileImage(filePath string, opt *AccountOptParams) (*UserResult, *string, error) {
	u := fmt.Sprintf("account/update_profile_image.json")
	params := map[string]string{}

	if opt != nil {
		if opt.Mode != "" {
			params["mode"] = opt.Mode
		}
		if opt.Format != "" {
			params["format"] = opt.Format
		}
	}

	req, err := s.client.NewUploadRequest(http.MethodPost, u, params, "image", filePath)
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

// Notification shall get the unread counts for mentions, direct
// messages and friend requests of the current user
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.notification
func (s *AccountService) Notification() (*NotificationResult, *string, error) {
	u := fmt.Sprintf("account/notification.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newNotification := new(NotificationResult)
	resp, err := s.client.Do(req, newNotification)
	if err != nil {
		return nil, nil, err
	}

	return newNotification, resp.BodyStrPtr, nil
}

// NotifyNum shall get new notification number of the current app
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.notify-num
func (s *AccountService) NotifyNum() (*NotifyNumResult, *string, error) {
	u := fmt.Sprintf("account/notify_num.json")

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newNotifyNum := new(NotifyNumResult)
	resp, err := s.client.Do(req, newNotifyNum)
	if err != nil {
		return nil, nil, err
	}

	return newNotifyNum, resp.BodyStrPtr, nil
}

// UpdateNotifyNum shall update the notification number of the current app
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/account.update-notify-num
func (s *AccountService) UpdateNotifyNum(opt *AccountOptParams) (*NotifyNumResult, *string, error) {
	u := fmt.Sprintf("account/update_notify_num.json")
	params := url.Values{}

	if opt != nil {
		if opt.NotifyNum != 0 {
			params.Add("notify_num", strconv.FormatInt(opt.NotifyNum, 10))
		}
	}

	req, err := s.client.NewRequest(http.MethodPost, u, params.Encode())
	if err != nil {
		return nil, nil, err
	}

	newNotifyNum := new(NotifyNumResult)
	resp, err := s.client.Do(req, newNotifyNum)
	if err != nil {
		return nil, nil, err
	}

	return newNotifyNum, resp.BodyStrPtr, nil
}
