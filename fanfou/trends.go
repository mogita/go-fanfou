package fanfou

import (
	"fmt"
)

// TrendsService handles communication with the trends related
// methods of the Fanfou API.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/API-Endpoints#trends
type TrendsService struct {
	client *Client
}

// Trends specifies Fanfou's trends data structure
type Trends struct {
	AsOf   string        `json:"as_of,omitempty"`
	Trends []*TrendsItem `json:"trends,omitempty"`
}

// TrendsItem specifies Fanfou's trends item data structure
type TrendsItem struct {
	Name  string `json:"name,omitempty"`
	Query string `json:"query,omitempty"`
	URL   string `json:"url,omitempty"`
}

// List shall get information about the most recent trends.
//
// Fanfou API docs: https://github.com/mogita/FanFouAPIDoc/wiki/trends.list
func (s *TrendsService) List() (*Trends, error) {
	u := fmt.Sprintf("trends/list.json")
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	trends := new(Trends)
	_, err = s.client.Do(req, trends)
	return trends, err
}