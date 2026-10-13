package igdbModel

type Game struct {
	Game               int    `json:"game"`
	UID                string `json:"uid"`
	Year               int    `json:"year,omitempty"`
	ID                 int    `json:"id"`
	AgeRatings         []int  `json:"age_ratings,omitempty"`
	AlternativeNames   []int  `json:"alternative_names,omitempty"`
	Category           int    `json:"category"`
	Cover              int    `json:"cover,omitempty"`
	CreatedAt          int    `json:"created_at"`
	ExternalGames      []int  `json:"external_games,omitempty"`
	FirstReleaseDate   int    `json:"first_release_date,omitempty"`
	GameModes          []int  `json:"game_modes,omitempty"`
	Genres             []int  `json:"genres,omitempty"`
	InvolvedCompanies  []int  `json:"involved_companies,omitempty"`
	Keywords           []int  `json:"keywords,omitempty"`
	Name               string `json:"name"`
	Platforms          []int  `json:"platforms,omitempty"`
	PlayerPerspectives []int  `json:"player_perspectives,omitempty"`
	ReleaseDates       []int  `json:"release_dates,omitempty"`
	Screenshots        []int  `json:"screenshots,omitempty"`
	SimilarGames       []int  `json:"similar_games,omitempty"`
	Slug               string `json:"slug"`
	Storyline          string `json:"storyline,omitempty"`
	Summary            string `json:"summary,omitempty"`
	Tags               []int  `json:"tags,omitempty"`
	Themes             []int  `json:"themes,omitempty"`
	UpdatedAt          int    `json:"updated_at"`
	URL                string `json:"url"`
	Videos             []int  `json:"videos,omitempty"`
	Websites           []int  `json:"websites"`
	Checksum           string `json:"checksum"`
	GameLocalizations  []int  `json:"game_localizations,omitempty"`
	Collection         int    `json:"collection,omitempty"`
	MultiplayerModes   []int  `json:"multiplayer_modes,omitempty"`
	Status             int    `json:"status,omitempty"`
	LanguageSupports   []int  `json:"language_supports,omitempty"`
	Collections        []int  `json:"collections,omitempty"`
	Hypes              int    `json:"hypes,omitempty"`
	Artworks           []int  `json:"artworks,omitempty"`
	VersionParent      int    `json:"version_parent,omitempty"`
	VersionTitle       string `json:"version_title,omitempty"`
}
