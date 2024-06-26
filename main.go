package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Nms struct {
	NmIds []int `json:"nm_ids"`
}

type InstaInfo struct {
	NmId     int
	ReelInfo Reel
}

type Reel struct {
	ItemCode string
	ItemLink string
	Username string
	FullName string
	
	AuthorLink    string
	LikesCount    int
	CommentsCount int
	RepostCount   int
	Date          time.Time
}

type Response struct {
	Data    Data   `json:"data"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Data struct {
	Modules    []Module `json:"reels_serp_modules"`
	HasMore    bool     `json:"has_more"`
	ReelsMaxId string   `json:"reels_max_id"`
	PageIndex  int      `json:"page_index"`
	RankToken  string   `json:"rank_token"`
	Status     string   `json:"status"`
}

type Module struct {
	Type  string `json:"module_type"`
	Clips []Clip `json:"clips"`
}

type Clip struct {
	Media Media `json:"media"`
}

type Media struct {
	TakenAt                                              int64                  `json:"taken_at"`
	PK                                                   int64                  `json:"pk"`
	ID                                                   string                 `json:"id"`
	Fbid                                                 int64                  `json:"fbid"`
	DeviceTimestamp                                      int64                  `json:"device_timestamp"`
	CaptionIsEdited                                      bool                   `json:"caption_is_edited"`
	StrongID                                             string                 `json:"strong_id__"`
	DeletedReason                                        int                    `json:"deleted_reason"`
	HasSharedToFb                                        int                    `json:"has_shared_to_fb"`
	ShouldRequestAds                                     bool                   `json:"should_request_ads"`
	HasDelayedMetadata                                   bool                   `json:"has_delayed_metadata"`
	ViewStateItemType                                    int                    `json:"view_state_item_type"`
	LoggingInfoToken                                     string                 `json:"logging_info_token"`
	MezqlToken                                           string                 `json:"mezql_token"`
	ShareCountDisabled                                   bool                   `json:"share_count_disabled"`
	HideViewAllCommentEntrypoint                         bool                   `json:"hide_view_all_comment_entrypoint"`
	IsVisualReplyCommenterNoticeEnabled                  bool                   `json:"is_visual_reply_commenter_notice_enabled"`
	LikeAndViewCountsDisabled                            bool                   `json:"like_and_view_counts_disabled"`
	IsPostLiveClipsMedia                                 bool                   `json:"is_post_live_clips_media"`
	ClipsDeliveryParameters                              []interface{}          `json:"clips_delivery_parameters"`
	IsReshareOfTextPostAppMediaInIG                      bool                   `json:"is_reshare_of_text_post_app_media_in_ig"`
	HasPrivatelyLiked                                    bool                   `json:"has_privately_liked"`
	FilterType                                           int                    `json:"filter_type"`
	ClientCacheKey                                       string                 `json:"client_cache_key"`
	IntegrityReviewDecision                              string                 `json:"integrity_review_decision"`
	CanSeeInsightsAsBrand                                bool                   `json:"can_see_insights_as_brand"`
	MediaType                                            int                    `json:"media_type"`
	Code                                                 string                 `json:"code"`
	Caption                                              Caption                `json:"caption"`
	FundraiserTag                                        FundraiserTag          `json:"fundraiser_tag"`
	SharingFrictionInfo                                  SharingFrictionInfo    `json:"sharing_friction_info"`
	PlayCount                                            int                    `json:"play_count"`
	OriginalMediaHasVisualReplyMedia                     bool                   `json:"original_media_has_visual_reply_media"`
	CreatorViewerInsights                                []interface{}          `json:"creator_viewer_insights"`
	FbUserTags                                           FbUserTags             `json:"fb_user_tags"`
	CoauthorProducers                                    []interface{}          `json:"coauthor_producers"`
	InvitedCoauthorProducers                             []interface{}          `json:"invited_coauthor_producers"`
	IsInProfileGrid                                      bool                   `json:"is_in_profile_grid"`
	ProfileGridControlEnabled                            bool                   `json:"profile_grid_control_enabled"`
	MediaCroppingInfo                                    MediaCroppingInfo      `json:"media_cropping_info"`
	User                                                 User                   `json:"user"`
	ImageVersions2                                       ImageVersions2         `json:"image_versions2"`
	OriginalWidth                                        int                    `json:"original_width"`
	OriginalHeight                                       int                    `json:"original_height"`
	IsArtistPick                                         bool                   `json:"is_artist_pick"`
	EnableMediaNotesProduction                           bool                   `json:"enable_media_notes_production"`
	ProductType                                          string                 `json:"product_type"`
	IsPaidPartnership                                    bool                   `json:"is_paid_partnership"`
	InventorySource                                      string                 `json:"inventory_source"`
	MusicMetadata                                        interface{}            `json:"music_metadata"`
	SocialContext                                        []interface{}          `json:"social_context"`
	OrganicTrackingToken                                 string                 `json:"organic_tracking_token"`
	IsThirdPartyDownloadsEligible                        bool                   `json:"is_third_party_downloads_eligible"`
	IgMediaSharingDisabled                               bool                   `json:"ig_media_sharing_disabled"`
	AreRemixesCrosspostable                              bool                   `json:"are_remixes_crosspostable"`
	BoostUnavailableIdentifier                           interface{}            `json:"boost_unavailable_identifier"`
	BoostUnavailableReason                               interface{}            `json:"boost_unavailable_reason"`
	SubscribeCtaVisible                                  bool                   `json:"subscribe_cta_visible"`
	IsAutoCreated                                        bool                   `json:"is_auto_created"`
	IsCutoutStickerAllowed                               bool                   `json:"is_cutout_sticker_allowed"`
	IsReuseAllowed                                       bool                   `json:"is_reuse_allowed"`
	EnableWaist                                          bool                   `json:"enable_waist"`
	Owner                                                User                   `json:"owner"`
	FbAggregatedLikeCount                                int                    `json:"fb_aggregated_like_count"`
	FbAggregatedCommentCount                             int                    `json:"fb_aggregated_comment_count"`
	IsTaggedMediaSharedToViewerProfileGrid               bool                   `json:"is_tagged_media_shared_to_viewer_profile_grid"`
	ShouldShowAuthorPogForTaggedMediaSharedToProfileGrid bool                   `json:"should_show_author_pog_for_tagged_media_shared_to_profile_grid"`
	IsEligibleForMediaNoteRecsNux                        bool                   `json:"is_eligible_for_media_note_recs_nux"`
	CommentCount                                         int                    `json:"comment_count"`
	IsCommentsGifComposerEnabled                         bool                   `json:"is_comments_gif_composer_enabled"`
	CommentInformTreatment                               CommentInformTreatment `json:"comment_inform_treatment"`
	HasLiked                                             bool                   `json:"has_liked"`
	LikeCount                                            int                    `json:"like_count"`
	TopLikers                                            []string               `json:"top_likers"`
	VideoSubtitlesConfidence                             float64                `json:"video_subtitles_confidence"`
	VideoSubtitlesLocale                                 string                 `json:"video_subtitles_locale"`
	IsDashEligible                                       int                    `json:"is_dash_eligible"`
	VideoDashManifest                                    string                 `json:"video_dash_manifest"`
	VideoCodec                                           string                 `json:"video_codec"`
	NumberOfQualities                                    int                    `json:"number_of_qualities"`
	VideoVersions                                        []VideoVersion         `json:"video_versions"`
	VideoDuration                                        float64                `json:"video_duration"`
	HasAudio                                             bool                   `json:"has_audio"`
	ClipsTabPinnedUserIds                                []interface{}          `json:"clips_tab_pinned_user_ids"`
	ClipsMetadata                                        ClipsMetadata          `json:"clips_metadata"`
	CanViewerSave                                        bool                   `json:"can_viewer_save"`
	CanViewerReshare                                     bool                   `json:"can_viewer_reshare"`
	ReshareCount                                         int                    `json:"reshare_count"`
	ShopRoutingUserID                                    interface{}            `json:"shop_routing_user_id"`
	IsOrganicProductTaggingEligible                      bool                   `json:"is_organic_product_tagging_eligible"`
	ProductSuggestions                                   []interface{}          `json:"product_suggestions"`
	CanViewMorePreviewComments                           bool                   `json:"can_view_more_preview_comments"`
	IsOpenToPublicSubmission                             bool                   `json:"is_open_to_public_submission"`
}

type Caption struct {
	BitFlags           int    `json:"bit_flags"`
	CreatedAt          int64  `json:"created_at"`
	CreatedAtUtc       int64  `json:"created_at_utc"`
	DidReportAsSpam    bool   `json:"did_report_as_spam"`
	IsRankedComment    bool   `json:"is_ranked_comment"`
	PK                 string `json:"pk"`
	ShareEnabled       bool   `json:"share_enabled"`
	ContentType        string `json:"content_type"`
	MediaID            int64  `json:"media_id"`
	Status             string `json:"status"`
	Type               int    `json:"type"`
	UserID             int64  `json:"user_id"`
	HasTranslation     bool   `json:"has_translation"`
	Text               string `json:"text"`
	User               User   `json:"user"`
	IsCovered          bool   `json:"is_covered"`
	PrivateReplyStatus int    `json:"private_reply_status"`
}

type FundraiserTag struct {
	HasStandaloneFundraiser bool `json:"has_standalone_fundraiser"`
}

type SharingFrictionInfo struct {
	ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
	BloksAppURL               interface{} `json:"bloks_app_url"`
	SharingFrictionPayload    interface{} `json:"sharing_friction_payload"`
}

type FbUserTags struct {
	In []interface{} `json:"in"`
}

type MediaCroppingInfo struct {
	SquareCrop SquareCrop `json:"square_crop"`
}

type SquareCrop struct {
	CropLeft   float64 `json:"crop_left"`
	CropRight  float64 `json:"crop_right"`
	CropTop    float64 `json:"crop_top"`
	CropBottom float64 `json:"crop_bottom"`
}

type User struct {
	FbidV2                            int64            `json:"fbid_v2"`
	FeedPostReshareDisabled           bool             `json:"feed_post_reshare_disabled"`
	FullName                          string           `json:"full_name"`
	ID                                string           `json:"id"`
	IsPrivate                         bool             `json:"is_private"`
	IsUnpublished                     bool             `json:"is_unpublished"`
	PK                                int64            `json:"pk"`
	PKID                              string           `json:"pk_id"`
	ShowAccountTransparencyDetails    bool             `json:"show_account_transparency_details"`
	StrongID                          string           `json:"strong_id__"`
	ThirdPartyDownloadsEnabled        int              `json:"third_party_downloads_enabled"`
	AccountBadges                     []interface{}    `json:"account_badges"`
	FanClubInfo                       FanClubInfo      `json:"fan_club_info"`
	FriendshipStatus                  FriendshipStatus `json:"friendship_status"`
	HasAnonymousProfilePicture        bool             `json:"has_anonymous_profile_picture"`
	IsFavorite                        bool             `json:"is_favorite"`
	IsVerified                        bool             `json:"is_verified"`
	LatestReelMedia                   int64            `json:"latest_reel_media"`
	ProfilePicID                      string           `json:"profile_pic_id"`
	ProfilePicURL                     string           `json:"profile_pic_url"`
	TransparencyProductEnabled        bool             `json:"transparency_product_enabled"`
	Username                          string           `json:"username"`
	EligibleForTextAppActivationBadge bool             `json:"eligible_for_text_app_activation_badge"`
}

type FanClubInfo struct {
	FanClubID                            interface{} `json:"fan_club_id"`
	FanClubName                          interface{} `json:"fan_club_name"`
	IsFanClubReferralEligible            interface{} `json:"is_fan_club_referral_eligible"`
	FanConsiderationPageRevampEligiblity interface{} `json:"fan_consideration_page_revamp_eligiblity"`
	IsFanClubGiftingEligible             interface{} `json:"is_fan_club_gifting_eligible"`
	SubscriberCount                      interface{} `json:"subscriber_count"`
	ConnectedMemberCount                 interface{} `json:"connected_member_count"`
	AutosaveToExclusiveHighlight         interface{} `json:"autosave_to_exclusive_highlight"`
	HasEnoughSubscribersForSSC           interface{} `json:"has_enough_subscribers_for_ssc"`
}

type FriendshipStatus struct {
	Following      bool `json:"following"`
	IsBestie       bool `json:"is_bestie"`
	IsRestricted   bool `json:"is_restricted"`
	IsFeedFavorite bool `json:"is_feed_favorite"`
}

type ImageVersions2 struct {
	Candidates                        []Candidate                       `json:"candidates"`
	AdditionalCandidates              AdditionalCandidates              `json:"additional_candidates"`
	SmartThumbnailEnabled             bool                              `json:"smart_thumbnail_enabled"`
	ScrubberSpritesheetInfoCandidates ScrubberSpritesheetInfoCandidates `json:"scrubber_spritesheet_info_candidates"`
}

type Candidate struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	URL          string `json:"url"`
	ScansProfile string `json:"scans_profile"`
}

type AdditionalCandidates struct {
	IgtvFirstFrame Candidate   `json:"igtv_first_frame"`
	FirstFrame     Candidate   `json:"first_frame"`
	SmartFrame     interface{} `json:"smart_frame"`
}

type ScrubberSpritesheetInfoCandidates struct {
	Default ScrubberSpritesheetInfo `json:"default"`
}

type ScrubberSpritesheetInfo struct {
	VideoLength                float64  `json:"video_length"`
	ThumbnailWidth             int      `json:"thumbnail_width"`
	ThumbnailHeight            int      `json:"thumbnail_height"`
	ThumbnailDuration          float64  `json:"thumbnail_duration"`
	SpriteURLs                 []string `json:"sprite_urls"`
	ThumbnailsPerRow           int      `json:"thumbnails_per_row"`
	TotalThumbnailNumPerSprite int      `json:"total_thumbnail_num_per_sprite"`
	MaxThumbnailsPerSprite     int      `json:"max_thumbnails_per_sprite"`
	SpriteWidth                int      `json:"sprite_width"`
	SpriteHeight               int      `json:"sprite_height"`
	RenderedWidth              int      `json:"rendered_width"`
	FileSizeKB                 int      `json:"file_size_kb"`
}

type VideoVersion struct {
	Type   int    `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
	ID     string `json:"id"`
}

