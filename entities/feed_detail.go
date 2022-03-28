package ent

type FeedDetail struct {
	RawData
	ErrorMsg
	Data struct {
		Id                int           `json:"id"`
		Type              int           `json:"type"`
		Fid               int           `json:"fid"`
		Forwardid         string        `json:"forwardid"`
		SourceId          string        `json:"source_id"`
		Uid               int           `json:"uid"`
		Username          string        `json:"username"`
		TopicType         int           `json:"ttype"`
		TopicID           int64         `json:"tid"`
		TopicTitle        string        `json:"ttitle"`
		TopicPic          string        `json:"tpic"`
		TopicURL          string        `json:"turl"`
		TopicInfo         string        `json:"tinfo"`
		MessageTitle      string        `json:"message_title"`
		MessageKeywords   string        `json:"message_keywords"`
		MessageCover      string        `json:"message_cover"`
		Message           string        `json:"message"`
		Pic               string        `json:"pic"`
		MessageLength     int           `json:"message_length"`
		Issummary         int           `json:"issummary"`
		Istag             int           `json:"istag"`
		IsHtmlArticle     int           `json:"is_html_article"`
		Tags              string        `json:"tags"`
		Label             string        `json:"label"`
		UserTags          string        `json:"user_tags"`
		MediaType         int           `json:"media_type"`
		MediaPic          string        `json:"media_pic"`
		MediaUrl          string        `json:"media_url"`
		ExtraType         int           `json:"extra_type"`
		ExtraKey          string        `json:"extra_key"`
		ExtraTitle        string        `json:"extra_title"`
		ExtraUrl          string        `json:"extra_url"`
		ExtraPic          string        `json:"extra_pic"`
		ExtraInfo         string        `json:"extra_info"`
		ExtraStatus       int           `json:"extra_status"`
		Location          string        `json:"location"`
		FromID            int           `json:"fromid"`
		FromName          string        `json:"fromname"`
		LikeNum           int           `json:"likenum"`
		BuryNum           int           `json:"burynum"`
		CommentNum        int           `json:"commentnum"`
		ReplyNum          int           `json:"replynum"`
		ForwardNum        int           `json:"forwardnum"`
		ReportNum         int           `json:"reportnum"`
		RelatedNum        int           `json:"relatednum"`
		FavoriteNum       int           `json:"favnum"`
		ShareNum          int           `json:"share_num"`
		CommentBlockNum   int           `json:"comment_block_num"`
		QuestionAnswerNum int           `json:"question_answer_num"`
		QuestionFollowNum int           `json:"question_follow_num"`
		HitNum            int           `json:"hitnum"`
		ViewNum           int           `json:"viewnum"`
		FeedScore         int           `json:"feed_score"`
		RankScore         int           `json:"rank_score"`
		VoteScore         int           `json:"vote_score"`
		AtCount           int           `json:"at_count"`
		UrlCount          int           `json:"url_count"`
		TagCount          int           `json:"tag_count"`
		ChangeCount       int           `json:"change_count"`
		Recommend         int           `json:"recommend"`
		IsAnonymous       int           `json:"is_anonymous"`
		IsHidden          int           `json:"is_hidden"`
		IsHeadline        int           `json:"is_headline"`
		DisallowReply     int           `json:"disallow_reply"`
		Status            int           `json:"status"`
		MessageStatus     int           `json:"message_status"`
		BlockStatus       int           `json:"block_status"`
		Dateline          int           `json:"dateline"`
		Lastupdate        int           `json:"lastupdate"`
		LastChangeTime    int           `json:"last_change_time"`
		DeviceTitle       string        `json:"device_title"`
		DeviceName        string        `json:"device_name"`
		DeviceRom         string        `json:"device_rom"`
		DeviceBuild       string        `json:"device_build"`
		RecentReplyIds    string        `json:"recent_reply_ids"`
		RecentHotReplyIds string        `json:"recent_hot_reply_ids"`
		RecentLikeList    string        `json:"recent_like_list"`
		RelatedDyhIds     string        `json:"related_dyh_ids"`
		PostSignature     string        `json:"post_signature"`
		MessageSignature  string        `json:"message_signature"`
		FetchType         string        `json:"fetchType"`
		EntityId          int           `json:"entityId"`
		AvatarFetchType   string        `json:"avatarFetchType"`
		UserAvatar        string        `json:"userAvatar"`
		EntityTemplate    string        `json:"entityTemplate"`
		EntityType        string        `json:"entityType"`
		Url               string        `json:"url"`
		FeedType          string        `json:"feedType"`
		FeedTypeName      string        `json:"feedTypeName"`
		TurlTarget        string        `json:"turlTarget"`
		IsModified        int           `json:"isModified"`
		EnableModify      int           `json:"enableModify"`
		Info              string        `json:"info"`
		InfoHtml          string        `json:"infoHtml"`
		Title             string        `json:"title"`
		PicArr            []string      `json:"picArr"`
		DeviceTitleUrl    string        `json:"device_title_url"`
		Relateddata       []interface{} `json:"relateddata"`
		MediaInfo         string        `json:"media_info"`
		ShareUrl          string        `json:"shareUrl"`
		ExtraFromApi      string        `json:"extra_fromApi"`
		SourceFeed        interface{}   `json:"sourceFeed"`
		ForwardSourceType string        `json:"forwardSourceType"`
		CanDisallowReply  int           `json:"canDisallowReply"`
		DisallowRepost    int           `json:"disallow_repost"`
		LongLocation      string        `json:"long_location"`
		IsWhiteFeed       int           `json:"is_white_feed"`
		EditorTitle       string        `json:"editor_title"`
		IncludeGoodsIds   []string      `json:"include_goods_ids"`
		TopReplyIds       []interface{} `json:"top_reply_ids"`
		MessageBrief      string        `json:"message_brief"`
		MessageRawOutput  string        `json:"message_raw_output"`
		ReplyRows         []interface{} `json:"replyRows"`
		ReplyRowsCount    int           `json:"replyRowsCount"`
		ReplyRowsMore     int           `json:"replyRowsMore"`
		UserInfo          UserInfo      `json:"userInfo"`
		UserAction        struct {
			Like            int `json:"like"`
			Favorite        int `json:"favorite"`
			Follow          int `json:"follow"`
			Collect         int `json:"collect"`
			FollowAuthor    int `json:"followAuthor"`
			AuthorFollowYou int `json:"authorFollowYou"`
		} `json:"userAction"`
		HotReplyRows []struct {
			ReplyRow
			ReplyRows      []ReplyRow    `json:"replyRows,omitempty"`
			ReplyRowsCount int           `json:"replyRowsCount,omitempty"`
			ReplyRowsMore  int           `json:"replyRowsMore,omitempty"`
			Title          string        `json:"title,omitempty"`
			Url            string        `json:"url,omitempty"`
			Entities       []interface{} `json:"entities,omitempty"`
			EntityFixed    int           `json:"entityFixed,omitempty"`
			ExtraDataArr   struct {
				CardId       string `json:"cardId"`
				CardPageName string `json:"cardPageName"`
				SponsorType  string `json:"sponsorType"`
				CardStatName string `json:"cardStatName"`
			} `json:"extraDataArr,omitempty"`
			ExtraData string `json:"extraData,omitempty"`
		} `json:"hotReplyRows"`
		TopReplyRows []interface{} `json:"topReplyRows"`
	} `json:"data"`
}

