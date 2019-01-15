package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUsersService_Tagged(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/tagged.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]`)
		if err != nil {
			t.Errorf("users.tagged mock server error: %+v", err)
		}
	})

	users, err := client.Users.Tagged("test_tag", &UsersOptParams{
		Page:   1,
		Count:  1,
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("users.tagged returned error: %v", err)
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
		t.Errorf("users.tagged returned %+v, want %+v", users, want)
	}
}

func TestUsersService_Show(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/show.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("users.show mock server error: %+v", err)
		}
	})

	user, err := client.Users.Show(&UsersOptParams{
		ID:     "test_id",
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("users.show returned error: %v", err)
	}

	want := &User{
		ID:              "test_id",
		Name:            "test1",
		ScreenName:      "test2",
		Location:        "test3",
		Gender:          "test4",
		ProfileImageURL: "test7",
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("users.show returned %+v, want %+v", user, want)
	}
}
