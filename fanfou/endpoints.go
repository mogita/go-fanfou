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
	"SearchPublicTimeline": endpointItem{
		URL:    apiBase + "/search/public_timeline.json",
		Method: http.MethodGet,
	},
	"SearchUsers": endpointItem{
		URL:    apiBase + "/search/users.json",
		Method: http.MethodGet,
	},
	"SearchUserTimeline": endpointItem{
		URL:    apiBase + "/search/user_timeline.json",
		Method: http.MethodGet,
	},

	"BlocksIDs": endpointItem{
		URL:    apiBase + "/blocks/ids.json",
		Method: http.MethodGet,
	},
	"BlocksBlocking": endpointItem{
		URL:    apiBase + "/blocks/blocking.json",
		Method: http.MethodGet,
	},
	"BlocksCreate": endpointItem{
		URL:    apiBase + "/blocks/create.json",
		Method: http.MethodPost,
	},
	"BlocksExists": endpointItem{
		URL:    apiBase + "/blocks/exists.json",
		Method: http.MethodGet,
	},
	"BlocksDestroy": endpointItem{
		URL:    apiBase + "/blocks/destroy.json",
		Method: http.MethodPost,
	},

	"UsersTagged": endpointItem{
		URL:    apiBase + "/users/tagged.json",
		Method: http.MethodGet,
	},
	"UsersShow": endpointItem{
		URL:    apiBase + "/users/show.json",
		Method: http.MethodGet,
	},
	"UsersTagList": endpointItem{
		URL:    apiBase + "/users/tag_list.json",
		Method: http.MethodGet,
	},
	"UsersFollowers": endpointItem{
		URL:    apiBase + "/users/followers.json",
		Method: http.MethodGet,
	},
	"UsersRecommendation": endpointItem{
		URL:    apiBase + "/2/users/recommendation.json",
		Method: http.MethodGet,
	},
	"UsersCancelRecommendation": endpointItem{
		URL:    apiBase + "/2/users/cancel_recommendation.json",
		Method: http.MethodPost,
	},
	"UsersFriends": endpointItem{
		URL:    apiBase + "/users/friends.json",
		Method: http.MethodGet,
	},

	"AccountVerifyCredentials": endpointItem{
		URL:    apiBase + "/account/verify_credentials.json",
		Method: http.MethodGet,
	},
	"AccountUpdateProfileImage": endpointItem{
		URL:    apiBase + "/account/update_profile_image.json",
		Method: http.MethodPost,
	},
	"AccountRateLimitStatus": endpointItem{
		URL:    apiBase + "/account/rate_limit_status.json",
		Method: http.MethodGet,
	},
	"AccountUpdateProfile": endpointItem{
		URL:    apiBase + "/account/update_profile.json",
		Method: http.MethodPost,
	},
	"AccountNotification": endpointItem{
		URL:    apiBase + "/account/notification.json",
		Method: http.MethodGet,
	},
	"AccountUpdateNotifyNum": endpointItem{
		URL:    apiBase + "/account/update_notify_num.json",
		Method: http.MethodPost,
	},
	"AccountNotifyNum": endpointItem{
		URL:    apiBase + "/account/notify_num.json",
		Method: http.MethodGet,
	},

	"SavedSearchesCreate": endpointItem{
		URL:    apiBase + "/saved_searches/create.json",
		Method: http.MethodPost,
	},
	"SavedSearchesDestroy": endpointItem{
		URL:    apiBase + "/saved_searches/destroy.json",
		Method: http.MethodPost,
	},
	"SavedSearchesShow": endpointItem{
		URL:    apiBase + "/saved_searches/show.json",
		Method: http.MethodGet,
	},
	"SavedSearchesList": endpointItem{
		URL:    apiBase + "/saved_searches/list.json",
		Method: http.MethodGet,
	},

	"PhotosUserTimeline": endpointItem{
		URL:    apiBase + "/photos/user_timeline.json",
		Method: http.MethodGet,
	},
	"PhotosUpload": endpointItem{
		URL:    apiBase + "/photos/upload.json",
		Method: http.MethodPost,
	},

	"TrendsList": endpointItem{
		URL:    apiBase + "/trends/list.json",
		Method: http.MethodGet,
	},

	"FollowersIDs": endpointItem{
		URL:    apiBase + "/followers/ids.json",
		Method: http.MethodGet,
	},

	"FavoritesDestroy": endpointItem{
		URL:    apiBase + "/favorites/destroy.json",
		Method: http.MethodPost,
	},
	"Favorites": endpointItem{
		URL:    apiBase + "/favorites.json",
		Method: http.MethodGet,
	},
	"FavoritesCreate": endpointItem{
		URL:    apiBase + "/favorites/create.json",
		Method: http.MethodPost,
	},

	"FriendshipsCreate": endpointItem{
		URL:    apiBase + "/friendships/create.json",
		Method: http.MethodPost,
	},
	"FriendshipsDestroy": endpointItem{
		URL:    apiBase + "/friendships/destroy.json",
		Method: http.MethodPost,
	},
	"FriendshipsRequests": endpointItem{
		URL:    apiBase + "/friendships/requests.json",
		Method: http.MethodGet,
	},
	"FriendshipsDeny": endpointItem{
		URL:    apiBase + "/friendships/deny.json",
		Method: http.MethodPost,
	},
	"FriendshipsExists": endpointItem{
		URL:    apiBase + "/friendships/exists.json",
		Method: http.MethodGet,
	},
	"FriendshipsAccept": endpointItem{
		URL:    apiBase + "/friendships/accept.json",
		Method: http.MethodPost,
	},
	"FriendshipsShow": endpointItem{
		URL:    apiBase + "/friendships/show.json",
		Method: http.MethodGet,
	},

	"FriendsIDs": endpointItem{
		URL:    apiBase + "/friends/ids.json",
		Method: http.MethodGet,
	},

	"StatusesDestroy": endpointItem{
		URL:    apiBase + "/statuses/destroy.json",
		Method: http.MethodPost,
	},
	"StatusesHomeTimeline": endpointItem{
		URL:    apiBase + "/statuses/home_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesPublicTimeline": endpointItem{
		URL:    apiBase + "/statuses/public_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesReplies": endpointItem{
		URL:    apiBase + "/statuses/replies.json",
		Method: http.MethodGet,
	},
	"StatusesFollowers": endpointItem{
		URL:    apiBase + "/statuses/followers.json",
		Method: http.MethodGet,
	},
	"StatusesUpdate": endpointItem{
		URL:    apiBase + "/statuses/update.json",
		Method: http.MethodPost,
	},
	"StatusesUserTimeline": endpointItem{
		URL:    apiBase + "/statuses/user_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesFriends": endpointItem{
		URL:    apiBase + "/statuses/friends.json",
		Method: http.MethodGet,
	},
	"StatusesContextTimeline": endpointItem{
		URL:    apiBase + "/statuses/context_timeline.json",
		Method: http.MethodGet,
	},
	"StatusesMentions": endpointItem{
		URL:    apiBase + "/statuses/mentions.json",
		Method: http.MethodGet,
	},
	"StatusesShow": endpointItem{
		URL:    apiBase + "/statuses/show.json",
		Method: http.MethodGet,
	},

	"DirectMessagesDestroy": endpointItem{
		URL:    apiBase + "/direct_messages/destroy.json",
		Method: http.MethodPost,
	},
	"DirectMessagesConversation": endpointItem{
		URL:    apiBase + "/direct_messages/conversation.json",
		Method: http.MethodGet,
	},
	"DirectMessagesNew": endpointItem{
		URL:    apiBase + "/direct_messages/new.json",
		Method: http.MethodPost,
	},
	"DirectMessagesConversationList": endpointItem{
		URL:    apiBase + "/direct_messages/conversation_list.json",
		Method: http.MethodGet,
	},
	"DirectMessagesInbox": endpointItem{
		URL:    apiBase + "/direct_messages/inbox.json",
		Method: http.MethodGet,
	},
	"DirectMessagesSent": endpointItem{
		URL:    apiBase + "/direct_messages/sent.json",
		Method: http.MethodGet,
	},
}

// search

func (client *httpClientWrapper) SearchPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["SearchPublicTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUsers(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["SearchUsers"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUsers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["SearchUserTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// blocks

func (client *httpClientWrapper) BlocksIDs(params *ReqParams) ([]responseUserID, []byte, error) {
	ep := endpoints["BlocksIDs"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksIDs: %+v", err)
	}

	ret := []responseUserID{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksBlocking(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["BlocksBlocking"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksBlocking: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksCreate(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["BlocksCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksExists(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["BlocksExists"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksExists: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksDestroy(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["BlocksDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksDestroy: %+v", err)
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
		return nil, nil, fmt.Errorf("Failed to request UsersTagged: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersShow(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["UsersShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersShow: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersTagList(params *ReqParams) ([]responseTag, []byte, error) {
	ep := endpoints["UsersTagList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagList: %+v", err)
	}

	ret := []responseTag{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersFollowers"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersRecommendation(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersRecommendation"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersRecommendation: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersCancelRecommendation(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["UsersCancelRecommendation"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersCancelRecommendation: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["UsersFriends"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFriends: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// account

func (client *httpClientWrapper) AccountVerifyCredentials(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["AccountVerifyCredentials"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

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
	ep := endpoints["AccountRateLimitStatus"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountRateLimitStatus: %+v", err)
	}

	ret := responseRateLimitStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateProfile(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["AccountUpdateProfile"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateProfile: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotification(params *ReqParams) (*responseAccountNotification, []byte, error) {
	ep := endpoints["AccountNotification"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotification: %+v", err)
	}

	ret := responseAccountNotification{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	ep := endpoints["AccountUpdateNotifyNum"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	ep := endpoints["AccountNotifyNum"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotifyNum: %+v", err)
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
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesCreate: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesDestroy(params *ReqParams) (*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesDestroy: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesShow(params *ReqParams) (*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesShow: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesList(params *ReqParams) ([]*responseSavedSearch, []byte, error) {
	ep := endpoints["SavedSearchesList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesList: %+v", err)
	}

	ret := []*responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// photos

func (client *httpClientWrapper) PhotosUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["PhotosUserTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

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

// trends

func (client *httpClientWrapper) TrendsList(params *ReqParams) (*responseTrends, []byte, error) {
	ep := endpoints["TrendsList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request TrendsList: %+v", err)
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
		return nil, nil, fmt.Errorf("Failed to request FollowersIDs: %+v", err)
	}

	ret := []responseUserID{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// favorites

func (client *httpClientWrapper) FavoritesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["FavoritesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) Favorites(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["Favorites"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request Favorites: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FavoritesCreate(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["FavoritesCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesCreate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// friendships

func (client *httpClientWrapper) FriendshipsCreate(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsCreate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsDestroy(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsRequests(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["FriendshipsRequests"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsRequests: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsDeny(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsDeny"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsDeny: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsExists(params *ReqParams) (bool, []byte, error) {
	ep := endpoints["FriendshipsExists"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to request FriendshipsExists: %+v", err)
	}

	var ret bool
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsAccept(params *ReqParams) (*responseUser, []byte, error) {
	ep := endpoints["FriendshipsAccept"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsAccept: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsShow(params *ReqParams) (*responseFriendship, []byte, error) {
	ep := endpoints["FriendshipsShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsShow: %+v", err)
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
		return nil, nil, fmt.Errorf("Failed to request FriendsIDs: %+v", err)
	}

	ret := []responseUserID{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// statuses

func (client *httpClientWrapper) StatusesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["StatusesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesHomeTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesHomeTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesHomeTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesPublicTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesReplies(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesReplies"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesReplies: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["StatusesFollowers"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesUpdate(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["StatusesUpdate"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUpdate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesUserTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	ep := endpoints["StatusesFriends"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesFriends: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesContextTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesContextTimeline"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesContextTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesMentions(params *ReqParams) ([]*responseStatus, []byte, error) {
	ep := endpoints["StatusesMentions"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesMentions: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesShow(params *ReqParams) (*responseStatus, []byte, error) {
	ep := endpoints["StatusesShow"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesShow: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// direct messages

func (client *httpClientWrapper) DirectMessagesDestroy(params *ReqParams) (*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesDestroy"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesDestroy: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversation(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesConversation"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesConversation: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesNew(params *ReqParams) (*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesNew"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesNew: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversationList(params *ReqParams) ([]*responseDirectMessageConversationItem, []byte, error) {
	ep := endpoints["DirectMessagesConversationList"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesConversationList: %+v", err)
	}

	ret := []*responseDirectMessageConversationItem{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesInbox(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesInbox"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesInbox: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesSent(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	ep := endpoints["DirectMessagesSent"]
	data, err := client.makeRequest(ep.Method, ep.URL, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesSent: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
