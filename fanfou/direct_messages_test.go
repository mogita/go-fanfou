package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDirectMessagesService_Conversation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/direct_messages/conversation.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}, "in_reply_to": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}}, {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}, "in_reply_to": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}}]`)
		if err != nil {
			t.Errorf("direct_messages.conversation mock server error: %+v", err)
		}
	})

	user, err := client.DirectMessages.Conversation("test_id", &DirectMessagesOptParams{
		Page:    1,
		Count:   1,
		MaxID:   "test_max_id",
		SinceID: "test_since_id",
		Mode:    "test_mode",
		Format:  "test_format",
	})
	if err != nil {
		t.Errorf("direct_messages.conversation returned error: %v", err)
	}

	want := []DirectMessageResult{
		{
			ID:                  "test_id",
			Text:                "test_text",
			SenderID:            "test_sender_id",
			RecipientID:         "test_recipient_id",
			CreatedAt:           "Thu Nov 17 03:45:20 +0000 2011",
			SenderScreenName:    "test_sender_screen_name",
			RecipientScreenName: "test_recipient_screen_name",
			Sender: &UserResult{
				ID: "test_id",
			},
			Recipient: &UserResult{
				ID: "test_id",
			},
			InReplyTo: &DirectMessageResult{
				ID:                  "test_id",
				Text:                "test_text",
				SenderID:            "test_sender_id",
				RecipientID:         "test_recipient_id",
				CreatedAt:           "Thu Nov 17 03:45:20 +0000 2011",
				SenderScreenName:    "test_sender_screen_name",
				RecipientScreenName: "test_recipient_screen_name",
				Sender: &UserResult{
					ID: "test_id",
				},
				Recipient: &UserResult{
					ID: "test_id",
				},
			},
		},
		{
			ID:                  "test_id",
			Text:                "test_text",
			SenderID:            "test_sender_id",
			RecipientID:         "test_recipient_id",
			CreatedAt:           "Thu Nov 17 03:45:20 +0000 2011",
			SenderScreenName:    "test_sender_screen_name",
			RecipientScreenName: "test_recipient_screen_name",
			Sender: &UserResult{
				ID: "test_id",
			},
			Recipient: &UserResult{
				ID: "test_id",
			},
			InReplyTo: &DirectMessageResult{
				ID:                  "test_id",
				Text:                "test_text",
				SenderID:            "test_sender_id",
				RecipientID:         "test_recipient_id",
				CreatedAt:           "Thu Nov 17 03:45:20 +0000 2011",
				SenderScreenName:    "test_sender_screen_name",
				RecipientScreenName: "test_recipient_screen_name",
				Sender: &UserResult{
					ID: "test_id",
				},
				Recipient: &UserResult{
					ID: "test_id",
				},
			},
		},
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("direct_messages.conversation returned %+v, want %+v", user, want)
	}
}

func TestDirectMessagesService_New(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/direct_messages/new.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}, "in_reply_to": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}}`)
		if err != nil {
			t.Errorf("direct_messages.new mock server error: %+v", err)
		}
	})

	status, err := client.DirectMessages.New("test_id", "test_text", &DirectMessagesOptParams{
		Mode:        "test5",
		InReplyToID: "test6",
	})
	if err != nil {
		t.Errorf("direct_messages.new returned error: %v", err)
	}

	want := &DirectMessageResult{
		ID:                  "test_id",
		Text:                "test_text",
		SenderID:            "test_sender_id",
		RecipientID:         "test_recipient_id",
		CreatedAt:           "Thu Nov 17 03:45:20 +0000 2011",
		SenderScreenName:    "test_sender_screen_name",
		RecipientScreenName: "test_recipient_screen_name",
		Sender: &UserResult{
			ID: "test_id",
		},
		Recipient: &UserResult{
			ID: "test_id",
		},
		InReplyTo: &DirectMessageResult{
			ID:                  "test_id",
			Text:                "test_text",
			SenderID:            "test_sender_id",
			RecipientID:         "test_recipient_id",
			CreatedAt:           "Thu Nov 17 03:45:20 +0000 2011",
			SenderScreenName:    "test_sender_screen_name",
			RecipientScreenName: "test_recipient_screen_name",
			Sender: &UserResult{
				ID: "test_id",
			},
			Recipient: &UserResult{
				ID: "test_id",
			},
		},
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("direct_messages.new returned %+v, want %+v", status, want)
	}
}