type ClipsMetadata struct {
	BreakingContentInfo             interface{}             `json:"breaking_content_info"`
	BreakingCreatorInfo             interface{}             `json:"breaking_creator_info"`
	ChallengeInfo                   interface{}             `json:"challenge_info"`
	ClipsCreationEntryPoint         string                  `json:"clips_creation_entry_point"`
	FeaturedLabel                   interface{}             `json:"featured_label"`
	IsPublicChatWelcomeVideo        bool                    `json:"is_public_chat_welcome_video"`
	NuxInfo                         interface{}             `json:"nux_info"`
	ProfessionalClipsUpsellType     int                     `json:"professional_clips_upsell_type"`
	ReelsOnTheRiseInfo              interface{}             `json:"reels_on_the_rise_info"`
	ShowTips                        interface{}             `json:"show_tips"`
	AchievementsInfo                AchievementsInfo        `json:"achievements_info"`
	AdditionalAudioInfo             AdditionalAudioInfo     `json:"additional_audio_info"`
	AssetRecommendationInfo         interface{}             `json:"asset_recommendation_info"`
	AudioRankingInfo                AudioRankingInfo        `json:"audio_ranking_info"`
	AudioType                       string                  `json:"audio_type"`
	BrandedContentTagInfo           BrandedContentTagInfo   `json:"branded_content_tag_info"`
	ContentAppreciationInfo         ContentAppreciationInfo `json:"content_appreciation_info"`
	ContextualHighlightInfo         interface{}             `json:"contextual_highlight_info"`
	CutoutStickerInfo               []interface{}           `json:"cutout_sticker_info"`
	DisableUseInClipsClientCache    bool                    `json:"disable_use_in_clips_client_cache"`
	ExternalMediaInfo               interface{}             `json:"external_media_info"`
	IsFanClubPromoVideo             bool                    `json:"is_fan_club_promo_video"`
	IsSharedToFb                    bool                    `json:"is_shared_to_fb"`
	MashupInfo                      MashupInfo              `json:"mashup_info"`
	MerchandisingPillInfo           interface{}             `json:"merchandising_pill_info"`
	MusicCanonicalID                string                  `json:"music_canonical_id"`
	MusicInfo                       interface{}             `json:"music_info"`
	OriginalSoundInfo               OriginalSoundInfo       `json:"original_sound_info"`
	ConsumptionInfo                 ConsumptionInfo         `json:"consumption_info"`
	AllowCreatorToRename            bool                    `json:"allow_creator_to_rename"`
	CanRemixBeSharedToFb            bool                    `json:"can_remix_be_shared_to_fb"`
	CanRemixBeSharedToFbExpansion   bool                    `json:"can_remix_be_shared_to_fb_expansion"`
	FormattedClipsMediaCount        interface{}             `json:"formatted_clips_media_count"`
	AudioParts                      []interface{}           `json:"audio_parts"`
	AudioPartsByFilter              []interface{}           `json:"audio_parts_by_filter"`
	IsExplicit                      bool                    `json:"is_explicit"`
	OriginalAudioSubtype            string                  `json:"original_audio_subtype"`
	IsAudioAutomaticallyAttributed  bool                    `json:"is_audio_automatically_attributed"`
	IsReuseDisabled                 bool                    `json:"is_reuse_disabled"`
	IsXpostFromFb                   bool                    `json:"is_xpost_from_fb"`
	XpostFbCreatorInfo              interface{}             `json:"xpost_fb_creator_info"`
	IsOriginalAudioDownloadEligible bool                    `json:"is_original_audio_download_eligible"`
	TrendRank                       interface{}             `json:"trend_rank"`
	AudioFilterInfos                []interface{}           `json:"audio_filter_infos"`
	OaOwnerIsMusicArtist            bool                    `json:"oa_owner_is_music_artist"`
	IsEligibleForAudioEffects       bool                    `json:"is_eligible_for_audio_effects"`
}

