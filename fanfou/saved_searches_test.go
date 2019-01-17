package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSavedSearchesService_Show(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/saved_searches/show.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"id": 21071, "name": "fanfou|test", "query": "fanfou|test", "created_at": "Thu Nov 10 09:05:03 +0000 2011"}`)
		if err != nil {
			t.Errorf("saved_searches.show mock server error: %+v", err)
		}
	})

	user, err := client.SavedSearches.Show("test_id")
	if err != nil {
		t.Errorf("saved_searches.show returned error: %v", err)
	}

	want := &SavedSearchResult{
		ID:        21071,
		Name:      "fanfou|test",
		Query:     "fanfou|test",
		CreatedAt: "Thu Nov 10 09:05:03 +0000 2011",
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("saved_searches.show returned %+v, want %+v", user, want)
	}
}

func TestSavedSearchesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/saved_searches/list.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": 21071, "name": "fanfou|test", "query": "fanfou|test", "created_at": "Thu Nov 10 09:05:03 +0000 2011"},{"id": 21071, "name": "fanfou|test", "query": "fanfou|test", "created_at": "Thu Nov 10 09:05:03 +0000 2011"}]`)
		if err != nil {
			t.Errorf("saved_searches.list mock server error: %+v", err)
		}
	})

	user, err := client.SavedSearches.List()
	if err != nil {
		t.Errorf("saved_searches.list returned error: %v", err)
	}

	want := []SavedSearchResult{
		{
			ID:        21071,
			Name:      "fanfou|test",
			Query:     "fanfou|test",
			CreatedAt: "Thu Nov 10 09:05:03 +0000 2011",
		}, {
			ID:        21071,
			Name:      "fanfou|test",
			Query:     "fanfou|test",
			CreatedAt: "Thu Nov 10 09:05:03 +0000 2011",
		},
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("saved_searches.list returned %+v, want %+v", user, want)
	}
}
