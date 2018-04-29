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
	apiAccountVerifyCredentials  = apiBase + "/account/verify_credentials.json"
	apiAccountUpdateProfileImage = apiBase + "/account/update_profile_image.json"
	apiAccountRateLimitStatus    = apiBase + "/account/rate_limit_status.json"
	apiAccountUpdateProfile      = apiBase + "/account/update_profile.json"
	apiAccountNotification       = apiBase + "/account/notification.json"
	apiAccountUpdateNotifyNum    = apiBase + "/account/update_notify_num.json"
	apiAccountNotifyNum          = apiBase + "/account/notify_num.json"
)

func (client *baseClient) AccountVerifyCredentials(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountVerifyCredentials, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountVerifyCredentials: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) AccountUpdateProfileImage(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiAccountUpdateProfileImage, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateProfileImage: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) AccountRateLimitStatus(params *ReqParams) (*responseRateLimitStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountRateLimitStatus, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountRateLimitStatus: %+v", err)
	}

	ret := responseRateLimitStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) AccountUpdateProfile(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiAccountUpdateProfile, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateProfile: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) AccountNotification(params *ReqParams) (*responseAccountNotification, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountNotification, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotification: %+v", err)
	}

	ret := responseAccountNotification{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) AccountUpdateNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiAccountUpdateNotifyNum, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) AccountNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountNotifyNum, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiSavedSearchesCreate  = apiBase + "/saved_searches/create.json"
	apiSavedSearchesDestroy = apiBase + "/saved_searches/destroy.json"
	apiSavedSearchesShow    = apiBase + "/saved_searches/show.json"
	apiSavedSearchesList    = apiBase + "/saved_searches/list.json"
)

func (client *baseClient) SavedSearchesCreate(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiSavedSearchesCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesCreate: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) SavedSearchesDestroy(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiSavedSearchesDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesDestroy: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) SavedSearchesShow(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSavedSearchesShow, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesShow: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) SavedSearchesList(params *ReqParams) ([]*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSavedSearchesList, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesList: %+v", err)
	}

	ret := []*responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

const (
	apiPhotosUserTimeline = apiBase + "/photos/user_timeline.json"
	apiPhotosUpload       = apiBase + "/photos/upload.json"
)

func (client *baseClient) PhotosUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiPhotosUserTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request PhotosUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) PhotosUpload(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiPhotosUpload, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request PhotosUpload: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiTrendsList = apiBase + "/trends/list.json"
)

func (client *baseClient) TrendsList(params *ReqParams) (*responseTrends, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiTrendsList, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request TrendsList: %+v", err)
	}

	ret := responseTrends{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiFollowersIDs = apiBase + "/followers/ids.json"
)

func (client *baseClient) FollowersIDs(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFollowersIDs, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FollowersIDs: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

const (
	apiFavoritesDestroy = apiBase + "/favorites/destroy.json"
	apiFavorites        = apiBase + "/favorites.json"
	apiFavoritesCreate  = apiBase + "/favorites/create.json"
)

func (client *baseClient) FavoritesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFavoritesDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) Favorites(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFavorites, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request Favorites: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *baseClient) FavoritesCreate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFavoritesCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesCreate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiStatusesUpdate = apiBase + "/statuses/update.json"
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