type AchievementsInfo struct {
	ShowAchievements      bool        `json:"show_achievements"`
	NumEarnedAchievements interface{} `json:"num_earned_achievements"`
}

type AdditionalAudioInfo struct {
	AdditionalAudioUsername string                 `json:"additional_audio_username"`
	AudioReattributionInfo  AudioReattributionInfo `json:"audio_reattribution_info"`
}

type AudioReattributionInfo struct {
	ShouldAllowRestore bool `json:"should_allow_restore"`
}

type AudioRankingInfo struct {
	BestAudioClusterID string `json:"best_audio_cluster_id"`
}

type BrandedContentTagInfo struct {
	CanAddTag bool `json:"can_add_tag"`
}

type ContentAppreciationInfo struct {
	Enabled             bool        `json:"enabled"`
	EntryPointContainer interface{} `json:"entry_point_container"`
}

type MashupInfo struct {
	MashupsAllowed                      bool        `json:"mashups_allowed"`
	CanToggleMashupsAllowed             bool        `json:"can_toggle_mashups_allowed"`
	HasBeenMashedUp                     bool        `json:"has_been_mashed_up"`
	IsLightWeightCheck                  bool        `json:"is_light_weight_check"`
	FormattedMashupsCount               interface{} `json:"formatted_mashups_count"`
	OriginalMedia                       interface{} `json:"original_media"`
	PrivacyFilteredMashupsMediaCount    interface{} `json:"privacy_filtered_mashups_media_count"`
	NonPrivacyFilteredMashupsMediaCount interface{} `json:"non_privacy_filtered_mashups_media_count"`
	MashupType                          interface{} `json:"mashup_type"`
	IsCreatorRequestingMashup           bool        `json:"is_creator_requesting_mashup"`
	HasNonmimicableAdditionalAudio      bool        `json:"has_nonmimicable_additional_audio"`
	IsPivotPageAvailable                bool        `json:"is_pivot_page_available"`
}

