// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Video struct {
	VideoID        string `json:"videoId"`
	VideoURL       string `json:"videoUrl"`
	VideoTitle     string `json:"videoTitle"`
	VideoDesc      string `json:"videoDesc"`
	VideoCat       string `json:"videoCat"`
	VideoThumb     string `json:"videoThumb"`
	PlaylistID     int    `json:"playlistId"`
	LikeCount      int    `json:"likeCount"`
	DislikeCount   int    `json:"dislikeCount"`
	AgeRestriction bool   `json:"AgeRestriction"`
	Visibility     string `json:"visibility"`
	Location       string `json:"location"`
	Status         string `json:"status"`
	Premium        bool   `json:"premium"`
}

type Comment struct {
	CommentID string `json:"commentId"`
	UserID    int    `json:"userId"`
	VideoID   int    `json:"videoId"`
	Comment   string `json:"comment"`
	Like      int    `json:"like"`
	Dislike   int    `json:"dislike"`
}

type NewVideo struct {
	VideoURL       string `json:"videoUrl"`
	VideoTitle     string `json:"videoTitle"`
	VideoDesc      string `json:"videoDesc"`
	VideoCat       string `json:"videoCat"`
	VideoThumb     string `json:"videoThumb"`
	PlaylistID     int    `json:"playlistId"`
	LikeCount      int    `json:"likeCount"`
	DislikeCount   int    `json:"dislikeCount"`
	AgeRestriction bool   `json:"AgeRestriction"`
	Visibility     string `json:"visibility"`
	Location       string `json:"location"`
	Status         string `json:"status"`
	Premium        bool   `json:"premium"`
	ViewCount      int    `json:"viewCount"`
}

type Playlist struct {
	PlaylistID   string  `json:"playlistId"`
	PlaylistName string  `json:"playlistName"`
	UserID       int     `json:"userId"`
	Visibility   *string `json:"visibility"`
}

type Reply struct {
	ReplyID   string `json:"replyId"`
	CommentID int    `json:"commentId"`
	UserID    int    `json:"userId"`
	Reply     string `json:"reply"`
}

type User struct {
	UserID          string `json:"userId"`
	Member          bool   `json:"member"`
	ImgURL          string `json:"imgUrl"`
	Email           string `json:"email"`
	UserName        string `json:"userName"`
	SubscriberCount int    `json:"subscriberCount"`
}
