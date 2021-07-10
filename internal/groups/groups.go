package groups

type GroupList struct {
	Count  int64   `json:"count"`
	Groups []Group `json:"items"`
}

type Group struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	IsClosed     int64  `json:"is_closed"`
	Deactivated  string `json:"deactivated"`
	IsAdmin      int64  `json:"is_admin"`
	AdminLevel   int64  `json:"admin_level"`
	IsMember     int64  `json:"is_member"`
	IsAdvertiser int64  `json:"is_advertiser"`
	InvitedBy    int64  `json:"invited_by"`
	Type         string `json:"type"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
}
