package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// DirectMessagesService handles communication with the direct messages related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#direct-messages
type DirectMessagesService struct {
	client *Client
}

// DirectMessageResult specifies Fanfou's direct messages structure
type DirectMessageResult struct {
	ID                  string               `json:"id,omitempty"`
	Text                string               `json:"text,omitempty"`
	SenderID            string               `json:"sender_id,omitempty"`
	RecipientID         string               `json:"recipient_id,omitempty"`
	CreatedAt           string               `json:"created_at,omitempty"`
	SenderScreenName    string               `json:"sender_screen_name,omitempty"`
	RecipientScreenName string               `json:"recipient_screen_name,omitempty"`
	Sender              *UserResult          `json:"sender,omitempty"`
	Recipient           *UserResult          `json:"recipient,omitempty"`
	InReplyTo           *DirectMessageResult `json:"in_reply_to,omitempty"`
}

// DirectMessagesOptParams specifies the optional params for direct messages API
type DirectMessagesOptParams struct {
	ID      string
	Page    int64
	Count   int64
	MaxID   string
	SinceID string
	Mode    string
	Format  string
}

// Conversation shall get the conversation of direct messages between the specified
// user and the current user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/friendships.create
func (s *DirectMessagesService) Conversation(ID string, opt *DirectMessagesOptParams) ([]DirectMessageResult, error) {
	u := fmt.Sprintf("direct_messages/conversation.json")
	params := url.Values{
		"id": []string{ID},
	}

	if opt != nil {
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.MaxID != "" {
			params.Add("max_id", opt.MaxID)
		}
		if opt.SinceID != "" {
			params.Add("since_id", opt.SinceID)
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

	newDirectMessages := new([]DirectMessageResult)
	_, err = s.client.Do(req, newDirectMessages)
	if err != nil {
		return nil, err
	}

	return *newDirectMessages, nil
}
