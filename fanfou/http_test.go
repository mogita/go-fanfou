package fanfou

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func()
	}{
		{
			desc: "test newfileUploadRequest",
			testFunc: func() {
				httpReq, err := newfileUploadRequest("test_uri", map[string]string{"test_param": "test_value"}, "test_photo", "./http.go")

				ff, _, _ := httpReq.FormFile("test_photo")
				ff2, _, _ := httpReq.FormFile("test_photo_not_exists")

				assert.Nil(t, err)
				assert.Equal(t, httpReq.URL.Path, "test_uri")
				assert.Equal(t, httpReq.PostFormValue("test_param"), "test_value")
				assert.NotNil(t, ff)
				assert.Nil(t, ff2)
			},
		},
		{
			desc: "test paramsToURLValues",
			testFunc: func() {
				sample := ReqParams{
					ID:    "test_id",
					Name:  "test_name",
					Count: "66",
				}

				conv := paramsToURLValues(&sample)

				assert.Equal(t, conv.Get("id"), sample.ID)
				assert.Equal(t, conv.Get("name"), sample.Name)
				assert.Equal(t, conv.Get("count"), sample.Count)
				assert.Equal(t, conv.Get("non_existence"), "")
			},
		},
	}

	for _, testCase := range testCases {
		testCase.testFunc()
	}
}
