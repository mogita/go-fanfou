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
