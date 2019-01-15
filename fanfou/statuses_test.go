package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStatusesService_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/update.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{
			"status":                "status text",
			"in_reply_to_status_id": "test1",
			"in_reply_to_user_id":   "test2",
			"repost_status_id":      "test3",
			"source":                "test4",
			"location":              "test7",
		})
		_, err := fmt.Fprint(w, `{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}`)
		if err != nil {
			t.Errorf("statuses.update mock server error: %+v", err)
		}
	})

	status, err := client.Statuses.Update("status text", &StatusesOptParams{
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Mode:              "test5",
		Format:            "test6",
		Location:          "test7",
	})
	if err != nil {
		t.Errorf("statuses.update returned error: %v", err)
	}

	want := &Status{
		ID:                "test_id",
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Location:          "test7",
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("statuses.update returned %+v, want %+v", status, want)
	}
}

func TestStatusesService_Show(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/show.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}`)
		if err != nil {
			t.Errorf("statuses.show mock server error: %+v", err)
		}
	})

	status, err := client.Statuses.Show("test_id", &StatusesOptParams{
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("statuses.show returned error: %v", err)
	}

	want := &Status{
		ID:                "test_id",
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Location:          "test7",
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("statuses.show returned %+v, want %+v", status, want)
	}
}

func TestStatusesService_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}`)
		if err != nil {
			t.Errorf("statuses.destroy mock server error: %+v", err)
		}
	})

	status, err := client.Statuses.Destroy("test_id", &StatusesOptParams{
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("statuses.destroy returned error: %v", err)
	}

	want := &Status{
		ID:                "test_id",
		InReplyToStatusID: "test1",
		InReplyToUserID:   "test2",
		RepostStatusID:    "test3",
		Source:            "test4",
		Location:          "test7",
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("statuses.destroy returned %+v, want %+v", status, want)
	}
}

func TestStatusesService_HomeTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/home_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("statuses.home_timeline mock server error: %+v", err)
		}
	})

	statuses, err := client.Statuses.HomeTimeline(&StatusesOptParams{
		ID:      "test_user_id",
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Page:    1,
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("statuses.home_timeline returned error: %v", err)
	}

	want := []Status{
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

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("statuses.home_timeline returned %+v, want %+v", statuses, want)
	}
}

func TestStatusesService_PublicTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/public_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("statuses.public_timeline mock server error: %+v", err)
		}
	})

	statuses, err := client.Statuses.PublicTimeline(&StatusesOptParams{
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Page:    1,
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("statuses.home_timeline returned error: %v", err)
	}

	want := []Status{
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

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("statuses.public_timeline returned %+v, want %+v", statuses, want)
	}
}

func TestStatusesService_Replies(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/replies.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("statuses.replies mock server error: %+v", err)
		}
	})

	statuses, err := client.Statuses.Replies(&StatusesOptParams{
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Page:    1,
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("statuses.replies returned error: %v", err)
	}

	want := []Status{
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

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("statuses.replies returned %+v, want %+v", statuses, want)
	}
}

func TestStatusesService_Mentions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/mentions.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("statuses.mentions mock server error: %+v", err)
		}
	})

	statuses, err := client.Statuses.Mentions(&StatusesOptParams{
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Page:    1,
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("statuses.mentions returned error: %v", err)
	}

	want := []Status{
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

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("statuses.mentions returned %+v, want %+v", statuses, want)
	}
}

func TestStatusesService_UserTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/user_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("statuses.user_timeline mock server error: %+v", err)
		}
	})

	statuses, err := client.Statuses.UserTimeline(&StatusesOptParams{
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Page:    1,
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("statuses.user_timeline returned error: %v", err)
	}

	want := []Status{
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

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("statuses.user_timeline returned %+v, want %+v", statuses, want)
	}
}

func TestStatusesService_ContextTimeline(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/statuses/context_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `[{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"},{"id": "test_id", "in_reply_to_status_id": "test1", "in_reply_to_user_id": "test2", "repost_status_id": "test3", "source": "test4", "location": "test7"}]`)
		if err != nil {
			t.Errorf("statuses.context_timeline mock server error: %+v", err)
		}
	})

	statuses, err := client.Statuses.ContextTimeline("test_id", &StatusesOptParams{
		SinceID: "test_since_id",
		MaxID:   "test_max_id",
		Page:    1,
		Count:   1,
		Mode:    "test5",
		Format:  "test6",
	})
	if err != nil {
		t.Errorf("statuses.context_timeline returned error: %v", err)
	}

	want := []Status{
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

	if !reflect.DeepEqual(statuses, want) {
		t.Errorf("statuses.context_timeline returned %+v, want %+v", statuses, want)
	}
}
