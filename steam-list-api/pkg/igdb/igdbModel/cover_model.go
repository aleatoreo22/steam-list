package igdbModel

type Cover struct {
	ID           int    `json:"id"`
	Game         int    `json:"game"`
	Height       int    `json:"height"`
	ImageID      string `json:"image_id"`
	URL          string `json:"url"`
	Width        int    `json:"width"`
	Checksum     string `json:"checksum"`
	AlphaChannel bool   `json:"alpha_channel,omitempty"`
	Animated     bool   `json:"animated,omitempty"`
}
