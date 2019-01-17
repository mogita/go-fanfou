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

func TestSearchService_UserTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/search/user_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("search.user_timeline mock server error: %+v", err)
		}
	})

	statuses, err := client.Search.UserTimeline("test_query", &SearchOptParams{
		ID:      "test_id",
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("search.user_timeline returned error: %v", err)
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
		t.Errorf("search.user_timeline returned %+v, want %+v", statuses, want)
	}
}

func TestSearchService_Users(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/search/users.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"total_number":42,"users":[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]}`)
		if err != nil {
			t.Errorf("search.users mock server error: %+v", err)
		}
	})

	users, err := client.Search.Users("test_query", &SearchOptParams{
		Page:   1,
		Count:  1,
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("search.users returned error: %v", err)
	}

	want := &SearchUsersResult{
		TotalNumber: 42,
		Users: []UserResult{
			{
				ID:              "test_id",
				Name:            "test1",
				ScreenName:      "test2",
				Location:        "test3",
				Gender:          "test4",
				ProfileImageURL: "test7",
			},
			{
				ID:              "test_id",
				Name:            "test1",
				ScreenName:      "test2",
				Location:        "test3",
				Gender:          "test4",
				ProfileImageURL: "test7",
			},
		},
	}

	if !reflect.DeepEqual(users, want) {
		t.Errorf("search.users returned %+v, want %+v", users, want)
	}
}
