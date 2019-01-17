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

func TestDirectMessagesService_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/direct_messages/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}, "in_reply_to": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}}`)
		if err != nil {
			t.Errorf("direct_messages.destroy mock server error: %+v", err)
		}
	})

	status, err := client.DirectMessages.Destroy("test_id")
	if err != nil {
		t.Errorf("direct_messages.destroy returned error: %v", err)
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
		t.Errorf("direct_messages.destroy returned %+v, want %+v", status, want)
	}
}

func TestDirectMessagesService_ConversationList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/direct_messages/conversation_list.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"dm": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}, "otherid": "test_other_id", "msg_num": 11, "new_conv": true}, {"dm": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}, "otherid": "test_other_id", "msg_num": 11, "new_conv": true}]`)
		if err != nil {
			t.Errorf("direct_messages.conversation_list mock server error: %+v", err)
		}
	})

	result, err := client.DirectMessages.ConversationList(&DirectMessagesOptParams{
		Count: 1,
		Page:  1,
		Mode:  "test_mode",
	})
	if err != nil {
		t.Errorf("direct_messages.conversation_list returned error: %v", err)
	}

	want := &DirectMessageConversationListResult{
		{
			Dm: &DirectMessageResult{
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
				InReplyTo: nil,
			},
			Otherid: "test_other_id",
			MsgNum:  11,
			NewConv: true,
		},
		{
			Dm: &DirectMessageResult{
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
				InReplyTo: nil,
			},
			Otherid: "test_other_id",
			MsgNum:  11,
			NewConv: true,
		},
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("direct_messages.conversation_list returned %+v, want %+v", result, want)
	}
}

func TestDirectMessagesService_Inbox(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/direct_messages/inbox.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}, "in_reply_to": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}}, {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}, "in_reply_to": {"id": "test_id", "text": "test_text", "sender_id": "test_sender_id", "recipient_id": "test_recipient_id", "created_at": "Thu Nov 17 03:45:20 +0000 2011", "sender_screen_name": "test_sender_screen_name", "recipient_screen_name": "test_recipient_screen_name", "sender": {"id": "test_id"}, "recipient": {"id": "test_id"}}}]`)
		if err != nil {
			t.Errorf("direct_messages.inbox mock server error: %+v", err)
		}
	})

	user, err := client.DirectMessages.Inbox(&DirectMessagesOptParams{
		Page:    1,
		Count:   1,
		MaxID:   "test_max_id",
		SinceID: "test_since_id",
		Mode:    "test_mode",
	})
	if err != nil {
		t.Errorf("direct_messages.inbox returned error: %v", err)
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
		t.Errorf("direct_messages.inbox returned %+v, want %+v", user, want)
	}
}
