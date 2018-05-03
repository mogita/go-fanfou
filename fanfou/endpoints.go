package fanfou

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// search

var endpoints = map[string]string{
	"SearchPublicTimeline": apiBase + "/search/public_timeline.json",
	"SearchUsers":          apiBase + "/search/users.json",
	"SearchUserTimeline":   apiBase + "/search/user_timeline.json",

	"BlocksIDs":      apiBase + "/blocks/ids.json",
	"BlocksBlocking": apiBase + "/blocks/blocking.json",
	"BlocksCreate":   apiBase + "/blocks/create.json",
	"BlocksExists":   apiBase + "/blocks/exists.json",
	"BlocksDestroy":  apiBase + "/blocks/destroy.json",

	"UsersTagged":               apiBase + "/users/tagged.json",
	"UsersShow":                 apiBase + "/users/show.json",
	"UsersTagList":              apiBase + "/users/tag_list.json",
	"UsersFollowers":            apiBase + "/users/followers.json",
	"UsersRecommendation":       apiBase + "/2/users/recommendation.json",
	"UsersCancelRecommendation": apiBase + "/2/users/cancel_recommendation.json",
	"UsersFriends":              apiBase + "/users/friends.json",

	"AccountVerifyCredentials":  apiBase + "/account/verify_credentials.json",
	"AccountUpdateProfileImage": apiBase + "/account/update_profile_image.json",
	"AccountRateLimitStatus":    apiBase + "/account/rate_limit_status.json",
	"AccountUpdateProfile":      apiBase + "/account/update_profile.json",
	"AccountNotification":       apiBase + "/account/notification.json",
	"AccountUpdateNotifyNum":    apiBase + "/account/update_notify_num.json",
	"AccountNotifyNum":          apiBase + "/account/notify_num.json",

	"SavedSearchesCreate":  apiBase + "/saved_searches/create.json",
	"SavedSearchesDestroy": apiBase + "/saved_searches/destroy.json",
	"SavedSearchesShow":    apiBase + "/saved_searches/show.json",
	"SavedSearchesList":    apiBase + "/saved_searches/list.json",

	"PhotosUserTimeline": apiBase + "/photos/user_timeline.json",
	"PhotosUpload":       apiBase + "/photos/upload.json",

	"TrendsList": apiBase + "/trends/list.json",

	"FollowersIDs": apiBase + "/followers/ids.json",

	"FavoritesDestroy": apiBase + "/favorites/destroy.json",
	"Favorites":        apiBase + "/favorites.json",
	"FavoritesCreate":  apiBase + "/favorites/create.json",

	"FriendshipsCreate":   apiBase + "/friendships/create.json",
	"FriendshipsDestroy":  apiBase + "/friendships/destroy.json",
	"FriendshipsRequests": apiBase + "/friendships/requests.json",
	"FriendshipsDeny":     apiBase + "/friendships/deny.json",
	"FriendshipsExists":   apiBase + "/friendships/exists.json",
	"FriendshipsAccept":   apiBase + "/friendships/accept.json",
	"FriendshipsShow":     apiBase + "/friendships/show.json",

	"FriendsIDs": apiBase + "/friends/ids.json",

	"StatusesDestroy":         apiBase + "/statuses/destroy.json",
	"StatusesHomeTimeline":    apiBase + "/statuses/home_timeline.json",
	"StatusesPublicTimeline":  apiBase + "/statuses/public_timeline.json",
	"StatusesReplies":         apiBase + "/statuses/replies.json",
	"StatusesFollowers":       apiBase + "/statuses/followers.json",
	"StatusesUpdate":          apiBase + "/statuses/update.json",
	"StatusesUserTimeline":    apiBase + "/statuses/user_timeline.json",
	"StatusesFriends":         apiBase + "/statuses/friends.json",
	"StatusesContextTimeline": apiBase + "/statuses/context_timeline.json",
	"StatusesMentions":        apiBase + "/statuses/mentions.json",
	"StatusesShow":            apiBase + "/statuses/show.json",

	"DirectMessagesDestroy":          apiBase + "/direct_messages/destroy.json",
	"DirectMessagesConversation":     apiBase + "/direct_messages/conversation.json",
	"DirectMessagesNew":              apiBase + "/direct_messages/new.json",
	"DirectMessagesConversationList": apiBase + "/direct_messages/conversation_list.json",
	"DirectMessagesInbox":            apiBase + "/direct_messages/inbox.json",
	"DirectMessagesSent":             apiBase + "/direct_messages/sent.json",
}

