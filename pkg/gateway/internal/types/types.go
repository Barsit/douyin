// Code generated by goctl. DO NOT EDIT.
package types

type FavoriteOptReq struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int64  `form:"action_type" validate:"required,numeric,min=1,max=2"`
}

type FavoriteOptRes struct {
	Status
}

type FollowOptReq struct {
	Token      string `form:"token"`
	FollowId   int64  `form:"to_user_id"`
	ActionType int64  `form:"action_type" validate:"required,numeric,min=1,max=2"`
}

type FollowOptRes struct {
	Status
}

type CommentOptReq struct {
	Token       string `form:"token" validate:"required"`
	VideoId     int64  `form:"video_id" validate:"required,numeric"`
	ActionType  int64  `form:"action_type" validate:"required,numeric,min=1,max=2"`
	CommentText string `form:"comment_text,omitempty, optional"`
	CommentId   int64  `form:"comment_id,omitempty, optional"`
}

type Comment struct {
	CommentId  int64  `json:"id" copyier:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateTime string `json:"create_date"`
}

type CommentOptRes struct {
	Status
	Comment *Comment `json:"comment,omitempty"`
}

type FavoriteListReq struct {
	IdWithTokenReq
}

type FavoriteListRes struct {
	Status
	FavoriteList []*PubVideo `json:"video_list,omitempty"`
}

type FollowListReq struct {
	IdWithTokenReq
}

type FollowListRes struct {
	Status
	UserFollowlist []*User `json:"user_list,omitempty"`
}

type FollowerListReq struct {
	IdWithTokenReq
}

type FollowerListRes struct {
	Status
	UserFollowerlist []*User `json:"user_list,omitempty"`
}

type CommentListReq struct {
	Token   string `form:"token" validate:"required"`
	VideoId int64  `form:"video_id" validate:"required"`
}

type CommentListRes struct {
	Status
	CommentList []*Comment `json:"comment_list,omitempty"`
}

type Status struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg,omitempty"`
}

type UserReq struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type IdWithTokenReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type IdWithTokenRes struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserRegisterReq struct {
	UserReq
}

type UserRegisterRes struct {
	Status
	IdWithTokenRes
}

type UserLoginReq struct {
	UserReq
}

type UserLoginRes struct {
	Status
	IdWithTokenRes
}

type User struct {
	UserId          int64  `json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type UserInfoReq struct {
	IdWithTokenReq
}

type UserInfoRes struct {
	Status
	User *User `json:"user,omitempty"`
}

type PubVideoReq struct {
	Token string `form:"token"`
	Data  []byte `form:"data"`
	Title string `form:"title"`
}

type PubVideoRes struct {
	Status
}

type GetPubVideoListReq struct {
	Token  string `form:"token"`
	UserId int64  `form:"user_id"`
}

type PubVideo struct {
	Id            int64  `json:"id"`
	User          User   `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type GetPubVideoListRes struct {
	Status
	VideoPubList []*PubVideo `json:"video_list,omitempty"`
}

type FeedVideoListReq struct {
	Token    string `form:"token,optional"`
	LastTime int64  `form:"last_time,optional"`
}

type FeedVideoListRes struct {
	Status
	NextTime  int64       `json:"next_time,omitempty"`
	VideoList []*PubVideo `json:"video_list,omitempty"`
}

type MessageReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int64  `form:"action_type"`
	Content    string `form:"content"`
}

type MessageRes struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg,omitempty"`
}

type MessageListReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type MessageListRes struct {
	Status
	VideoPubList []*Message `json:"message_list,omitempty"`
}

type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}
