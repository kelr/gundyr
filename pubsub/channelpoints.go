package pubsub

import (
	"time"
)

const (
	channelPointTopic = "channel-points-channel-v1."
)

// ChannelPointsEvent contains the type and data payload for a channel points event
type ChannelPointsEvent struct {
	Type string            `json:"type"`
	Data ChannelPointsData `json:"data"`
}

// ChannelPointsData contains the time the reward was redeemed and the redemption data
type ChannelPointsData struct {
	TimeStamp  time.Time      `json:"timestamp"`
	Redemption RedemptionData `json:"redemption"`
}

// RedemptionData contains metadata about the redeemed reward
type RedemptionData struct {
	ID         string           `json:"id"`
	User       RedemptionUser   `json:"user"`
	ChannelID  string           `json:"channel_id"`
	RedeemedAt time.Time        `json:"redeemed_at"`
	Reward     RedemptionReward `json:"reward"`
	UserInput  string           `json:"user_input"`
	Status     string           `json:"status"`
}

// RedemptionUser represents the user who redeemed the reward
type RedemptionUser struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
}

// RedemptionReward represents information about the reward redeemed
type RedemptionReward struct {
	ID                    string                 `json:"id"`
	ChannelID             string                 `json:"channel_id"`
	Title                 string                 `json:"title"`
	Prompt                string                 `json:"prompt"`
	Cost                  int                    `json:"cost"`
	IsUserInputRequired   bool                   `json:"is_user_input_required"`
	IsSubOnly             bool                   `json:"is_sub_only"`
	Image                 RedemptionImage        `json:"image"`
	DefaultImage          RedemptionImage        `json:"default_image"`
	BackgroundColor       string                 `json:"background_color"`
	IsEnabled             bool                   `json:"is_enabled"`
	IsPaused              bool                   `json:"is_paused"`
	IsInStock             bool                   `json:"is_in_stock"`
	MaxPerStream          RedemptionMaxPerStream `json:"max_per_stream"`
	ShouldRedemptionsSkip bool                   `json:"should_redemptions_skip_request_queue"`
	TemplateID            string                 `json:"template_id"`
	UpdatedForIndicatorAt time.Time              `json:"updated_for_indicator_at"`
}

// RedemptionImage represents the cute image used on the redemption button
type RedemptionImage struct {
	URL1x string `json:"url_1x"`
	URL2x string `json:"url_2x"`
	URL4x string `json:"url_4x"`
}

// RedemptionMaxPerStream represents information about redemption limits per stream
type RedemptionMaxPerStream struct {
	IsEnabled    bool `json:"is_enabled"`
	MaxPerStream int  `json:"max_per_stream"`
}

// ListenChannelPoints subscribes a handler function to the Channel Points topic with the provided id.
// The handler will be called with a populated ChannelPointsData struct when the event is received.
func (c *Client) ListenChannelPoints(handler func(*ChannelPointsData)) {
	c.channelPointHandler = handler
	if c.IsConnected() {
		c.listen(&[]string{channelPointTopic + c.ID})
	}
}

// UnlistenChannelPoints removes the current handler function from the channel points event topic and
// unlistens from the topic.
func (c *Client) UnlistenChannelPoints() {
	c.channelPointHandler = nil
	if c.IsConnected() {
		c.unlisten(&[]string{channelPointTopic + c.ID})
	}
}
