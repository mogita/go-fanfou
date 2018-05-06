package fanfou

import (
	"fmt"
	"net/http"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

const (
	mockConsumerKey    = "mock_consumer_key"
	mockConsumerSecret = "mock_consumer_secret"
	mockRequestToken   = "mock_request_token"
	mockRequestSecret  = "mock_request_secret"
	mockAccessToken    = "mock_access_token"
	mockAccessSecret   = "mock_access_secret"
)

const (
	contentTypeJSON = "application/json; charset=utf-8"
	contentTypeHTML = "text/html; charset=utf-8"
)

type mockResult struct {
	Code        int
	ContentType string
	Body        string
}

type mockEndpointItem struct {
	URL         string
	Method      string
	Result200   mockResult
	Result400   mockResult
	Result500   mockResult
	ResultChaos mockResult
}

const (
	mockStatus = `{"created_at":"Wed Nov 09 07:15:21 +0000 2011","id":"UcIlC04F2pQ","rawid":123456,"text":"看看","source":"网页","truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","repost_status_id":"gcIghPhCxQ","repost_status":{"id":"CKT71D9TaCs","created_at":"Fri May 04 15:22:00 +0000 2018","in_reply_to_screen_name":"","source":"<a href=\"http://fan.zico.im\" target=\"_blank\">有饭</a>","rawid":225235039,"in_reply_to_status_id":"","location":"小千的怀里","in_reply_to_user_id":"","text":"这个点了我还在吃西瓜 我不胖谁胖。","truncated":false,"favorited":false,"is_self":false,"user":{"location":"小千的怀里","birthday":"1995-06-07","statuses_count":7511,"url":"http://www.JacksonYi.com","profile_sidebar_fill_color":"#F4F8ED","profile_sidebar_border_color":"#CEDDC2","notifications":false,"friends_count":99,"unique_id":"~GMquI3wphCM","name":"千与椒","screen_name":"千与椒","utc_offset":28800,"profile_background_color":"#E9F0DB","id":"Jackson1128","gender":"女","protected":false,"profile_background_image_url":"http://s3.meituan.net/v1/mss_3d027b52ec5a4d589e68050845611e68/avatar/b0/00/sr/av_1477496505.jpg?1477496505","profile_image_url_large":"http://s3.meituan.net/v1/mss_3d027b52ec5a4d589e68050845611e68/avatar/l0/00/sr/av.jpg?1448762617","profile_background_tile":true,"profile_image_url":"http://s3.meituan.net/v1/mss_3d027b52ec5a4d589e68050845611e68/avatar/s0/00/sr/av.jpg?1448762617","followers_count":99,"following":false,"profile_text_color":"#8D7947","created_at":"Thu Jul 23 14:03:43 +0000 2015","photo_count":1341,"favourites_count":0,"profile_link_color":"#AF838C","description":"无趣。日常叨逼叨。"},"photo":{"thumburl":"http://photo.fanfou.com/v1/mss_3d027b52ec5a4d589e68050845611e68/ff/n0/0f/pd/64_213740.jpg@120w_120h_1l.jpg","largeurl":"http://photo.fanfou.com/v1/mss_3d027b52ec5a4d589e68050845611e68/ff/n0/0f/pd/64_213740.jpg@596w_1l.jpg","url":"http://fanfou.com/photo/K4OVN_lU4UE","imageurl":"http://photo.fanfou.com/v1/mss_3d027b52ec5a4d589e68050845611e68/ff/n0/0f/pd/64_213740.jpg@200w_200h_1l.jpg"}},"repost_user_id":"clock","repost_screen_name":"钟钟","favorited":false,"user":{"id":"superisaac","name":"杲i杲","screen_name":"杲i杲","location":"北京 朝阳区","gender":"男","birthday":"","description":"爱工作, 爱编程, 爱摄影, 爱暴走, 也爱哗哗的灌水. 我不是富二代, 我是技术宅男, 我要减肥.","profile_image_url":"http://avatar3.fanfou.com/s0/01/3a/ve.jpg?1304991298","profile_image_url_large":"http://avatar3.fanfou.com/l0/01/3a/ve.jpg?1304991298","url":"","protected":false,"followers_count":1329,"friends_count":2011,"favourites_count":29,"statuses_count":5259,"following":true,"notifications":true,"created_at":"Thu Feb 24 10:01:24 +0000 2011","utc_offset":28800,"profile_background_color":"#C4B9A3","profile_text_color":"#3C2215","profile_link_color":"#BC834A","profile_sidebar_fill_color":"#F0EADC","profile_sidebar_border_color":"#FFFFFF","profile_background_image_url":"http://static.fanfou.com/img/bg/14.jpg","profile_background_tile":false},"photo":{"imageurl":"http://photo.fanfou.com/s0/02/f2/7q_317973.jpg","thumburl":"http://photo.fanfou.com/t0/02/f2/7q_317973.jpg","largeurl":"http://photo.fanfou.com/n0/02/f2/7q_317973.jpg"}}`

	mockUser = `{"id":"test","name":"测试昵称","screen_name":"测试昵称","location":"北京 海淀区","gender":"男","birthday":"2105-03-11","description":"测试帐号","profile_image_url":"http://avatar3.fanfou.com/s0/00/5n/sk.jpg?1320913295","profile_image_url_large":"http://avatar3.fanfou.com/l0/00/5n/sk.jpg?1320913295","url":"http://fanfou.com/test","protected":true,"followers_count":9,"friends_count":16,"favourites_count":23,"statuses_count":124,"following":false,"notifications":false,"created_at":"Sat Jun 09 23:56:33 +0000 2007","utc_offset":28800,"profile_background_color":"#ffffe5","profile_text_color":"#004040","profile_link_color":"#ff0000","profile_sidebar_fill_color":"#ffefbf","profile_sidebar_border_color":"#ffac80","profile_background_image_url":"http://avatar.fanfou.com/b0/00/5n/sk_1320749993.jpg","profile_background_tile":true,"status":{"created_at":"Thu Nov 10 09:37:34 +0000 2011","id":"XRFWGErKgGI","text":"这是神马？","source":"<a href=\"http://abc.fanfouapps.com\" target=\"_blank\">ABC</a>","truncated":false,"in_reply_to_lastmsg_id":"","in_reply_to_user_id":"","favorited":false,"in_reply_to_screen_name":""}}`

	mockAccountRateLimit = `{"reset_time":"Mon Nov 14 08:57:28 +0000 2011","remaining_hits":150,"hourly_limit":150,"reset_time_in_seconds":1321261048}`

	mockAccountNotificationStatus = `{"mentions":1,"direct_messages":0,"friend_requests":5}`

	mockAccountNotifyNum = `{"result":"ok","notify_num":5}`

	mockSavedSearch = `{"id":21071,"query":"fanfou|test","name":"fanfou|test","created_at":"Thu Nov 10 09:05:03 +0000 2011"}`

	mockTrends = `{"as_of":"Thu Nov 10 09:57:23 +0000 2011","trends":[{"name":"萤火一号","query":"萤火一号|火星|变轨","url":"http://fanfou.com/q/%E8%90%A4%E7%81%AB%E4%B8%80%E5%8F%B7%7C%E7%81%AB%E6%98%9F%7C%E5%8F%98%E8%BD%A8"},{"name":"土耳其地震","query":"土耳其|地震","url":"http://fanfou.com/q/%E5%9C%9F%E8%80%B3%E5%85%B6%7C%E5%9C%B0%E9%9C%87"},{"name":"《失恋33天》","query":"33天|白百何","url":"http://fanfou.com/q/33%E5%A4%A9%7C%E7%99%BD%E7%99%BE%E4%BD%95"},{"name":"股市大跌","query":"股市|国债","url":"http://fanfou.com/q/%E8%82%A1%E5%B8%82%7C%E5%9B%BD%E5%80%BA"},{"name":"光棍节","query":"光棍|神棍|六一","url":"http://fanfou.com/q/%E5%85%89%E6%A3%8D%7C%E7%A5%9E%E6%A3%8D%7C%E5%85%AD%E4%B8%80"},{"name":"北方降温","query":"降温|冷空气","url":"http://fanfou.com/q/%E9%99%8D%E6%B8%A9%7C%E5%86%B7%E7%A9%BA%E6%B0%94"}]}`

	mockFriendship = `{"relationship":{"source":{"id":"test","screen_name":"测试昵称","following":"false","followed_by":"false","notifications_enabled":"false","blocking":"true"},"target":{"id":"debug","screen_name":"debug","following":"false","followed_by":"false"}}}`

	mockDirectMessage = `{"id":"wgFPBXHGdwQ","text":"test_direct_message","sender_id":"test","recipient_id":"lisztli","created_at":"Thu Nov 17 03:45:20 +0000 2011","sender_screen_name":"测试昵称","recipient_screen_name":"LisztLi","sender":{"id":"test","name":"测试昵称","screen_name":"测试昵称","location":"北京 海淀区","gender":"男","birthday":"2105-03-11","description":"测试帐号","profile_image_url":"http://avatar3.fanfou.com/s0/00/5n/sk.jpg?1320913295","profile_image_url_large":"http://avatar3.fanfou.com/l0/00/5n/sk.jpg?1320913295","url":"http://fanfou.com/test","protected":true,"followers_count":6,"friends_count":12,"favourites_count":25,"statuses_count":134,"following":false,"notifications":false,"created_at":"Sat Jun 09 23:56:33 +0000 2007","utc_offset":28800,"profile_background_color":"#ffffe5","profile_text_color":"#004040","profile_link_color":"#ff0000","profile_sidebar_fill_color":"#ffefbf","profile_sidebar_border_color":"#ffac80","profile_background_image_url":"http://avatar.fanfou.com/b0/00/5n/sk_1321293707.jpg","profile_background_tile":true},"recipient":{"id":"lisztli","name":"LisztLi","screen_name":"LisztLi","location":"北京 海淀区","gender":"男","birthday":"1986-02-07","description":"","profile_image_url":"http://avatar3.fanfou.com/s0/00/3z/3k.jpg?1308030845","profile_image_url_large":"http://avatar3.fanfou.com/l0/00/3z/3k.jpg?1308030845","url":"","protected":false,"followers_count":162,"friends_count":143,"favourites_count":12,"statuses_count":2668,"following":false,"notifications":false,"created_at":"Wed Sep 05 01:11:07 +0000 2007","utc_offset":28800,"profile_background_color":"#acdae5","profile_text_color":"#222222","profile_link_color":"#0066cc","profile_sidebar_fill_color":"#e2f2da","profile_sidebar_border_color":"#b2d1a3","profile_background_image_url":"http://avatar.fanfou.com/b0/00/3z/3k_1308411449.jpg","profile_background_tile":true},"in_reply_to":{"id":"2077","text":"reply_direct_message","sender_id":"lisztli","recipient_id":"test","created_at":"Thu Nov 16 03:45:20 +0000 2011","recipient_screen_name":"测试昵称","sender_screen_name":"LisztLi","recipient":{"id":"test","name":"测试昵称","screen_name":"测试昵称","location":"北京 海淀区","gender":"男","birthday":"2105-03-11","description":"测试帐号","profile_image_url":"http://avatar3.fanfou.com/s0/00/5n/sk.jpg?1320913295","profile_image_url_large":"http://avatar3.fanfou.com/l0/00/5n/sk.jpg?1320913295","url":"http://fanfou.com/test","protected":true,"followers_count":6,"friends_count":12,"favourites_count":25,"statuses_count":134,"following":false,"notifications":false,"created_at":"Sat Jun 09 23:56:33 +0000 2007","utc_offset":28800,"profile_background_color":"#ffffe5","profile_text_color":"#004040","profile_link_color":"#ff0000","profile_sidebar_fill_color":"#ffefbf","profile_sidebar_border_color":"#ffac80","profile_background_image_url":"http://avatar.fanfou.com/b0/00/5n/sk_1321293707.jpg","profile_background_tile":true},"sender":{"id":"lisztli","name":"LisztLi","screen_name":"LisztLi","location":"北京 海淀区","gender":"男","birthday":"1986-02-07","description":"","profile_image_url":"http://avatar3.fanfou.com/s0/00/3z/3k.jpg?1308030845","profile_image_url_large":"http://avatar3.fanfou.com/l0/00/3z/3k.jpg?1308030845","url":"","protected":false,"followers_count":162,"friends_count":143,"favourites_count":12,"statuses_count":2668,"following":false,"notifications":false,"created_at":"Wed Sep 05 01:11:07 +0000 2007","utc_offset":28800,"profile_background_color":"#acdae5","profile_text_color":"#222222","profile_link_color":"#0066cc","profile_sidebar_fill_color":"#e2f2da","profile_sidebar_border_color":"#b2d1a3","profile_background_image_url":"http://avatar.fanfou.com/b0/00/3z/3k_1308411449.jpg","profile_background_tile":true}}}`

	mockDirectMessageConversationItem = `{"dm":{"id":"-SVsfHE_1RU","text":"asdf","sender_id":"zengke","recipient_id":"moon","created_at":"Wed Nov 02 07:49:55 +0000 2011","sender_screen_name":"曾科","recipient_screen_name":"穆荣均","sender":{"id":"zengke","name":"曾科","screen_name":"曾科","location":"","gender":"","birthday":"","description":"mujixx","profile_image_url":"http://avatar1.fanfou.com/s0/00/bm/pt.jpg?1320216698","profile_image_url_large":"http://avatar1.fanfou.com/l0/00/bm/pt.jpg?1320216698","url":"","protected":true,"followers_count":0,"friends_count":0,"favourites_count":0,"statuses_count":0,"following":true,"notifications":true,"created_at":"Mon Oct 06 11:03:08 +0000 2008","utc_offset":28800},"recipient":{"id":"moon","name":"穆荣均","screen_name":"穆荣均","location":"北京","gender":"男","birthday":"","description":"..","profile_image_url":"http://avatar.fanfou.com/s0/00/3e/r4.jpg?1319098444","profile_image_url_large":"http://avatar.fanfou.com/l0/00/3e/r4.jpg?1319098444","url":"","protected":false,"followers_count":0,"friends_count":0,"favourites_count":0,"statuses_count":0,"following":false,"notifications":false,"created_at":"Sat May 12 15:58:58 +0000 2007","utc_offset":28800}},"otherid":"zengke","msg_num":11,"new_conv":true}`

	mockClientError = `{"request":"/account/show.json","error":"参数错误"}`

	mockServerError = ``

	mockChaosError = ``
)

var mockEndpoints = map[string]mockEndpointItem{
	"SearchPublicTimeline": mockEndpointItem{
		URL:    endpoints["SearchPublicTimeline"].URL,
		Method: endpoints["SearchPublicTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"SearchUsers": mockEndpointItem{
		URL:    endpoints["SearchUsers"].URL,
		Method: endpoints["SearchUsers"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"SearchUserTimeline": mockEndpointItem{
		URL:    endpoints["SearchUserTimeline"].URL,
		Method: endpoints["SearchUserTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"BlocksIDs": mockEndpointItem{
		URL:    endpoints["BlocksIDs"].URL,
		Method: endpoints["BlocksIDs"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", `"id1"`, `"id2"`, `"id3"`),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"BlocksBlocking": mockEndpointItem{
		URL:    endpoints["BlocksBlocking"].URL,
		Method: endpoints["BlocksBlocking"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"BlocksCreate": mockEndpointItem{
		URL:    endpoints["BlocksCreate"].URL,
		Method: endpoints["BlocksCreate"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"BlocksExists": mockEndpointItem{
		URL:    endpoints["BlocksExists"].URL,
		Method: endpoints["BlocksExists"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"BlocksDestroy": mockEndpointItem{
		URL:    endpoints["BlocksDestroy"].URL,
		Method: endpoints["BlocksDestroy"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"UsersTagged": mockEndpointItem{
		URL:    endpoints["UsersTagged"].URL,
		Method: endpoints["UsersTagged"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"UsersShow": mockEndpointItem{
		URL:    endpoints["UsersShow"].URL,
		Method: endpoints["UsersShow"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"UsersTagList": mockEndpointItem{
		URL:    endpoints["UsersTagList"].URL,
		Method: endpoints["UsersTagList"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", `"tag1"`, `"tag2"`, `"tag3"`),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"UsersFollowers": mockEndpointItem{
		URL:    endpoints["UsersFollowers"].URL,
		Method: endpoints["UsersFollowers"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"UsersRecommendation": mockEndpointItem{
		URL:    endpoints["UsersRecommendation"].URL,
		Method: endpoints["UsersRecommendation"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"UsersCancelRecommendation": mockEndpointItem{
		URL:    endpoints["UsersCancelRecommendation"].URL,
		Method: endpoints["UsersCancelRecommendation"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"UsersFriends": mockEndpointItem{
		URL:    endpoints["UsersFriends"].URL,
		Method: endpoints["UsersFriends"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"AccountVerifyCredentials": mockEndpointItem{
		URL:    endpoints["AccountVerifyCredentials"].URL,
		Method: endpoints["AccountVerifyCredentials"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"AccountUpdateProfileImage": mockEndpointItem{
		URL:    endpoints["AccountUpdateProfileImage"].URL,
		Method: endpoints["AccountUpdateProfileImage"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"AccountRateLimitStatus": mockEndpointItem{
		URL:    endpoints["AccountRateLimitStatus"].URL,
		Method: endpoints["AccountRateLimitStatus"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockAccountRateLimit,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"AccountUpdateProfile": mockEndpointItem{
		URL:    endpoints["AccountUpdateProfile"].URL,
		Method: endpoints["AccountUpdateProfile"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"AccountNotification": mockEndpointItem{
		URL:    endpoints["AccountNotification"].URL,
		Method: endpoints["AccountNotification"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockAccountNotificationStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"AccountUpdateNotifyNum": mockEndpointItem{
		URL:    endpoints["AccountUpdateNotifyNum"].URL,
		Method: endpoints["AccountUpdateNotifyNum"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockAccountNotifyNum,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"AccountNotifyNum": mockEndpointItem{
		URL:    endpoints["AccountNotifyNum"].URL,
		Method: endpoints["AccountNotifyNum"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockAccountNotifyNum,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"SavedSearchesCreate": mockEndpointItem{
		URL:    endpoints["SavedSearchesCreate"].URL,
		Method: endpoints["SavedSearchesCreate"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockSavedSearch,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"SavedSearchesDestroy": mockEndpointItem{
		URL:    endpoints["SavedSearchesDestroy"].URL,
		Method: endpoints["SavedSearchesDestroy"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockSavedSearch,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"SavedSearchesShow": mockEndpointItem{
		URL:    endpoints["SavedSearchesShow"].URL,
		Method: endpoints["SavedSearchesShow"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockSavedSearch,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"SavedSearchesList": mockEndpointItem{
		URL:    endpoints["SavedSearchesList"].URL,
		Method: endpoints["SavedSearchesList"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockSavedSearch, mockSavedSearch, mockSavedSearch),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"PhotosUserTimeline": mockEndpointItem{
		URL:    endpoints["PhotosUserTimeline"].URL,
		Method: endpoints["PhotosUserTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"PhotosUpload": mockEndpointItem{
		URL:    endpoints["PhotosUpload"].URL,
		Method: endpoints["PhotosUpload"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"TrendsList": mockEndpointItem{
		URL:    endpoints["TrendsList"].URL,
		Method: endpoints["TrendsList"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockTrends,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"FollowersIDs": mockEndpointItem{
		URL:    endpoints["FollowersIDs"].URL,
		Method: endpoints["FollowersIDs"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", `"id1"`, `"id2"`, `"id3"`),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"FavoritesDestroy": mockEndpointItem{
		URL:    endpoints["FavoritesDestroy"].URL,
		Method: endpoints["FavoritesDestroy"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"Favorites": mockEndpointItem{
		URL:    endpoints["Favorites"].URL,
		Method: endpoints["Favorites"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FavoritesCreate": mockEndpointItem{
		URL:    endpoints["FavoritesCreate"].URL,
		Method: endpoints["FavoritesCreate"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"FriendshipsCreate": mockEndpointItem{
		URL:    endpoints["FriendshipsCreate"].URL,
		Method: endpoints["FriendshipsCreate"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FriendshipsDestroy": mockEndpointItem{
		URL:    endpoints["FriendshipsDestroy"].URL,
		Method: endpoints["FriendshipsDestroy"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FriendshipsRequests": mockEndpointItem{
		URL:    endpoints["FriendshipsRequests"].URL,
		Method: endpoints["FriendshipsRequests"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FriendshipsDeny": mockEndpointItem{
		URL:    endpoints["FriendshipsDeny"].URL,
		Method: endpoints["FriendshipsDeny"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FriendshipsExists": mockEndpointItem{
		URL:    endpoints["FriendshipsExists"].URL,
		Method: endpoints["FriendshipsExists"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			// this is exactly how the server responds, yes a string
			Body: "true",
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FriendshipsAccept": mockEndpointItem{
		URL:    endpoints["FriendshipsAccept"].URL,
		Method: endpoints["FriendshipsAccept"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockUser,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"FriendshipsShow": mockEndpointItem{
		URL:    endpoints["FriendshipsShow"].URL,
		Method: endpoints["FriendshipsShow"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockFriendship,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"FriendsIDs": mockEndpointItem{
		URL:    endpoints["FriendsIDs"].URL,
		Method: endpoints["FriendsIDs"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", `"id1"`, `"id2"`, `"id3"`),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"StatusesDestroy": mockEndpointItem{
		URL:    endpoints["StatusesDestroy"].URL,
		Method: endpoints["StatusesDestroy"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesHomeTimeline": mockEndpointItem{
		URL:    endpoints["StatusesHomeTimeline"].URL,
		Method: endpoints["StatusesHomeTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesPublicTimeline": mockEndpointItem{
		URL:    endpoints["StatusesPublicTimeline"].URL,
		Method: endpoints["StatusesPublicTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesReplies": mockEndpointItem{
		URL:    endpoints["StatusesReplies"].URL,
		Method: endpoints["StatusesReplies"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesFollowers": mockEndpointItem{
		URL:    endpoints["StatusesFollowers"].URL,
		Method: endpoints["StatusesFollowers"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesUpdate": mockEndpointItem{
		URL:    endpoints["StatusesUpdate"].URL,
		Method: endpoints["StatusesUpdate"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesUserTimeline": mockEndpointItem{
		URL:    endpoints["StatusesUserTimeline"].URL,
		Method: endpoints["StatusesUserTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesFriends": mockEndpointItem{
		URL:    endpoints["StatusesFriends"].URL,
		Method: endpoints["StatusesFriends"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockUser, mockUser, mockUser),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesContextTimeline": mockEndpointItem{
		URL:    endpoints["StatusesContextTimeline"].URL,
		Method: endpoints["StatusesContextTimeline"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesMentions": mockEndpointItem{
		URL:    endpoints["StatusesMentions"].URL,
		Method: endpoints["StatusesMentions"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockStatus, mockStatus, mockStatus),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"StatusesShow": mockEndpointItem{
		URL:    endpoints["StatusesShow"].URL,
		Method: endpoints["StatusesShow"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockStatus,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},

	"DirectMessagesDestroy": mockEndpointItem{
		URL:    endpoints["DirectMessagesDestroy"].URL,
		Method: endpoints["DirectMessagesDestroy"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockDirectMessage,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"DirectMessagesConversation": mockEndpointItem{
		URL:    endpoints["DirectMessagesConversation"].URL,
		Method: endpoints["DirectMessagesConversation"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockDirectMessage, mockDirectMessage, mockDirectMessage),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"DirectMessagesNew": mockEndpointItem{
		URL:    endpoints["DirectMessagesNew"].URL,
		Method: endpoints["DirectMessagesNew"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        mockDirectMessage,
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"DirectMessagesConversationList": mockEndpointItem{
		URL:    endpoints["DirectMessagesConversationList"].URL,
		Method: endpoints["DirectMessagesConversationList"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockDirectMessageConversationItem, mockDirectMessageConversationItem, mockDirectMessageConversationItem),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"DirectMessagesInbox": mockEndpointItem{
		URL:    endpoints["DirectMessagesInbox"].URL,
		Method: endpoints["DirectMessagesInbox"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockDirectMessage, mockDirectMessage, mockDirectMessage),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
	"DirectMessagesSent": mockEndpointItem{
		URL:    endpoints["DirectMessagesSent"].URL,
		Method: endpoints["DirectMessagesSent"].Method,
		Result200: mockResult{
			Code:        200,
			ContentType: contentTypeJSON,
			Body:        fmt.Sprintf("[%s,%s,%s]", mockDirectMessage, mockDirectMessage, mockDirectMessage),
		},
		Result400: mockResult{
			Code:        400,
			ContentType: contentTypeJSON,
			Body:        mockClientError,
		},
		Result500: mockResult{
			Code:        500,
			ContentType: contentTypeHTML,
		},
		ResultChaos: mockResult{
			Code:        200,
			ContentType: contentTypeHTML,
		},
	},
}

func init() {
	httpmock.RegisterResponder("GET", requestTokenURL, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(200, fmt.Sprintf(`oauth_token=%s&oauth_token_secret=%s`, mockRequestToken, mockRequestSecret)), nil
	})

	httpmock.RegisterResponder("GET", accessTokenURL, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(200, fmt.Sprintf(`oauth_token=%s&oauth_token_secret=%s`, mockAccessToken, mockAccessSecret)), nil
	})

	httpmock.Activate()
}
