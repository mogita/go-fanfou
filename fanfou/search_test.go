package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSearchService_PublicTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/search/public_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("search.public_timeline mock server error: %+v", err)
		}
	})

	statuses, err := client.Search.PublicTimeline("test_query", &SearchOptParams{
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("search.public_timeline returned error: %v", err)
	}

	want := []Status{
		{
			ID:                "test_id",
			InReplyToStatusID: "test1",
			InReplyToUserID:   "test2",
			RepostStatusID:    "test3",
			Source:            "test4",
			Location:          "test7",
		},
		{
			ID:                "test_id",
			InReplyToStatusID: "test1",
			InReplyToUserID:   "test2",
			RepostStatusID:    "test3",
			Source:            "test4",
			Location:          "test7",
		},
	}

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("search.public_timeline returned %+v, want %+v", statuses, want)
	}
}