type OriginalSoundInfo struct {
	AudioAssetID                    int64           `json:"audio_asset_id"`
	MusicCanonicalID                interface{}     `json:"music_canonical_id"`
	ProgressiveDownloadURL          string          `json:"progressive_download_url"`
	DashManifest                    string          `json:"dash_manifest"`
	IgArtist                        User            `json:"ig_artist"`
	ShouldMuteAudio                 bool            `json:"should_mute_audio"`
	OriginalMediaID                 int64           `json:"original_media_id"`
	HideRemixing                    bool            `json:"hide_remixing"`
	DurationInMS                    int64           `json:"duration_in_ms"`
	TimeCreated                     int64           `json:"time_created"`
	OriginalAudioTitle              string          `json:"original_audio_title"`
	ConsumptionInfo                 ConsumptionInfo `json:"consumption_info"`
	AllowCreatorToRename            bool            `json:"allow_creator_to_rename"`
	CanRemixBeSharedToFb            bool            `json:"can_remix_be_shared_to_fb"`
	CanRemixBeSharedToFbExpansion   bool            `json:"can_remix_be_shared_to_fb_expansion"`
	FormattedClipsMediaCount        interface{}     `json:"formatted_clips_media_count"`
	AudioParts                      []interface{}   `json:"audio_parts"`
	AudioPartsByFilter              []interface{}   `json:"audio_parts_by_filter"`
	IsExplicit                      bool            `json:"is_explicit"`
	OriginalAudioSubtype            string          `json:"original_audio_subtype"`
	IsAudioAutomaticallyAttributed  bool            `json:"is_audio_automatically_attributed"`
	IsReuseDisabled                 bool            `json:"is_reuse_disabled"`
	IsXpostFromFb                   bool            `json:"is_xpost_from_fb"`
	XpostFbCreatorInfo              interface{}     `json:"xpost_fb_creator_info"`
	IsOriginalAudioDownloadEligible bool            `json:"is_original_audio_download_eligible"`
	TrendRank                       interface{}     `json:"trend_rank"`
	AudioFilterInfos                []interface{}   `json:"audio_filter_infos"`
	OaOwnerIsMusicArtist            bool            `json:"oa_owner_is_music_artist"`
	IsEligibleForAudioEffects       bool            `json:"is_eligible_for_audio_effects"`
}

