package fanfou

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClients(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func()
	}{
		{
			desc: "test NewClientWithOAuth",
			testFunc: func() {
				client := NewClientWithOAuth("", "")
				httpClient := &http.Client{}

				assert.Nil(t, client.http)
				assert.Equal(t, reflect.TypeOf(client.http), reflect.TypeOf(httpClient))
			},
		},
	}

	for _, testCase := range testCases {
		testCase.testFunc()
	}
}