type UserInfo struct {
	Uid                 int    `json:"uid"`
	Username            string `json:"username"`
	Admintype           int    `json:"admintype"`
	Groupid             int    `json:"groupid"`
	Usergroupid         int    `json:"usergroupid"`
	Level               int    `json:"level"`
	Experience          int    `json:"experience"`
	Status              int    `json:"status"`
	BlockStatus         int    `json:"block_status"`
	Usernamestatus      int    `json:"usernamestatus"`
	Avatarstatus        int    `json:"avatarstatus"`
	AvatarCoverStatus   int    `json:"avatar_cover_status"`
	Regdate             int    `json:"regdate"`
	Logintime           int    `json:"logintime"`
	VerifyTitle         string `json:"verify_title"`
	VerifyStatus        int    `json:"verify_status"`
	UserType            int    `json:"user_type"`
	VerifyShowType      int    `json:"verify_show_type"`
	AvatarPluginStatus  int    `json:"avatar_plugin_status"`
	FetchType           string `json:"fetchType"`
	EntityType          string `json:"entityType"`
	EntityId            int    `json:"entityId"`
	DisplayUsername     string `json:"displayUsername"`
	Url                 string `json:"url"`
	UserAvatar          string `json:"userAvatar"`
	UserSmallAvatar     string `json:"userSmallAvatar"`
	UserBigAvatar       string `json:"userBigAvatar"`
	Cover               string `json:"cover"`
	VerifyIcon          string `json:"verify_icon"`
	VerifyLabel         string `json:"verify_label"`
	IsDeveloper         int    `json:"isDeveloper"`
	NextLevelExperience int    `json:"next_level_experience"`
	NextLevelPercentage string `json:"next_level_percentage"`
	LevelTodayMessage   string `json:"level_today_message"`
	LevelDetailUrl      string `json:"level_detail_url"`
	AvatarPluginUrl     string `json:"avatar_plugin_url"`
	FeedPluginUrl       string `json:"feed_plugin_url"`
	FeedPluginOpenUrl   string `json:"feed_plugin_open_url"`
}

type ReplyRow struct {
	Id              int         `json:"id,omitempty"`
	Ftype           int         `json:"ftype,omitempty"`
	Fid             int         `json:"fid,omitempty"`
	Rid             int         `json:"rid,omitempty"`
	Rrid            int         `json:"rrid,omitempty"`
	Uid             int         `json:"uid,omitempty"`
	Username        string      `json:"username,omitempty"`
	Ruid            int         `json:"ruid,omitempty"`
	Rusername       string      `json:"rusername,omitempty"`
	Pic             string      `json:"pic"`
	Message         string      `json:"message,omitempty"`
	Replynum        int         `json:"replynum,omitempty"`
	Likenum         int         `json:"likenum,omitempty"`
	Burynum         int         `json:"burynum,omitempty"`
	Reportnum       int         `json:"reportnum,omitempty"`
	RankScore       int         `json:"rank_score,omitempty"`
	Dateline        int         `json:"dateline,omitempty"`
	Lastupdate      int         `json:"lastupdate"`
	IsFolded        int         `json:"is_folded,omitempty"`
	Status          int         `json:"status,omitempty"`
	MessageStatus   int         `json:"message_status,omitempty"`
	BlockStatus     int         `json:"block_status,omitempty"`
	RecentReplyIds  string      `json:"recent_reply_ids,omitempty"`
	FeedUid         int         `json:"feedUid,omitempty"`
	FetchType       string      `json:"fetchType,omitempty"`
	EntityId        interface{} `json:"entityId"`
	AvatarFetchType string      `json:"avatarFetchType,omitempty"`
	UserAvatar      string      `json:"userAvatar,omitempty"`
	EntityTemplate  string      `json:"entityTemplate"`
	EntityType      string      `json:"entityType"`
	InfoHtml        string      `json:"infoHtml,omitempty"`
	IsFeedAuthor    int         `json:"isFeedAuthor,omitempty"`
	PicArr          []string    `json:"picArr,omitempty"`
	UserInfo        UserInfo    `json:"userInfo,omitempty"`
}
