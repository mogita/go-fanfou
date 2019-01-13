package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCommentsService_Add(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/update.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{
			"status":                "status text",
			"in_reply_to_status_id": "test1",
			"in_reply_to_user_id":   "test2",
			"repost_status_id":      "test3",
			"source":                "test4",
			"location":              "test7",
		})
		_, err := fmt.Fprint(w, `{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}`)
		if err != nil {
			t.Errorf("statuses.update mock server error: %+v", err)
		}
	})

	status, err := client.Statuses.Update("status text", &StatusesOptParams{
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Mode:              "test5",
		Format:            "test6",
		Location:          "test7",
	})
	if err != nil {
		t.Errorf("statuses.update returned error: %v", err)
	}

	want := &Status{
		ID:                "test_id",
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Location:          "test7",
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("statuses.update returned %+v, want %+v", status, want)
	}
}
