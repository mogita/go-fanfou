package fanfou

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mrjones/oauth"

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
				var err error
				client := NewClientWithOAuth(mockConsumerKey, mockConsumerSecret)
				accessToken := oauth.AccessToken{
					Token:  mockAccessToken,
					Secret: mockAccessSecret,
				}
				client.http, err = client.OAuthConsumer.MakeHttpClient(&accessToken)

				assert.Nil(t, err)
				assert.NotNil(t, client.http)

				clientType := reflect.TypeOf(client)
				for i := 0; i < clientType.NumMethod(); i++ {
					method := clientType.Method(i)
					if method.Name == "DoAuth" || method.Name == "GetTokenAndAuthURL" {
						continue
					}

					v := reflect.ValueOf(client).MethodByName(method.Name).Call([]reflect.Value{reflect.ValueOf(&ReqParams{})})
					fmt.Println(v[2])
				}
			},
		},
	}

	for _, testCase := range testCases {
		testCase.testFunc()
	}
}
