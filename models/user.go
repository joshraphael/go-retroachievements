package models

// GetUserProfile describes elements of a users profile
type GetUserProfile struct {
	// Username of the profile
	User string `json:"User"`

	// URL resource to the image of the profile avatar
	UserPic string `json:"UserPic"`

	// Date the user joined
	MemberSince DateTime `json:"MemberSince"`

	// Last game played rich presence text
	RichPresenceMsg string `json:"RichPresenceMsg"`

	// ID of last game played
	LastGameID int `json:"LastGameID"`

	// Number of achievements unlocked by players
	ContribCount int `json:"ContribCount"`

	// Number of points awarded to players
	ContribYield int `json:"ContribYield"`

	// Number of points awarded on this profile
	TotalPoints int `json:"TotalPoints"`

	// Number of softcore points awarded on this profile
	TotalSoftcorePoints int `json:"TotalSoftcorePoints"`

	// Number of RetroPoints awarded on this profile
	TotalTruePoints int `json:"TotalTruePoints"`

	// Site permissions (0 = Normal, 1 = Jr. Dev, 2 = Developer, 3 = Moderator, 4 = Admin)
	Permissions int `json:"Permissions"`

	//  "1" if the user is untracked, otherwise "0"
	Untracked int `json:"Untracked"`

	// ID of the profile
	ID int `json:"ID"`

	// Allows other users to comment on their profile wall
	UserWallActive bool `json:"UserWallActive"`

	// Custom status message displayed on profile
	Motto string `json:"Motto"`
}

type GetUserRecentAchievements struct {
	// Title of the achievement
	Title string `json:"Title"`

	// Description of the achievement
	Description string `json:"Description"`

	// Points awarded
	Points int `json:"Points"`

	// Ratio of points the achievemnet is worth
	TrueRatio int `json:"TrueRatio"`

	// Username of the author of the achievement
	Author string `json:"Author"`

	// The date the achievement was unlocked
	Date DateTime `json:"Date"`

	// Mode the achievement was unlocked in: 1 if in hardcore mode, 0 if not
	HardcoreMode int `json:"HardcoreMode"`

	// The ID of the achievement
	AchievementID int `json:"AchievementID"`

	// Name of the padge resource
	BadgeName string `json:"BadgeName"`

	// Type of achievement (standard, missable, progression, win_condition)
	Type string `json:"Type"`

	// Title of the game
	GameTitle string `json:"GameTitle"`

	// URL resource of the game icon
	GameIcon string `json:"GameIcon"`

	// ID of the game
	GameID int `json:"GameID"`

	// Common name of the console
	ConsoleName string `json:"ConsoleName"`

	// URL resource to the image used for the achievment badge
	BadgeURL string `json:"BadgeURL"`

	// URL resource to the game page
	GameURL string `json:"GameURL"`
}
