package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFriendsService_IDs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/friends/ids.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `["id1","id2","id3"]`)
		if err != nil {
			t.Errorf("friends.ids mock server error: %+v", err)
		}
	})

	users, _, err := client.Friends.IDs(&FriendsOptParams{
		ID:    "test_id",
		Count: 1,
		Page:  1,
	})
	if err != nil {
		t.Errorf("friends.ids returned error: %v", err)
	}

	want := &UserIDs{
		"id1",
		"id2",
		"id3",
	}

	if !reflect.DeepEqual(users, want) {
		t.Errorf("friends.ids returned %+v, want %+v", users, want)
	}
}
