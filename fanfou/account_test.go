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

func TestAccountService_UpdateProfile(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/update_profile.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"id": "test_id", "name": "test1", "screen_name": "test2", "location": "test3", "gender": "test4", "profile_image_url": "test7"}`)
		if err != nil {
			t.Errorf("account.update_profile mock server error: %+v", err)
		}
	})

	user, err := client.Account.UpdateProfile(&AccountOptParams{
		Mode:        "test5",
		Format:      "test6",
		URL:         "test_url",
		Location:    "test_location",
		Description: "test_description",
		Name:        "test_name",
		Email:       "test_email",
	})
	if err != nil {
		t.Errorf("account.update_profile returned error: %v", err)
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
		t.Errorf("account.update_profile returned %+v, want %+v", user, want)
	}
}

func TestAccountService_Notification(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/notification.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"mentions": 1, "direct_messages": 0, "friend_requests": 5}`)
		if err != nil {
			t.Errorf("account.notification mock server error: %+v", err)
		}
	})

	result, err := client.Account.Notification()
	if err != nil {
		t.Errorf("account.notification returned error: %v", err)
	}

	want := &NotificationResult{
		Mentions:       1,
		DirectMessages: 0,
		FriendRequests: 5,
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("account.notification returned %+v, want %+v", result, want)
	}
}

func TestAccountService_NotifyNum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/notify_num.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"result": "ok", "notify_num": 5}`)
		if err != nil {
			t.Errorf("account.notify_num mock server error: %+v", err)
		}
	})

	result, err := client.Account.NotifyNum()
	if err != nil {
		t.Errorf("account.notify_num returned error: %v", err)
	}

	want := &NotifyNumResult{
		Result:    "ok",
		NotifyNum: 5,
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("account.notify_num returned %+v, want %+v", result, want)
	}
}

func TestAccountService_UpdateNotifyNum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/update_notify_num.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := fmt.Fprint(w, `{"result": "ok", "notify_num": 5}`)
		if err != nil {
			t.Errorf("account.update_notify_num mock server error: %+v", err)
		}
	})

	result, err := client.Account.UpdateNotifyNum(&AccountOptParams{
		NotifyNum: 5,
	})
	if err != nil {
		t.Errorf("account.update_notify_num returned error: %v", err)
	}

	want := &NotifyNumResult{
		Result:    "ok",
		NotifyNum: 5,
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("account.update_notify_num returned %+v, want %+v", result, want)
	}
}
