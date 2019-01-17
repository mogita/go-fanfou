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

	want := &UserResult{
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

func TestUsersService_TagList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/tag_list.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `["tag1","tag2"]`)
		if err != nil {
			t.Errorf("users.tag_list mock server error: %+v", err)
		}
	})

	tags, err := client.Users.TagList(&UsersOptParams{
		ID: "test_id",
	})
	if err != nil {
		t.Errorf("users.tag_list returned error: %v", err)
	}

	want := []Tag{
		"tag1",
		"tag2",
	}

	if !reflect.DeepEqual(tags, want) {
		t.Errorf("users.tag_list returned %+v, want %+v", tags, want)
	}
}

func TestUsersService_Followers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/followers.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]`)
		if err != nil {
			t.Errorf("users.followers mock server error: %+v", err)
		}
	})

	users, err := client.Users.Followers(&UsersOptParams{
		ID:     "test_id",
		Page:   1,
		Count:  1,
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("users.followers returned error: %v", err)
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
		t.Errorf("users.followers returned %+v, want %+v", users, want)
	}
}

func TestUsersService_Friends(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/friends.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]`)
		if err != nil {
			t.Errorf("users.friends mock server error: %+v", err)
		}
	})

	users, err := client.Users.Friends(&UsersOptParams{
		ID:     "test_id",
		Page:   1,
		Count:  1,
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("users.friends returned error: %v", err)
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
		t.Errorf("users.friends returned %+v, want %+v", users, want)
	}
}

func TestUsersService_Recommendation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/2/users/recommendation.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"},{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}]`)
		if err != nil {
			t.Errorf("users.recommendation mock server error: %+v", err)
		}
	})

	users, err := client.Users.Recommendation(&UsersOptParams{
		Page:   1,
		Count:  1,
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("users.recommendation returned error: %v", err)
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
		t.Errorf("users.recommendation returned %+v, want %+v", users, want)
	}
}

func TestUsersService_CancelRecommendation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/2/users/cancel_recommendation.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("users.cancel_recommendation mock server error: %+v", err)
		}
	})

	user, err := client.Users.CancelRecommendation("test_id", &UsersOptParams{
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("users.cancel_recommendation returned error: %v", err)
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
		t.Errorf("users.cancel_recommendation returned %+v, want %+v", user, want)
	}
}
