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