type ConsumptionInfo struct {
	IsBookmarked              bool        `json:"is_bookmarked"`
	ShouldMuteAudioReason     string      `json:"should_mute_audio_reason"`
	IsTrendingInClips         bool        `json:"is_trending_in_clips"`
	ShouldMuteAudioReasonType interface{} `json:"should_mute_audio_reason_type"`
	DisplayMediaID            interface{} `json:"display_media_id"`
}

type CommentInformTreatment struct {
	ShouldHaveInformTreatment bool        `json:"should_have_inform_treatment"`
	Text                      string      `json:"text"`
	URL                       interface{} `json:"url"`
	ActionType                interface{} `json:"action_type"`
}

func main() {
	baseURL := "https://instagram-api-20231.p.rapidapi.com/api/reels_by_keyword?query=201771004"
	reelsMaxId := ""
	client := &http.Client{}
	var reelsSlice []Reel

	for {
		url := baseURL
		if reelsMaxId != "" {
			url += "&reels_max_id=" + reelsMaxId
		}

		// Create a new request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		// Set headers
		req.Header.Set("x-rapidapi-host", "instagram-api-20231.p.rapidapi.com")
		req.Header.Set("x-rapidapi-key", "e87d28030dmshf7754819f694597p198c99jsn16b59c0ebda9")

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Parse the response body
		var jsonResp Response
		if err := json.Unmarshal(body, &jsonResp); err != nil {
			log.Fatalf("Error parsing response body: %v", err)
		}

		clipsInfo := jsonResp.Data.Modules
		for _, module := range clipsInfo {
			for _, data := range module.Clips {
				reelInfo := Reel{
					ItemCode:      data.Media.Code,
					ItemLink:      fmt.Sprintf("http://www.instagram.com/%s", data.Media.Code),
					Username:      data.Media.User.Username,
					AuthorLink:    fmt.Sprintf("http://www.instagram.com/%s", data.Media.User.Username),
					LikesCount:    data.Media.LikeCount,
					CommentsCount: data.Media.CommentCount,
					RepostCount:   data.Media.ReshareCount,
				}

				timestamp := data.Media.TakenAt
				t := time.Unix(timestamp, 0)
				reelInfo.Date = t

				reelsSlice = append(reelsSlice, reelInfo)
			}
		}

		if !jsonResp.Data.HasMore {
			break
		}

		// Update reelsMaxId for the next request
		reelsMaxId = jsonResp.Data.ReelsMaxId
	}

	for _, reel := range reelsSlice {
		fmt.Printf("Clip ID: %s\n", reel.ItemCode)
		fmt.Printf("Clip Link: %s\n", reel.ItemLink)
		fmt.Printf("Author username: %s\n", reel.Username)
		fmt.Printf("Author link: %s\n", reel.AuthorLink)
		fmt.Printf("Likes count: %d\n", reel.LikesCount)
		fmt.Printf("Comments count: %d\n", reel.CommentsCount)
		fmt.Printf("Repost count: %d\n", reel.RepostCount)

		timestring := reel.Date.Format("2006-01-02 15:04:05")
		fmt.Printf("Date: %s\n", timestring)
	}
}
