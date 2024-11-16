package models

import "time"

// GetUserProfileParameters contains the parameters needed for getting a users profile
type GetUserProfileParameters struct {
	// The target username
	Username string
}

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

// GetUserRecentAchievementsParameters contains the parameters needed for getting a users recent achievements
type GetUserRecentAchievementsParameters struct {
	// The target username
	Username string

	// [Optional] Minutes to look back (Default 60)
	LookbackMinutes *int
}

// GetUserRecentAchievements describes elements of a users recent achievements
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

// GetAchievementsEarnedBetweenParameters contains the parameters needed for getting a users achievements earned between two dates
type GetAchievementsEarnedBetweenParameters struct {
	// The target username
	Username string

	// Time range start
	From time.Time

	// Time range end
	To time.Time
}

// GetAchievementsEarnedBetween describes elements of an achievement earned between two dates
type GetAchievementsEarnedBetween struct {
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

// GetAchievementsEarnedOnDayParameters contains the parameters needed for getting a users achievements earned on a specific day
type GetAchievementsEarnedOnDayParameters struct {
	// The target username
	Username string

	// Date
	Date time.Time
}

// GetAchievementsEarnedOnDay describes elements of an achievement earned on a specific day
type GetAchievementsEarnedOnDay struct {
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

// GetAchievementsEarnedOnDayParameters contains the parameters needed for getting a users game and achievement progress
type GetGameInfoAndUserProgressParameters struct {
	// The target username
	Username string

	// The target game ID
	GameID int

	// [Optional] Include additional award metadata (Default false)
	IncludeAwardMetadata *bool
}

type GetGameInfoAndUserProgressAchievement struct {
	ID                 int       `json:"ID"`
	NumAwarded         int       `json:"NumAwarded"`
	NumAwardedHardcore int       `json:"NumAwardedHardcore"`
	Title              string    `json:"Title"`
	Description        string    `json:"Description"`
	Points             int       `json:"Points"`
	TrueRatio          int       `json:"TrueRatio"`
	Author             string    `json:"Author"`
	DateModified       DateTime  `json:"DateModified"`
	DateCreated        DateTime  `json:"DateCreated"`
	BadgeName          string    `json:"BadgeName"`
	DisplayOrder       int       `json:"DisplayOrder"`
	MemAddr            string    `json:"MemAddr"`
	Type               string    `json:"type"`
	DateEarnedHardcore *DateTime `json:"DateEarnedHardcore"`
	DateEarned         *DateTime `json:"DateEarned"`
}

type GetGameInfoAndUserProgress struct {
	ID                         int                                           `json:"ID"`
	Title                      string                                        `json:"Title"`
	ConsoleID                  int                                           `json:"ConsoleID"`
	ForumTopicID               *int                                          `json:"ForumTopicID"`
	Flags                      int                                           `json:"Flags"`
	ImageIcon                  string                                        `json:"ImageIcon"`
	ImageTitle                 string                                        `json:"ImageTitle"`
	ImageIngame                string                                        `json:"ImageIngame"`
	ImageBoxArt                string                                        `json:"ImageBoxArt"`
	Publisher                  string                                        `json:"Publisher"`
	Developer                  string                                        `json:"Developer"`
	Genre                      string                                        `json:"Genre"`
	Released                   *DateOnly                                     `json:"Released"`
	ReleasedAtGranularity      *string                                       `json:"ReleasedAtGranularity"`
	IsFinal                    int                                           `json:"IsFinal"`
	RichPresencePatch          string                                        `json:"RichPresencePatch"`
	GuideURL                   *string                                       `json:"GuideURL"`
	ConsoleName                string                                        `json:"ConsoleName"`
	ParentGameID               *int                                          `json:"ParentGameID"`
	NumDistinctPlayers         int                                           `json:"NumDistinctPlayers"`
	NumAchievements            int                                           `json:"NumAchievements"`
	Achievements               map[int]GetGameInfoAndUserProgressAchievement `json:"Achievements"`
	NumDistinctPlayersCasual   int                                           `json:"NumDistinctPlayersCasual"`
	NumDistinctPlayersHardcore int                                           `json:"NumDistinctPlayersHardcore"`
	PlayersTotal               int                                           `json:"players_total"`
	AchievementsPublished      int                                           `json:"achievements_published"`
	PointsTotal                int                                           `json:"points_total"`
	NumAwardedToUser           int                                           `json:"NumAwardedToUser"`
	NumAwardedToUserHardcore   int                                           `json:"NumAwardedToUserHardcore"`
	UserCompletion             string                                        `json:"UserCompletion"`
	UserCompletionHardcore     string                                        `json:"UserCompletionHardcore"`
	HighestAwardKind           *string                                       `json:"HighestAwardKind"`
	HighestAwardDate           *RFC3339NumColonTZ                            `json:"HighestAwardDate"`
}

// GetUserCompletionProgressParameters contains the parameters needed for getting a users completion progress
type GetUserCompletionProgressParameters struct {
	// The target username
	Username string
}

// GetUserCompletionProgress
type GetUserCompletionProgress struct {
	Count   int                  `json:"Count"`
	Total   int                  `json:"Total"`
	Results []CompletionProgress `json:"Results"`
}

type CompletionProgress struct {
	GameID                int                `json:"GameID"`
	Title                 string             `json:"Title"`
	ImageIcon             string             `json:"ImageIcon"`
	ConsoleID             int                `json:"ConsoleID"`
	ConsoleName           string             `json:"ConsoleName"`
	MaxPossible           int                `json:"MaxPossible"`
	NumAwarded            int                `json:"NumAwarded"`
	NumAwardedHardcore    int                `json:"NumAwardedHardcore"`
	MostRecentAwardedDate RFC3339NumColonTZ  `json:"MostRecentAwardedDate"`
	HighestAwardKind      *string            `json:"HighestAwardKind"`
	HighestAwardDate      *RFC3339NumColonTZ `json:"HighestAwardDate"`
}

// GetUserAwardsParameters contains the parameters needed for getting a users awards
type GetUserAwardsParameters struct {
	// The target username
	Username string
}

type GetUserAwards struct {
	TotalAwardsCount          int     `json:"TotalAwardsCount"`
	HiddenAwardsCount         int     `json:"HiddenAwardsCount"`
	MasteryAwardsCount        int     `json:"MasteryAwardsCount"`
	CompletionAwardsCount     int     `json:"CompletionAwardsCount"`
	BeatenHardcoreAwardsCount int     `json:"BeatenHardcoreAwardsCount"`
	BeatenSoftcoreAwardsCount int     `json:"BeatenSoftcoreAwardsCount"`
	EventAwardsCount          int     `json:"EventAwardsCount"`
	SiteAwardsCount           int     `json:"SiteAwardsCount"`
	VisibleUserAwards         []Award `json:"VisibleUserAwards"`
}

type Award struct {
	AwardedAt      RFC3339NumColonTZ `json:"AwardedAt"`
	AwardType      string            `json:"AwardType"`
	AwardData      int               `json:"AwardData"`
	AwardDataExtra int               `json:"AwardDataExtra"`
	DisplayOrder   int               `json:"DisplayOrder"`
	Title          *string           `json:"Title"`
	ConsoleID      *int              `json:"ConsoleID"`
	ConsoleName    *string           `json:"ConsoleName"`
	Flags          *int              `json:"Flags"`
	ImageIcon      *string           `json:"ImageIcon"`
}

// GetUserClaimsParameters contains the parameters needed for getting a users claims
type GetUserClaimsParameters struct {
	// The target username
	Username string
}

type GetUserClaims struct {
	ID          int      `json:"ID"`
	User        string   `json:"User"`
	GameID      int      `json:"GameID"`
	GameTitle   string   `json:"GameTitle"`
	GameIcon    string   `json:"GameIcon"`
	ConsoleID   int      `json:"ConsoleID"`
	ConsoleName string   `json:"ConsoleName"`
	ClaimType   int      `json:"ClaimType"`
	SetType     int      `json:"SetType"`
	Status      int      `json:"Status"`
	Extension   int      `json:"Extension"`
	Special     int      `json:"Special"`
	Created     DateTime `json:"Created"`
	DoneTime    DateTime `json:"DoneTime"`
	Updated     DateTime `json:"Updated"`
	UserIsJrDev int      `json:"UserIsJrDev"`
	MinutesLeft int      `json:"MinutesLeft"`
}

// GetUserGameRankAndScoreParameters contains the parameters needed for getting a users rank and score
type GetUserGameRankAndScoreParameters struct {
	// The target username
	Username string

	// The target game ID
	GameID int
}

type GetUserGameRankAndScore struct {
	User       string    `json:"User"`
	UserRank   int       `json:"UserRank"`
	TotalScore int       `json:"TotalScore"`
	LastAward  *DateTime `json:"LastAward"`
}

// GetUserPointsParameters contains the parameters needed for getting a users points
type GetUserPointsParameters struct {
	// The target username
	Username string
}

type GetUserPoints struct {
	Points         int `json:"Points"`
	SoftcorePoints int `json:"SoftcorePoints"`
}
