package fanfou

import (
	"fmt"
	"net/url"
)

// StatusesService handles communication with the statuses related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#statuses
type StatusesService struct {
	client *Client
}

type Status struct {
	CreatedAt           string `json:"created_at,omitempty"`
	ID                  string `json:"id,omitempty"`
	Rawid               int64  `json:"rawid,omitempty"`
	Text                string `json:"text,omitempty"`
	Source              string `json:"source,omitempty"`
	Truncated           bool   `json:"truncated,omitempty"`
	InReplyToStatusID   string `json:"in_reply_to_status_id,omitempty"`
	InReplyToUserID     string `json:"in_reply_to_user_id,omitempty"`
	InReplyToScreenName string `json:"in_reply_to_screen_name,omitempty"`
	RepostStatusID      string `json:"repost_status_id,omitempty"`
	RepostStatus        string `json:"repost_status,omitempty"`
	RepostUserID        string `json:"repost_user_id,omitempty"`
	RepostScreenName    string `json:"repost_screen_name,omitempty"`
	Favorited           bool   `json:"favorited,omitempty"`
	User                struct {
		ID                        string `json:"id,omitempty"`
		Name                      string `json:"name,omitempty"`
		ScreenName                string `json:"screen_name,omitempty"`
		Location                  string `json:"location,omitempty"`
		Gender                    string `json:"gender,omitempty"`
		Birthday                  string `json:"birthday,omitempty"`
		Description               string `json:"description,omitempty"`
		ProfileImageURL           string `json:"profile_image_url,omitempty"`
		ProfileImageURLLarge      string `json:"profile_image_url_large,omitempty"`
		URL                       string `json:"url,omitempty"`
		Protected                 bool   `json:"protected,omitempty"`
		FollowersCount            int64  `json:"followers_count,omitempty"`
		FriendsCount              int64  `json:"friends_count,omitempty"`
		FavouritesCount           int64  `json:"favourites_count,omitempty"`
		StatusesCount             int64  `json:"statuses_count,omitempty"`
		Following                 bool   `json:"following,omitempty"`
		Notifications             bool   `json:"notifications,omitempty"`
		CreatedAt                 string `json:"created_at,omitempty"`
		UtcOffset                 int64  `json:"utc_offset,omitempty"`
		ProfileBackgroundColor    string `json:"profile_background_color,omitempty"`
		ProfileTextColor          string `json:"profile_text_color,omitempty"`
		ProfileLinkColor          string `json:"profile_link_color,omitempty"`
		ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color,omitempty"`
		ProfileSidebarBorderColor string `json:"profile_sidebar_border_color,omitempty"`
		ProfileBackgroundImageURL string `json:"profile_background_image_url,omitempty"`
		ProfileBackgroundTile     bool   `json:"profile_background_tile,omitempty"`
	} `json:"user,omitempty"`
	Photo struct {
		Imageurl string `json:"imageurl,omitempty"`
		Thumburl string `json:"thumburl,omitempty"`
		Largeurl string `json:"largeurl,omitempty"`
	} `json:"photo,omitempty"`
}

type StatusesOptParams struct {
	InReplyToStasusID string
	InReplyToUserID   string
	RepostStatusID    string
	Source            string
	Mode              string
	Format            string
	Location          string
}

// Post a new status
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/statuses.update
func (s *StatusesService) Update(status string, opt *StatusesOptParams) (*Status, error) {
	u := fmt.Sprintf("statuses/update.json")
	params := url.Values{
		"status": []string{status},
	}

	if opt != nil {
		if opt.InReplyToStasusID != "" {
			params.Add("in_reply_to_status_id", opt.InReplyToStasusID)
		}
		if opt.InReplyToUserID != "" {
			params.Add("in_reply_to_user_id", opt.InReplyToUserID)
		}
		if opt.RepostStatusID != "" {
			params.Add("repost_status_id", opt.RepostStatusID)
		}
		if opt.Source != "" {
			params.Add("source", opt.Source)
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
		if opt.Location != "" {
			params.Add("location", opt.Location)
		}
		u += "?" + params.Encode()
		fmt.Printf("%+v\n", u)
	}

	req, err := s.client.NewRequest("POST", u, params.Encode())
	if err != nil {
		return nil, err
	}

	newStatus := new(Status)
	_, err = s.client.Do(req, newStatus)
	if err != nil {
		return nil, err
	}

	return newStatus, nil
}
