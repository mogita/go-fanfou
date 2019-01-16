package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBlocksService_IDs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/blocks/ids.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `["id1","id2","id3"]`)
		if err != nil {
			t.Errorf("blocks.ids mock server error: %+v", err)
		}
	})

	users, err := client.Blocks.IDs()
	if err != nil {
		t.Errorf("blocks.ids returned error: %v", err)
	}

	want := &UserIDs{
		"id1",
		"id2",
		"id3",
	}

	if !reflect.DeepEqual(users, want) {
		t.Errorf("blocks.ids returned %+v, want %+v", users, want)
	}
}

func TestBlocksService_Blocking(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/blocks/blocking.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]`)
		if err != nil {
			t.Errorf("blocks.blocking mock server error: %+v", err)
		}
	})

	users, err := client.Blocks.Blocking(&BlocksOptParams{
		Page:  1,
		Count: 1,
		Mode:  "test5",
	})
	if err != nil {
		t.Errorf("blocks.blocking returned error: %v", err)
	}

	want := []User{
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
	}

	if !reflect.DeepEqual(users, want) {
		t.Errorf("blocks.blocking returned %+v, want %+v", users, want)
	}
}
