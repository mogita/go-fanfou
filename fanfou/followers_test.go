package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFollowersService_IDs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/followers/ids.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `["id1","id2","id3"]`)
		if err != nil {
			t.Errorf("followers.ids mock server error: %+v", err)
		}
	})

	users, err := client.Followers.IDs(&FollowersOptParams{
		ID:    "test_id",
		Count: 1,
		Page:  1,
	})
	if err != nil {
		t.Errorf("followers.ids returned error: %v", err)
	}

	want := &UserIDs{
		"id1",
		"id2",
		"id3",
	}

	if !reflect.DeepEqual(users, want) {
		t.Errorf("followers.ids returned %+v, want %+v", users, want)
	}
}
