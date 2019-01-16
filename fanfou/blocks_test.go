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
