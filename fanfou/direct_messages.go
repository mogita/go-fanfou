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

// DirectMessageConversationListResult specifies Fanfou's direct messages conversation list data structure
type DirectMessageConversationListResult []DirectMessageConversationListItem

// DirectMessageConversationListItem specifies Fanfou's direct messages conversation list item data structure
type DirectMessageConversationListItem struct {
	Dm      *DirectMessageResult `json:"dm"`
	Otherid string               `json:"otherid"`
	MsgNum  int64                `json:"msg_num"`
	NewConv bool                 `json:"new_conv"`
}

// DirectMessagesOptParams specifies the optional params for direct messages API
type DirectMessagesOptParams struct {
	ID          string
	InReplyToID string
	Page        int64
	Count       int64
	MaxID       string
	SinceID     string
	Mode        string
	Format      string
}

// Conversation shall get the conversation of direct messages between the specified
// user and the current user
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/direct-messages.conversation
func (s *DirectMessagesService) Conversation(ID string, opt *DirectMessagesOptParams) ([]DirectMessageResult, *string, error) {
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
		return nil, nil, err
	}

	newDirectMessages := new([]DirectMessageResult)
	resp, err := s.client.Do(req, newDirectMessages)
	if err != nil {
		return nil, nil, err
	}

	return *newDirectMessages, resp.BodyStrPtr, nil
}

// New shall send a new direct message to the specified user
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/direct-messages.new
func (s *DirectMessagesService) New(user, text string, opt *DirectMessagesOptParams) (*DirectMessageResult, *string, error) {
	u := fmt.Sprintf("direct_messages/new.json")
	params := url.Values{
		"user": []string{user},
		"text": []string{text},
	}

	if opt != nil {
		if opt.InReplyToID != "" {
			params.Add("in_reply_to_id", opt.InReplyToID)
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, nil, err
	}

	newDirectMessage := new(DirectMessageResult)
	resp, err := s.client.Do(req, newDirectMessage)
	if err != nil {
		return nil, nil, err
	}

	return newDirectMessage, resp.BodyStrPtr, nil
}

// Destroy shall delete a specified direct message
// ID represents the direct message ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/direct-messages.destroy
func (s *DirectMessagesService) Destroy(ID string) (*DirectMessageResult, *string, error) {
	u := fmt.Sprintf("direct_messages/destroy.json")
	params := url.Values{
		"id": []string{ID},
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodPost, u, "")
	if err != nil {
		return nil, nil, err
	}

	newDirectMessage := new(DirectMessageResult)
	resp, err := s.client.Do(req, newDirectMessage)
	if err != nil {
		return nil, nil, err
	}

	return newDirectMessage, resp.BodyStrPtr, nil
}

// ConversationList shall get the conversation list of the direct messages of
// the current user
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/direct-messages.conversation-list
func (s *DirectMessagesService) ConversationList(opt *DirectMessagesOptParams) (*DirectMessageConversationListResult, *string, error) {
	u := fmt.Sprintf("direct_messages/conversation_list.json")
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

	newDirectMessages := new(DirectMessageConversationListResult)
	resp, err := s.client.Do(req, newDirectMessages)
	if err != nil {
		return nil, nil, err
	}

	return newDirectMessages, resp.BodyStrPtr, nil
}

// Inbox shall get the latest direct messages in the inbox
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/direct-messages.inbox
func (s *DirectMessagesService) Inbox(opt *DirectMessagesOptParams) ([]DirectMessageResult, *string, error) {
	u := fmt.Sprintf("direct_messages/inbox.json")
	params := url.Values{}

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
		return nil, nil, err
	}

	newDirectMessages := new([]DirectMessageResult)
	resp, err := s.client.Do(req, newDirectMessages)
	if err != nil {
		return nil, nil, err
	}

	return *newDirectMessages, resp.BodyStrPtr, nil
}

// Sent shall get the sent direct messages by the current user
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/direct-messages.sent
func (s *DirectMessagesService) Sent(opt *DirectMessagesOptParams) ([]DirectMessageResult, *string, error) {
	u := fmt.Sprintf("direct_messages/sent.json")
	params := url.Values{}

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
		return nil, nil, err
	}

	newDirectMessages := new([]DirectMessageResult)
	resp, err := s.client.Do(req, newDirectMessages)
	if err != nil {
		return nil, nil, err
	}

	return *newDirectMessages, resp.BodyStrPtr, nil
}
