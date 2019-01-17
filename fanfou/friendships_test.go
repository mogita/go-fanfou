package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFriendshipsService_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/friendships/create.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("friendships.create mock server error: %+v", err)
		}
	})

	user, err := client.Friendships.Create("test_id", &FriendshipsOptParams{
		Mode: "test5",
	})
	if err != nil {
		t.Errorf("friendships.create returned error: %v", err)
	}

	want := &UserResult{
		ID:              "test_id",
		Name:            "test1",
		ScreenName:      "test2",
		Location:        "test3",
		Gender:          "test4",
		ProfileImageURL: "test7",
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("friendships.create returned %+v, want %+v", user, want)
	}
}

func TestFriendshipsService_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("friendships.destroy mock server error: %+v", err)
		}
	})

	user, err := client.Friendships.Destroy("test_id", &FriendshipsOptParams{
		Mode: "test5",
	})
	if err != nil {
		t.Errorf("friendships.destroy returned error: %v", err)
	}

	want := &UserResult{
		ID:              "test_id",
		Name:            "test1",
		ScreenName:      "test2",
		Location:        "test3",
		Gender:          "test4",
		ProfileImageURL: "test7",
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("friendships.destroy returned %+v, want %+v", user, want)
	}
}

func TestFriendshipsService_Requests(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/friendships/requests.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]`)
		if err != nil {
			t.Errorf("friendships.requests mock server error: %+v", err)
		}
	})

	users, err := client.Friendships.Requests(&FriendshipsOptParams{
		Page:   1,
		Count:  1,
		Mode:   "test5",
		Format: "test_format",
	})
	if err != nil {
		t.Errorf("friendships.requests returned error: %v", err)
	}

	want := []UserResult{
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
		t.Errorf("friendships.requests returned %+v, want %+v", users, want)
	}
}

func TestFriendshipsService_Deny(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/friendships/deny.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("friendships.deny mock server error: %+v", err)
		}
	})

	user, err := client.Friendships.Deny("test_id", &FriendshipsOptParams{
		Mode: "test5",
	})
	if err != nil {
		t.Errorf("friendships.deny returned error: %v", err)
	}

	want := &UserResult{
		ID:              "test_id",
		Name:            "test1",
		ScreenName:      "test2",
		Location:        "test3",
		Gender:          "test4",
		ProfileImageURL: "test7",
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("friendships.deny returned %+v, want %+v", user, want)
	}
}

func TestFriendshipsService_Exists(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/friendships/exists.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `"true"`)
		if err != nil {
			t.Errorf("friendships.exists mock server error: %+v", err)
		}
	})

	user, err := client.Friendships.Exists("test_id_a", "test_id_b")
	if err != nil {
		t.Errorf("friendships.exists returned error: %v", err)
	}

	if !reflect.DeepEqual(user, true) {
		t.Errorf("friendships.exists returned %+v, want %+v", user, true)
	}
}
