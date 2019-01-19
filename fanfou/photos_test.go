package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPhotosService_UserTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/photos/user_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("photos.user_timeline mock server error: %+v", err)
		}
	})

	statuses, _, err := client.Photos.UserTimeline(&PhotosOptParams{
		ID:      "test_id",
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Count:   1,
		Page:    1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("photos.user_timeline returned error: %v", err)
	}

	want := []StatusResult{
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
		t.Errorf("photos.user_timeline returned %+v, want %+v", statuses, want)
	}
}

func TestPhotosService_Upload(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/photos/upload.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}`)
		if err != nil {
			t.Errorf("photos.upload mock server error: %+v", err)
		}
	})

	status, _, err := client.Photos.Upload("./photos.go", &PhotosOptParams{
		Status:   "test_status",
		Source:   "test_source",
		Location: "test_location",
		Mode:     "test5",
		Format:   "test6",
	})
	if err != nil {
		t.Errorf("photos.upload returned error: %v", err)
	}

	want := &StatusResult{
		ID:                "test_id",
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Location:          "test7",
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("photos.upload returned %+v, want %+v", status, want)
	}
}
