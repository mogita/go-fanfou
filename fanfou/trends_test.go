package fanfou

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTrendsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/trends/list.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := fmt.Fprint(w, `{"as_of":"Thu Nov 10 09:57:23 +0000 2011", "trends":[{"name":"萤火一号","query":"萤火一号|火星|变轨","url":"http://fanfou.com/q/萤火一号%7C火星%7C变轨"}]}`)
		if err != nil {
			t.Errorf("trends.list mock server error: %+v", err)
		}
	})

	trends, err := client.Trends.List()
	if err != nil {
		t.Errorf("trends.list returned error: %v", err)
	}

	want := &Trends{
		AsOf: "Thu Nov 10 09:57:23 +0000 2011",
		Trends: []*TrendsItem{
			{
				Name:  "萤火一号",
				Query: "萤火一号|火星|变轨",
				URL:   "http://fanfou.com/q/萤火一号%7C火星%7C变轨",
			},
		},
	}

	if !reflect.DeepEqual(trends, want) {
		t.Errorf("trends.list returned %+v, want %+v", trends, want)
	}
}
