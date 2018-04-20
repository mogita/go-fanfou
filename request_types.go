package fanfou

// ReqParams defines the params for requests
type ReqParams struct {
	ID                string `json:"id,omitempty"`
	SinceID           string `json:"since_id,omitempty"`
	MaxID             string `json:"max_id,omitempty"`
	Count             string `json:"count,omitempty"`
	Mode              string `json:"mode,omitempty"`
	Format            string `json:"format,omitempty"`
	CallBack          string `json:"callback,omitempty"`
	Q                 string `json:"q,omitempty"`
	Page              string `json:"page,omitempty"`
	Tag               string `json:"tag,omitempty"`
	Image             string `json:"image,omitempty"`
	URL               string `json:"url,omitempty"`
	Location          string `json:"location,omitempty"`
	Description       string `json:"description,omitempty"`
	Name              string `json:"name,omitempty"`
	Email             string `json:"email,omitempty"`
	NotifyNum         string `json:"notify_num,omitempty"`
	Query             string `json:"query,omitempty"`
	Photo             string `json:"photo,omitempty"`
	Status            string `json:"status,omitempty"`
	Source            string `json:"source,omitempty"`
	UserA             string `json:"user_a,omitempty"`
	UserB             string `json:"user_b,omitempty"`
	SourceLoginName   string `json:"source_login_name,omitempty"`
	TargetLoginName   string `json:"target_login_name,omitempty"`
	SourceID          string `json:"source_id,omitempty"`
	TargetID          string `json:"target_id,omitempty"`
	InReplyToStatusID string `json:"in_reply_to_status_id,omitempty"`
	InReplyToUserID   string `json:"in_reply_to_user_id,omitempty"`
	RepostStatusID    string `json:"repost_status_id,omitempty"`
	User              string `json:"user,omitempty"`
	Text              string `json:"text,omitempty"`
	InReplyToID       string `json:"in_reply_to_id,omitempty"`
}
