package fanfou

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
func (s *PhotosService) UserTimeline(opt *PhotosOptParams) ([]StatusResult, error) {
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
		return nil, err
	}

	newStatuses := new([]StatusResult)
	_, err = s.client.Do(req, newStatuses)
	if err != nil {
		return nil, err
	}

	return *newStatuses, nil
}

// Upload shall send a new status with a photo
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/photos.upload
func (s *PhotosService) Upload(filePath string, opt *PhotosOptParams) (*StatusResult, error) {
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

	req, err := s.client.NewUploadRequest(http.MethodPost, u, params, "photo", filePath)
	if err != nil {
		return nil, err
	}

	newStatuses := new(StatusResult)
	_, err = s.client.Do(req, newStatuses)
	if err != nil {
		return nil, err
	}

	return newStatuses, nil
}
