package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_VerifyCredentials(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/verify_credentials.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("account.verify_credentials mock server error: %+v", err)
		}
	})

	user, err := client.Account.VerifyCredentials(&AccountOptParams{
		Mode:   "test5",
		Format: "test6",
	})
	if err != nil {
		t.Errorf("account.verify_credentials returned error: %v", err)
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
		t.Errorf("account.verify_credentials returned %+v, want %+v", user, want)
	}
}

func TestAccountService_RateLimitStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/rate_limit_status.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"reset_time": "Mon Nov 14 08:57:28 +0000 2011", "remaining_hits": 150, "hourly_limit": 150, "reset_time_in_seconds": 1321261048}`)
		if err != nil {
			t.Errorf("account.rate_limit_status mock server error: %+v", err)
		}
	})

	result, err := client.Account.RateLimitStatus()
	if err != nil {
		t.Errorf("account.rate_limit_status returned error: %v", err)
	}

	want := &RateLimitStatusResult{
		ResetTime:          "Mon Nov 14 08:57:28 +0000 2011",
		RemainingHits:      150,
		HourlyLimit:        150,
		ResetTimeInSeconds: 1321261048,
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("account.rate_limit_status returned %+v, want %+v", result, want)
	}
}
