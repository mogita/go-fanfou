package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFavoritesService_IDs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/favorites/id.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("favorites.id mock server error: %+v", err)
		}
	})

	user, err := client.Favorites.IDs(&FavoritesOptParams{
		ID:     "test_id",
		Page:   1,
		Count:  1,
		Mode:   "test_mode",
		Format: "test_format",
	})
	if err != nil {
		t.Errorf("favorites.id returned error: %v", err)
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

	if !reflect.DeepEqual(user, want) {
		t.Errorf("favorites.id returned %+v, want %+v", user, want)
	}
}

func TestStatusesService_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/favorites/create.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}`)
		if err != nil {
			t.Errorf("favorites.create mock server error: %+v", err)
		}
	})

	status, err := client.Favorites.Create("test_id", &FavoritesOptParams{
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("favorites.create returned error: %v", err)
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
		t.Errorf("favorites.create returned %+v, want %+v", status, want)
	}
}
