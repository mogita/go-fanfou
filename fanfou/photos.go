package fanfou

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// PhotosService handles communication with the saved photos related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#photos
type PhotosService struct {
	client *Client
}

// PhotosOptParams specifies the optional params for search API
type PhotosOptParams struct {
	ID       string
	Status   string
	Source   string
	Location string
	SinceID  string
	MaxID    string
	Page     int64
	Count    int64
	Mode     string
	Format   string
}

// UserTimeline shall get photos of the specified user, or of the current user
// if no ID specified
// ID represents the user ID
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/photos.user-timeline
func (s *PhotosService) UserTimeline(opt *PhotosOptParams) ([]StatusResult, *string, error) {
	u := fmt.Sprintf("photos/user_timeline.json")
	params := url.Values{}

	if opt != nil {
		if opt.ID != "" {
			params.Add("id", opt.ID)
		}
		if opt.SinceID != "" {
			params.Add("since_id", opt.SinceID)
		}
		if opt.MaxID != "" {
			params.Add("max_id", opt.MaxID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatInt(opt.Count, 10))
		}
		if opt.Page != 0 {
			params.Add("page", strconv.FormatInt(opt.Page, 10))
		}
		if opt.Mode != "" {
			params.Add("mode", opt.Mode)
		}
		if opt.Format != "" {
			params.Add("format", opt.Format)
		}
	}

	u += "?" + params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u, "")
	if err != nil {
		return nil, nil, err
	}

	newStatuses := new([]StatusResult)
	resp, err := s.client.Do(req, newStatuses)
	if err != nil {
		return nil, nil, err
	}

	return *newStatuses, resp.BodyStrPtr, nil
}

// Upload shall send a new status with a photo
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/photos.upload
func (s *PhotosService) Upload(filePath string, opt *PhotosOptParams) (*StatusResult, *string, error) {
	u := fmt.Sprintf("photos/upload.json")
	params := map[string]string{}

	if opt != nil {
		if opt.Status != "" {
			params["status"] = opt.Status
		}
		if opt.Source != "" {
			params["source"] = opt.Source
		}
		if opt.Location != "" {
			params["location"] = opt.Location
		}
		if opt.Mode != "" {
			params["mode"] = opt.Mode
		}
		if opt.Format != "" {
			params["format"] = opt.Format
		}
	}

	if URL, err := url.Parse(filePath); err == nil && URL.Scheme != "" {
		localPath, err := fetchFile(URL.String())
		if err != nil {
			return nil, nil, err
		}

		defer func() {
			err := os.Remove(localPath)
			if err != nil {
				fmt.Printf("failed to remove tmp file: %+v\n", err)
			}
		}()

		filePath = localPath
	}

	req, err := s.client.NewUploadRequest(http.MethodPost, u, params, "photo", filePath)
	if err != nil {
		return nil, nil, err
	}

	newStatuses := new(StatusResult)
	resp, err := s.client.Do(req, newStatuses)
	if err != nil {
		return nil, nil, err
	}

	return newStatuses, resp.BodyStrPtr, nil
}

func fetchFile(URL string) (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Printf("failed to close body: %+v\n", err)
		}
	}()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(bodyBytes)

	suffix := ""
	switch contentType {
	case "image/jpeg":
		suffix = ".jpg"
	case "image/png":
		suffix = ".png"
	case "image/gif":
		suffix = ".gif"
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randStr := randomString(16)

	tmpFile, err := os.Create(os.TempDir() + "/go-fanfou-tmp-" + randStr + suffix)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(tmpFile.Name(), bodyBytes, 0644)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