func (client *httpClientWrapper) SearchPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["SearchPublicTimeline"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUsers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["SearchUsers"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUsers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) SearchUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["SearchUserTimeline"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SearchUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// blocks

func (client *httpClientWrapper) BlocksIDs(params *ReqParams) ([]responseUserID, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["BlocksIDs"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksIDs: %+v", err)
	}

	ret := []responseUserID{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksBlocking(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["BlocksBlocking"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksBlocking: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) BlocksCreate(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["BlocksCreate"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksExists(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["BlocksExists"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksExists: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) BlocksDestroy(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["BlocksDestroy"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request BlocksDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// users

func (client *httpClientWrapper) UsersTagged(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["UsersTagged"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagged: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersShow(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["UsersShow"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersShow: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersTagList(params *ReqParams) ([]responseTag, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["UsersTagList"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersTagList: %+v", err)
	}

	ret := []responseTag{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["UsersFollowers"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersRecommendation(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["UsersRecommendation"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersRecommendation: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) UsersCancelRecommendation(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["UsersCancelRecommendation"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersCancelRecommendation: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) UsersFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["UsersFriends"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UsersFriends: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// account

func (client *httpClientWrapper) AccountVerifyCredentials(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["AccountVerifyCredentials"], params)

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
	data, err := client.makeRequest(http.MethodGet, endpoints["AccountRateLimitStatus"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountRateLimitStatus: %+v", err)
	}

	ret := responseRateLimitStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateProfile(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["AccountUpdateProfile"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateProfile: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotification(params *ReqParams) (*responseAccountNotification, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["AccountNotification"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotification: %+v", err)
	}

	ret := responseAccountNotification{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountUpdateNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["AccountUpdateNotifyNum"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountUpdateNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) AccountNotifyNum(params *ReqParams) (*responseNotifyNum, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["AccountNotifyNum"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request AccountNotifyNum: %+v", err)
	}

	ret := responseNotifyNum{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// saved searches

func (client *httpClientWrapper) SavedSearchesCreate(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["SavedSearchesCreate"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesCreate: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesDestroy(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["SavedSearchesDestroy"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesDestroy: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesShow(params *ReqParams) (*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["SavedSearchesShow"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesShow: %+v", err)
	}

	ret := responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) SavedSearchesList(params *ReqParams) ([]*responseSavedSearch, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["SavedSearchesList"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request SavedSearchesList: %+v", err)
	}

	ret := []*responseSavedSearch{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// photos

func (client *httpClientWrapper) PhotosUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["PhotosUserTimeline"], params)

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
	data, err := client.makeRequest(http.MethodGet, endpoints["TrendsList"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request TrendsList: %+v", err)
	}

	ret := responseTrends{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// followers

func (client *httpClientWrapper) FollowersIDs(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["FollowersIDs"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FollowersIDs: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// favorites

func (client *httpClientWrapper) FavoritesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["FavoritesDestroy"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) Favorites(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["Favorites"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request Favorites: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FavoritesCreate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["FavoritesCreate"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FavoritesCreate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// friendships

func (client *httpClientWrapper) FriendshipsCreate(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["FriendshipsCreate"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsCreate: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsDestroy(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["FriendshipsDestroy"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsDestroy: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsRequests(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["FriendshipsRequests"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsRequests: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsDeny(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["FriendshipsDeny"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsDeny: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsExists(params *ReqParams) (bool, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["FriendshipsExists"], params)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to request FriendshipsExists: %+v", err)
	}

	var ret bool
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) FriendshipsAccept(params *ReqParams) (*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["FriendshipsAccept"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsAccept: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) FriendshipsShow(params *ReqParams) (*responseFriendship, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["FriendshipsShow"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendshipsShow: %+v", err)
	}

	ret := responseFriendship{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// friends

func (client *httpClientWrapper) FriendsIDs(params *ReqParams) ([]responseUserID, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["FriendsIDs"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request FriendsIDs: %+v", err)
	}

	ret := []responseUserID{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// statuses

func (client *httpClientWrapper) StatusesDestroy(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["StatusesDestroy"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesDestroy: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesHomeTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesHomeTimeline"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesHomeTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesPublicTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesPublicTimeline"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesPublicTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesReplies(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesReplies"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesReplies: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFollowers(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesFollowers"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesFollowers: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesUpdate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["StatusesUpdate"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUpdate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) StatusesUserTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesUserTimeline"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUserTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesFriends(params *ReqParams) ([]*responseUser, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesFriends"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesFriends: %+v", err)
	}

	ret := []*responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesContextTimeline(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesContextTimeline"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesContextTimeline: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesMentions(params *ReqParams) ([]*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesMentions"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesMentions: %+v", err)
	}

	ret := []*responseStatus{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) StatusesShow(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["StatusesShow"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesShow: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

// direct messages

func (client *httpClientWrapper) DirectMessagesDestroy(params *ReqParams) (*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["DirectMessagesDestroy"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesDestroy: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversation(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["DirectMessagesConversation"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesConversation: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesNew(params *ReqParams) (*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, endpoints["DirectMessagesNew"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesNew: %+v", err)
	}

	ret := responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *httpClientWrapper) DirectMessagesConversationList(params *ReqParams) ([]*responseDirectMessageConversationItem, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["DirectMessagesConversationList"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesConversationList: %+v", err)
	}

	ret := []*responseDirectMessageConversationItem{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesInbox(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["DirectMessagesInbox"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesInbox: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

func (client *httpClientWrapper) DirectMessagesSent(params *ReqParams) ([]*responseDirectMessage, []byte, error) {
	data, err := client.makeRequest(http.MethodGet, endpoints["DirectMessagesSent"], params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request DirectMessagesSent: %+v", err)
	}

	ret := []*responseDirectMessage{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
