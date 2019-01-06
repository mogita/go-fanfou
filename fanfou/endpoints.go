package fanfou

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type endpointItem struct {
	URL    string
	Method string
}

var endpoints = map[string]endpointItem{
	"SearchPublicTimeline": {
		URL:    apiBase + "/search/public_timeline.json",
		Method: http.MethodGet,
	},
	"SearchUsers": {
		URL:    apiBase + "/search/users.json",
		Method: http.MethodGet,
	},
	"SearchUserTimeline": {
		URL:    apiBase + "/search/user_timeline.json",
		Method: http.MethodGet,
	},

	"BlocksIDs": {
		URL:    apiBase + "/blocks/ids.json",
		Method: http.MethodGet,
	},
	"BlocksBlocking": {
		URL:    apiBase + "/blocks/blocking.json",
		Method: http.MethodGet,
	},
	"BlocksCreate": {
		URL:    apiBase + "/blocks/create.json",
		Method: http.MethodPost,
	},
	"BlocksExists": {
		URL:    apiBase + "/blocks/exists.json",
		Method: http.MethodGet,
	},
	"BlocksDestroy": {
		URL:    apiBase + "/blocks/destroy.json",
		Method: http.MethodPost,
	},

	"UsersTagged": {
		URL:    apiBase + "/users/tagged.json",
		Method: http.MethodGet,
	},
	"UsersShow": {
		URL:    apiBase + "/users/show.json",
		Method: http.MethodGet,
	},
	"UsersTagList": {
		URL:    apiBase + "/users/tag_list.json",
		Method: http.MethodGet,
	},
	"UsersFollowers": {
		URL:    apiBase + "/users/followers.json",
		Method: http.MethodGet,
	},
	"UsersRecommendation": {
		URL:    apiBase + "/2/users/recommendation.json",
		Method: http.MethodGet,
	},
	"UsersCancelRecommendation": {
		URL:    apiBase + "/2/users/cancel_recommendation.json",
		Method: http.MethodPost,
	},
	"UsersFriends": {
		URL:    apiBase + "/users/friends.json",
		Method: http.MethodGet,
	},

	"AccountVerifyCredentials": {
		URL:    apiBase + "/account/verify_credentials.json",
		Method: http.MethodGet,
	},
	"AccountUpdateProfileImage": {
		URL:    apiBase + "/account/update_profile_image.json",
		Method: http.MethodPost,
	},
	"AccountRateLimitStatus": {
		URL:    apiBase + "/account/rate_limit_status.json",
		Method: http.MethodGet,
	},
	"AccountUpdateProfile": {
		URL:    apiBase + "/account/update_profile.json",
		Method: http.MethodPost,
	},
	"AccountNotification": {
		URL:    apiBase + "/account/notification.json",
		Method: http.MethodGet,
	},
	"AccountUpdateNotifyNum": {
		URL:    apiBase + "/account/update_notify_num.json",
		Method: http.MethodPost,
	},
	"AccountNotifyNum": {
		URL:    apiBase + "/account/notify_num.json",
		Method: http.MethodGet,
	},

	"SavedSearchesCreate": {
		URL:    apiBase + "/saved_searches/create.json",
		Method: http.MethodPost,
	},
	"SavedSearchesDestroy": {
		URL:    apiBase + "/saved_searches/destroy.json",
		Method: http.MethodPost,
	},
	"SavedSearchesShow": {
		URL:    apiBase + "/saved_searches/show.json",
		Method: http.MethodGet,
	},
	"SavedSearchesList": {
		URL:    apiBase + "/saved_searches/list.json",
		Method: http.MethodGet,
	},

	"PhotosUserTimeline": {
		URL:    apiBase + "/photos/user_timeline.json",
		Method: http.MethodGet,
	},
	"PhotosUpload": {
		URL:    apiBase + "/photos/upload.json",
		Method: http.MethodPost,
	},

	"TrendsList": {
		URL:    apiBase + "/trends/list.json",
		Method: http.MethodGet,
	},

	"FollowersIDs": {
		URL:    apiBase + "/followers/ids.json",
		Method: http.MethodGet,
	},

	"FavoritesDestroy": {
		URL:    apiBase + "/favorites/destroy.json",
		Method: http.MethodPost,
	},
	"Favorites": {
		URL:    apiBase + "/favorites.json",
		Method: http.MethodGet,
	},
	"FavoritesCreate": {
		URL:    apiBase + "/favorites/create.json",
		Method: http.MethodPost,
	},

	"FriendshipsCreate": {
		URL:    apiBase + "/friendships/create.json",
		Method: http.MethodPost,
	},
	"FriendshipsDestroy": {
		URL:    apiBase + "/friendships/destroy.json",
		Method: http.MethodPost,
	},
	"FriendshipsRequests": {
		URL:    apiBase + "/friendships/requests.json",
		Method: http.MethodGet,
	},
	"FriendshipsDeny": {
		URL:    apiBase + "/friendships/deny.json",
		Method: http.MethodPost,
	},
	"FriendshipsExists": {
		URL:    apiBase + "/friendships/exists.json",
		Method: http.MethodGet,
	},
	"FriendshipsAccept": {
		URL:    apiBase + "/friendships/accept.json",
		Method: http.MethodPost,
	},
	"FriendshipsShow": {
		URL:    apiBase + "/friendships/show.json",
		Method: http.MethodGet,
	},

	"FriendsIDs": {
		URL:    apiBase + "/friends/ids.json",
		Method: http.MethodGet,
	},

	"StatusesDestroy": {
		URL:    apiBase + "/statuses/destroy.json",
		Method: http.MethodPost,
	},
	"StatusesHomeTimeline": {
		URL:    apiBase + "/statuses/home_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesPublicTimeline": {
		URL:    apiBase + "/statuses/public_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesReplies": {
		URL:    apiBase + "/statuses/replies.json",
		Method: http.MethodGet,
	},
	"StatusesFollowers": {
		URL:    apiBase + "/statuses/followers.json",
		Method: http.MethodGet,
	},
	"StatusesUpdate": {
		URL:    apiBase + "/statuses/update.json",
		Method: http.MethodPost,
	},
	"StatusesUserTimeline": {
		URL:    apiBase + "/statuses/user_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesFriends": {
		URL:    apiBase + "/statuses/friends.json",
		Method: http.MethodGet,
	},
	"StatusesContextTimeline": {
		URL:    apiBase + "/statuses/context_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesMentions": {
		URL:    apiBase + "/statuses/mentions.json",
		Method: http.MethodGet,
	},
	"StatusesShow": {
		URL:    apiBase + "/statuses/show.json",
		Method: http.MethodGet,
	},

	"DirectMessagesDestroy": {
		URL:    apiBase + "/direct_messages/destroy.json",
		Method: http.MethodPost,
	},
	"DirectMessagesConversation": {
		URL:    apiBase + "/direct_messages/conversation.json",
		Method: http.MethodGet,
	},
	"DirectMessagesNew": {
		URL:    apiBase + "/direct_messages/new.json",
		Method: http.MethodPost,
	},
	"DirectMessagesConversationList": {
		URL:    apiBase + "/direct_messages/conversation_list.json",
		Method: http.MethodGet,
	},
	"DirectMessagesInbox": {
		URL:    apiBase + "/direct_messages/inbox.json",
		Method: http.MethodGet,
	},
	"DirectMessagesSent": {
		URL:    apiBase + "/direct_messages/sent.json",
		Method: http.MethodGet,
	},
}

// search

func (client *httpClientWrapper) SearchPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["SearchPublicTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SearchPublicTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUsers(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["SearchUsers"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SearchUsers: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["SearchUserTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SearchUserTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// blocks

func (client *httpClientWrapper) BlocksIDs(params *ReqParams) ([]responseUserID, []byte, error) {
	ep := endpoints["BlocksIDs"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request BlocksIDs: %+v", err)
	}

	var ret []responseUserID
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksBlocking(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["BlocksBlocking"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request BlocksBlocking: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksCreate(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["BlocksCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request BlocksCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksExists(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["BlocksExists"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request BlocksExists: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksDestroy(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["BlocksDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request BlocksDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// users

func (client *httpClientWrapper) UsersTagged(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersTagged"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersTagged: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersShow(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["UsersShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersShow: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersTagList(params *ReqParams) ([]responseTag, []byte, error) {
	ep := endpoints["UsersTagList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersTagList: %+v", err)
	}

	var ret []responseTag
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersFollowers"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersFollowers: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersRecommendation(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersRecommendation"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersRecommendation: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersCancelRecommendation(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["UsersCancelRecommendation"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersCancelRecommendation: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersFriends"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request UsersFriends: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// account

func (client *httpClientWrapper) AccountVerifyCredentials(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["AccountVerifyCredentials"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountVerifyCredentials: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateProfileImage(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest("image", "", params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountUpdateProfileImage: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountRateLimitStatus(params *ReqParams) (*responseRateLimitStatus, []byte, error) {
	ep := endpoints["AccountRateLimitStatus"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountRateLimitStatus: %+v", err)
	}

	ret := responseRateLimitStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateProfile(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["AccountUpdateProfile"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountUpdateProfile: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotification(params *ReqParams) (*responseAccountNotification, []byte, error) {
	ep := endpoints["AccountNotification"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountNotification: %+v", err)
	}

	ret := responseAccountNotification{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	ep := endpoints["AccountUpdateNotifyNum"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountUpdateNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	ep := endpoints["AccountNotifyNum"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request AccountNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// saved searches

func (client *httpClientWrapper) SavedSearchesCreate(params *ReqParams) (*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SavedSearchesCreate: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesDestroy(params *ReqParams) (*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SavedSearchesDestroy: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesShow(params *ReqParams) (*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SavedSearchesShow: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesList(params *ReqParams) ([]*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request SavedSearchesList: %+v", err)
	}

	ret := []*responseSavedSearch{nil}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// photos

func (client *httpClientWrapper) PhotosUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["PhotosUserTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request PhotosUserTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) PhotosUpload(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest("photo", "", params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request PhotosUpload: %+v", err)
	}

	var ret responseStatus
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// trends

func (client *httpClientWrapper) TrendsList(params *ReqParams) (*responseTrends, []byte, error) {
	ep := endpoints["TrendsList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request TrendsList: %+v", err)
	}

	ret := responseTrends{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// followers

func (client *httpClientWrapper) FollowersIDs(params *ReqParams) ([]responseUserID, []byte, error) {
	ep := endpoints["FollowersIDs"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FollowersIDs: %+v", err)
	}

	var ret []responseUserID
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// favorites

func (client *httpClientWrapper) FavoritesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["FavoritesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FavoritesDestroy: %+v", err)
	}

	var ret responseStatus
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) Favorites(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["Favorites"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request Favorites: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FavoritesCreate(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["FavoritesCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FavoritesCreate: %+v", err)
	}

	var ret responseStatus
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// friendships

func (client *httpClientWrapper) FriendshipsCreate(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendshipsCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsDestroy(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendshipsDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsRequests(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["FriendshipsRequests"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendshipsRequests: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsDeny(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsDeny"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendshipsDeny: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsExists(params *ReqParams) (bool, []byte, error) {
	ep := endpoints["FriendshipsExists"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return false, nil, fmt.Errorf("failed to request FriendshipsExists: %+v", err)
	}

	var ret bool
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsAccept(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsAccept"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendshipsAccept: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsShow(params *ReqParams) (*responseFriendship, []byte, error) {
	ep := endpoints["FriendshipsShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendshipsShow: %+v", err)
	}

	ret := responseFriendship{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// friends

func (client *httpClientWrapper) FriendsIDs(params *ReqParams) ([]responseUserID, []byte, error) {
	ep := endpoints["FriendsIDs"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request FriendsIDs: %+v", err)
	}

	var ret []responseUserID
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// statuses

func (client *httpClientWrapper) StatusesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["StatusesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesDestroy: %+v", err)
	}

	var ret responseStatus
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesHomeTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesHomeTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesHomeTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesPublicTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesPublicTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesReplies(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesReplies"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesReplies: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["StatusesFollowers"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesFollowers: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesUpdate(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["StatusesUpdate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesUpdate: %+v", err)
	}

	var ret responseStatus
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesUserTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesUserTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["StatusesFriends"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesFriends: %+v", err)
	}

	var ret []*responseUser
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesContextTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesContextTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesContextTimeline: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesMentions(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesMentions"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesMentions: %+v", err)
	}

	var ret []*responseStatus
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesShow(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["StatusesShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request StatusesShow: %+v", err)
	}

	var ret responseStatus
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// direct messages

func (client *httpClientWrapper) DirectMessagesDestroy(params *ReqParams) (*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request DirectMessagesDestroy: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversation(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesConversation"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request DirectMessagesConversation: %+v", err)
	}

	ret := []*responseDirectMessage{nil}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesNew(params *ReqParams) (*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesNew"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request DirectMessagesNew: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversationList(params *ReqParams) ([]*responseDirectMessageConversationItem, []byte, error) {
	ep := endpoints["DirectMessagesConversationList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request DirectMessagesConversationList: %+v", err)
	}

	ret := []*responseDirectMessageConversationItem{nil}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesInbox(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesInbox"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request DirectMessagesInbox: %+v", err)
	}

	ret := []*responseDirectMessage{nil}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesSent(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesSent"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to request DirectMessagesSent: %+v", err)
	}

	ret := []*responseDirectMessage{nil}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
