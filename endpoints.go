package fanfou

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiBase = "http://api.fanfou.com"
)

// search

const (
	apiSearchPublicTimeline = apiBase + "/search/public_timeline.json"
	apiSearchUsers          = apiBase + "/search/users.json"
	apiSearchUserTimeline   = apiBase + "/search/user_timeline.json"
)

func (client *baseClient) SearchPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSearchPublicTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) SearchUsers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSearchUsers, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUsers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) SearchUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSearchUserTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// blocks

const (
	apiBlocksIDs      = apiBase + "/blocks/ids.json"
	apiBlocksBlocking = apiBase + "/blocks/blocking.json"
	apiBlocksCreate   = apiBase + "/blocks/create.json"
	apiBlocksExists   = apiBase + "/blocks/exists.json"
	apiBlocksDestroy  = apiBase + "/blocks/destroy.json"
)

func (client *baseClient) BlocksIDs(params *ReqParams) (*responseBlockIDs, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiBlocksIDs, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksIDs: %+v", err)
	}

	ret := responseBlockIDs{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) BlocksBlocking(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiBlocksBlocking, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksBlocking: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) BlocksCreate(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiBlocksCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) BlocksExists(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiBlocksExists, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksExists: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) BlocksDestroy(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiBlocksDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiUsersTagged               = apiBase + "/users/tagged.json"
	apiUsersShow                 = apiBase + "/users/show.json"
	apiUsersTagList              = apiBase + "/users/tag_list.json"
	apiUsersFollowers            = apiBase + "/users/followers.json"
	apiUsersRecommendation       = apiBase + "/2/users/recommendation.json"
	apiUsersCancelRecommendation = apiBase + "/2/users/cancel_recommendation.json"
	apiUsersFriends              = apiBase + "/users/friends.json"
)

func (client *baseClient) UsersTagged(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersTagged, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagged: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) UsersShow(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersShow, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersShow: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) UsersTagList(params *ReqParams) (responseTags, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersTagList, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagList: %+v", err)
	}

	ret := responseTags{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) UsersFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersFollowers, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) UsersRecommendation(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersRecommendation, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersRecommendation: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) UsersCancelRecommendation(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiUsersCancelRecommendation, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersCancelRecommendation: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) UsersFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersFriends, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFriends: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

const (
	apiStatusesUpdate            = apiBase + "/statuses/update.json"
	apiAccountVerifyCredentials  = apiBase + "/account/verify_credentials.json"
	apiAccountUpdateProfile      = apiBase + "/account/update_profile.json"
	apiAccountUpdateNotifyNum    = apiBase + "/account/update_notify_num.json"
	apiSavedSearchesCreate       = apiBase + "/saved_searches/create.json"
	apiPhotosUpload              = apiBase + "/photos/upload.json"
	apiAccountUpdateProfileImage = apiBase + "/account/update_profile_image.json"
)

func (client *baseClient) StatusesUpdate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiStatusesUpdate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUpdate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}
