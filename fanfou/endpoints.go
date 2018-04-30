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

func (client *httpClientWrapper) SearchPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSearchPublicTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUsers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSearchUsers, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUsers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
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

func (client *httpClientWrapper) BlocksIDs(params *ReqParams) (*responseBlockIDs, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiBlocksIDs, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksIDs: %+v", err)
	}

	ret := responseBlockIDs{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksBlocking(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiBlocksBlocking, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksBlocking: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksCreate(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiBlocksCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksExists(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiBlocksExists, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksExists: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksDestroy(params *ReqParams) (*responseUser, []byte, error) {
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

func (client *httpClientWrapper) UsersTagged(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersTagged, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagged: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersShow(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersShow, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersShow: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersTagList(params *ReqParams) (responseTags, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersTagList, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagList: %+v", err)
	}

	ret := responseTags{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersFollowers, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersRecommendation(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiUsersRecommendation, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersRecommendation: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersCancelRecommendation(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiUsersCancelRecommendation, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersCancelRecommendation: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersFriends(params *ReqParams) ([]*responseUser, []byte, error) {
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

func (client *httpClientWrapper) AccountVerifyCredentials(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountVerifyCredentials, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountVerifyCredentials: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateProfileImage(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest("image", "", params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateProfileImage: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountRateLimitStatus(params *ReqParams) (*responseRateLimitStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountRateLimitStatus, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountRateLimitStatus: %+v", err)
	}

	ret := responseRateLimitStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateProfile(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiAccountUpdateProfile, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateProfile: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotification(params *ReqParams) (*responseAccountNotification, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiAccountNotification, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotification: %+v", err)
	}

	ret := responseAccountNotification{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiAccountUpdateNotifyNum, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
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

func (client *httpClientWrapper) SavedSearchesCreate(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiSavedSearchesCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesCreate: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesDestroy(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiSavedSearchesDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesDestroy: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesShow(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiSavedSearchesShow, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesShow: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesList(params *ReqParams) ([]*responseSavedSearch, []byte, error) {
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

func (client *httpClientWrapper) PhotosUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiPhotosUserTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request PhotosUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) PhotosUpload(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest("photo", "", params)

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

func (client *httpClientWrapper) TrendsList(params *ReqParams) (*responseTrends, []byte, error) {
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

func (client *httpClientWrapper) FollowersIDs(params *ReqParams) ([]*responseUser, []byte, error) {
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

func (client *httpClientWrapper) FavoritesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFavoritesDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) Favorites(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFavorites, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request Favorites: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FavoritesCreate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFavoritesCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesCreate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiFriendshipsCreate   = apiBase + "/friendships/create.json"
	apiFriendshipsDestroy  = apiBase + "/friendships/destroy.json"
	apiFriendshipsRequests = apiBase + "/friendships/requests.json"
	apiFriendshipsDeny     = apiBase + "/friendships/deny.json"
	apiFriendshipsExists   = apiBase + "/friendships/exists.json"
	apiFriendshipsAccept   = apiBase + "/friendships/accept.json"
	apiFriendshipsShow     = apiBase + "/friendships/show.json"
)

func (client *httpClientWrapper) FriendshipsCreate(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFriendshipsCreate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsDestroy(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFriendshipsDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsRequests(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFriendshipsRequests, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsRequests: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsDeny(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFriendshipsDeny, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsDeny: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsExists(params *ReqParams) (*bool, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFriendshipsExists, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsExists: %+v", err)
	}

	var ret bool
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsAccept(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiFriendshipsAccept, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsAccept: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsShow(params *ReqParams) (*responseFriendship, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFriendshipsShow, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsShow: %+v", err)
	}

	ret := responseFriendship{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiFriendsIDs = apiBase + "/friends/ids.json"
)

func (client *httpClientWrapper) FriendsIDs(params *ReqParams) (responseFriendsIDs, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiFriendsIDs, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendsIDs: %+v", err)
	}

	ret := responseFriendsIDs{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

const (
	apiStatusesDestroy         = apiBase + "/statuses/destroy.json"
	apiStatusesHomeTimeline    = apiBase + "/statuses/home_timeline.json"
	apiStatusesPublicTimeline  = apiBase + "/statuses/public_timeline.json"
	apiStatusesReplies         = apiBase + "/statuses/replies.json"
	apiStatusesFollowers       = apiBase + "/statuses/followers.json"
	apiStatusesUpdate          = apiBase + "/statuses/update.json"
	apiStatusesUserTimeline    = apiBase + "/statuses/user_timeline.json"
	apiStatusesFriends         = apiBase + "/statuses/friends.json"
	apiStatusesContextTimeline = apiBase + "/statuses/context_timeline.json"
	apiStatusesMentions        = apiBase + "/statuses/mentions.json"
	apiStatusesShow            = apiBase + "/statuses/show.json"
)

func (client *httpClientWrapper) StatusesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiStatusesDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesHomeTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesHomeTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesHomeTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesPublicTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesReplies(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesReplies, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesReplies: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesFollowers, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesUpdate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiStatusesUpdate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUpdate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesUserTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesFriends, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesFriends: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesContextTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesContextTimeline, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesContextTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesMentions(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesMentions, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesMentions: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesShow(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiStatusesShow, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesShow: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

const (
	apiDirectMessagesDestroy          = apiBase + "/direct_messages/destroy.json"
	apiDirectMessagesConversation     = apiBase + "/direct_messages/conversation.json"
	apiDirectMessagesNew              = apiBase + "/direct_messages/new.json"
	apiDirectMessagesConversationList = apiBase + "/direct_messages/conversation_list.json"
	apiDirectMessagesInbox            = apiBase + "/direct_messages/inbox.json"
	apiDirectMessagesSent             = apiBase + "/direct_messages/sent.json"
)

func (client *httpClientWrapper) DirectMessagesDestroy(params *ReqParams) (*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiDirectMessagesDestroy, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesDestroy: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversation(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiDirectMessagesConversation, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesConversation: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesNew(params *ReqParams) (*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiDirectMessagesNew, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesNew: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversationList(params *ReqParams) ([]*responseDirectMessageConversationItem, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiDirectMessagesConversationList, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesConversationList: %+v", err)
	}

	ret := []*responseDirectMessageConversationItem{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesInbox(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiDirectMessagesInbox, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesInbox: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesSent(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, apiDirectMessagesSent, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesSent: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
