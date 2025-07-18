# <p align="center">![ra_gopher](assets/ra_gopher_small.png "Retro Achievements Gopher")<br>go-retroachievements</p>

<p align="center">
    <i>A Golang library that lets you get achievement, user, game data and more from RetroAchievements.</i>
</p>

<p align="center">
    <a href="https://api-docs.retroachievements.org/getting-started.html"><strong>Documentation: Get Started</strong></a>
</p>

<br>
<hr />

[![GitHub License](https://img.shields.io/github/license/joshraphael/go-retroachievements)](https://github.com/joshraphael/go-retroachievements/blob/main/LICENSE)
[![godoc](https://pkg.go.dev/badge/github.com/joshraphael/go-retroachievements.svg)](https://pkg.go.dev/github.com/joshraphael/go-retroachievements)
[![pipeline](https://github.com/joshraphael/go-retroachievements/actions/workflows/go.yaml/badge.svg)](https://github.com/joshraphael/go-retroachievements/actions/workflows/go.yaml)
![coverage](https://raw.githubusercontent.com/joshraphael/go-retroachievements/badges/.badges/main/coverage.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshraphael/go-retroachievements)](https://goreportcard.com/report/github.com/joshraphael/go-retroachievements)
[![GitHub Tag](https://img.shields.io/github/v/tag/joshraphael/go-retroachievements)](https://github.com/joshraphael/go-retroachievements/tags)
[![GitHub repo size](https://img.shields.io/github/repo-size/joshraphael/go-retroachievements)](https://github.com/joshraphael/go-retroachievements/archive/main.zip)

## Installation
Use go get to install the latest version of the library.
```sh
go get github.com/joshraphael/go-retroachievements@latest
```
Then include it in your application:
```go
import "github.com/joshraphael/go-retroachievements"
import "github.com/joshraphael/go-retroachievements/models"
```

## Usage

Construct a new Retro Achievement client using your personal web API key

```go
client := retroachievements.NewClient("<your web API key>")
```

you can now use the client to call any of the available endpoints, for example:

```go
profile, err := client.GetUserProfile(models.GetUserProfileParameters{
    Username: "jamiras",
})
```

Check out the [examples](examples/) directory for how to call each endpoint, as well as our [GoDocs](https://pkg.go.dev/github.com/joshraphael/go-retroachievements)

## API
For convenience, the API docs and examples can be found in the tables below

<h3>User</h3>

|Function|Description|Links|
|-|-|-|
|`GetUserProfile()`|Get a user's basic profile information.|[docs](https://api-docs.retroachievements.org/v1/get-user-profile.html) \| [example](examples/user/getuserprofile/getuserprofile.go)|
|`GetUserRecentAchievements()`|Get a list of achievements recently earned by the user.|[docs](https://api-docs.retroachievements.org/v1/get-user-recent-achievements.html) \| [example](examples/user/getuserrecentachievements/getuserrecentachievements.go)|
|`GetAchievementsEarnedBetween()`|Get a list of achievements earned by a user between two dates.|[docs](https://api-docs.retroachievements.org/v1/get-achievements-earned-between.html) \| [example](examples/user/getachievementsearnedbetween/getachievementsearnedbetween.go)|
|`GetAchievementsEarnedOnDay()`|Get a list of achievements earned by a user on a given date.|[docs](https://api-docs.retroachievements.org/v1/get-achievements-earned-on-day.html) \| [example](examples/user/getachievementsearnedonday/getachievementsearnedonday.go)|
|`GetGameInfoAndUserProgress()`|Get metadata about a game as well as a user's progress on that game.|[docs](https://api-docs.retroachievements.org/v1/get-game-info-and-user-progress.html) \| [example](examples/user/getgameinfoanduserprogress/getgameinfoanduserprogress.go)|
|`GetUserCompletionProgress()`|Get metadata about all the user's played games and any awards associated with them.|[docs](https://api-docs.retroachievements.org/v1/get-user-completion-progress.html) \| [example](examples/user/getusercompletionprogress/getusercompletionprogress.go)|
|`GetUserAwards()`|Get a list of a user's site awards/badges.|[docs](https://api-docs.retroachievements.org/v1/get-user-awards.html) \| [example](examples/user/getuserawards/getuserawards.go)|
|`GetUserClaims()`|Get a list of set development claims made over the lifetime of a user.|[docs](https://api-docs.retroachievements.org/v1/get-user-claims.html) \| [example](examples/user/getuserclaims/getuserclaims.go)|
|`GetUserGameRankAndScore()`|Get metadata about how a user has performed on a given game.|[docs](https://api-docs.retroachievements.org/v1/get-user-game-rank-and-score.html) \| [example](examples/user/getusergamerankandscore/getusergamerankandscore.go)|
|`GetUserPoints()`|Get a user's total hardcore and softcore points.|[docs](https://api-docs.retroachievements.org/v1/get-user-points.html) \| [example](examples/user/getuserpoints/getuserpoints.go)|
|`GetUserProgress()`|Get a user's progress on a list of specified games.|[docs](https://api-docs.retroachievements.org/v1/get-user-progress.html) \| [example](examples/user/getuserprogress/getuserprogress.go)|
|`GetUserRecentlyPlayedGames()`|Get a list of games a user has recently played.|[docs](https://api-docs.retroachievements.org/v1/get-user-recently-played-games.html) \| [example](examples/user/getuserrecentlyplayedgames/getuserrecentlyplayedgames.go)|
|`GetUserSummary()`|Get a user's profile metadata.|[docs](https://api-docs.retroachievements.org/v1/get-user-summary.html) \| [example](examples/user/getusersummary/getusersummary.go)|
|`GetUserCompletedGames()`|[Deprecated] Get hardcore and softcore completion metadata about games a user has played.|[docs](https://api-docs.retroachievements.org/v1/get-user-completed-games.html) \| [example](examples/user/getusercompletedgames/getusercompletedgames.go)|
|`GetUserWantToPlayList()`|Get a user's "Want to Play Games" list.|[docs](https://api-docs.retroachievements.org/v1/get-user-want-to-play-list.html) \| [example](examples/user/getuserwanttoplaylist/getuserwanttoplaylist.go)|
|`GetUsersIFollow()`|Get the caller's "Following" users list.|[docs](https://api-docs.retroachievements.org/v1/get-users-i-follow.html) \| [example](examples/user/getusersifollow/getusersifollow.go)|
|`GetUsersFollowingMe()`|Get the caller's "Followers" users list.|[docs](https://api-docs.retroachievements.org/v1/get-users-following-me.html) \| [example](examples/user/getusersfollowingme/getusersfollowingme.go)|
|`GetUserSetRequests()`|Get a user's list of set requests.|[docs](https://api-docs.retroachievements.org/v1/get-user-set-requests.html) \| [example](examples/user/getusersetrequests/getusersetrequests.go)|

<h3>Game</h3>

|Function|Description|Links|
|-|-|-|
|`GetGame()`|Get basic metadata about a game.|[docs](https://api-docs.retroachievements.org/v1/get-game.html) \| [example](examples/game/getgame/getgame.go)|
|`GetGameExtended()`|Get extended metadata about a game.|[docs](https://api-docs.retroachievements.org/v1/get-game-extended.html) \| [example](examples/game/getgameextended/getgameextended.go)|
|`GetGameHashes()`|Get the hashes linked to a game.|[docs](https://api-docs.retroachievements.org/v1/get-game-hashes.html) \| [example](examples/game/getgamehashes/getgamehashes.go)|
|`GetAchievementCount()`|Get the list of achievement IDs for a game.|[docs](https://api-docs.retroachievements.org/v1/get-achievement-count.html) \| [example](examples/game/getachievementcount/getachievementcount.go)|
|`GetAchievementDistribution()`|Gets how many players have unlocked how many achievements for a game.|[docs](https://api-docs.retroachievements.org/v1/get-achievement-distribution.html) \| [example](examples/game/getachievementdistribution/getachievementdistribution.go)|
|`GetGameRankAndScore()`|Gets metadata about either the latest masters for a game, or the highest points earners for a game.|[docs](https://api-docs.retroachievements.org/v1/get-game-rank-and-score.html) \| [example](examples/game/getgamerankandscore/getgamerankandscore.go)|

<h3>Leaderboards</h3>

|Function|Description|Links|
|-|-|-|
|`GetGameLeaderboards()`|Gets a given games's list of leaderboards.|[docs](https://api-docs.retroachievements.org/v1/get-game-leaderboards.html) \| [example](examples/leaderboards/getgameleaderboards/getgameleaderboards.go)|
|`GetLeaderboardEntries()`|Gets a given leadboard's entries.|[docs](https://api-docs.retroachievements.org/v1/get-leaderboard-entries.html) \| [example](examples/leaderboards/getleaderboardentries/getleaderboardentries.go)|
|`GetUserGameLeaderboards()`|Gets a user's list of leaderboards for a given game.|[docs](https://api-docs.retroachievements.org/v1/get-user-game-leaderboards.html) \| [example](examples/leaderboards/getusergameleaderboards/getusergameleaderboards.go)|

<h3>System</h3>

|Function|Description|Links|
|-|-|-|
|`GetConsoleIDs()`|Gets the complete list of all system ID and name pairs on the site.|[docs](https://api-docs.retroachievements.org/v1/get-console-ids.html) \| [example](examples/system/getconsoleids/getconsoleids.go)|
|`GetGameList()`|Gets the complete list of games for a specified console on the site.|[docs](https://api-docs.retroachievements.org/v1/get-game-list.html) \| [example](examples/system/getgamelist/getgamelist.go)|

<h3>Achievement</h3>

|Function|Description|Links|
|-|-|-|
|`GetAchievementUnlocks()`|Gets a list of users who have earned an achievement.|[docs](https://api-docs.retroachievements.org/v1/get-achievement-unlocks.html) \| [example](examples/achievement/getachievementunlocks/getachievementunlocks.go)|

<h3>Comment</h3>

|Function|Description|Links|
|-|-|-|
|`GetComments()`|Gets comments of a specified kind: game, achievement, or user.|[docs](https://api-docs.retroachievements.org/v1/get-comments.html) \| [example](examples/comment/getcomments/getcomments.go)|

<h3>Feed</h3>

|Function|Description|Links|
|-|-|-|
|`GetRecentGameAwards()`|Gets all recently granted game awards across the site's userbase.|[docs](https://api-docs.retroachievements.org/v1/get-recent-game-awards.html) \| [example](examples/feed/getrecentgameawards/getrecentgameawards.go)|
|`GetActiveClaims()`|Gets information about all active set claims (max: 1000).|[docs](https://api-docs.retroachievements.org/v1/get-active-claims.html) \| [example](examples/feed/getactiveclaims/getactiveclaims.go)|
|`GetClaims()`|Gets information about all achievement set development claims of a specified kind: completed, dropped, or expired (max: 1000).|[docs](https://api-docs.retroachievements.org/v1/get-claims.html) \| [example](examples/feed/getclaims/getclaims.go)|
|`GetTopTenUsers()`|Gets the current top ten users, ranked by hardcore points, on the site.|[docs](https://api-docs.retroachievements.org/v1/get-top-ten-users.html) \| [example](examples/feed/gettoptenusers/gettoptenusers.go)|

<h3>Event</h3>

|Function|Description|Links|
|-|-|-|
|`GetAchievementOfTheWeek()`|Gets comprehensive metadata about the current Achievement of the Week.|[docs](https://api-docs.retroachievements.org/v1/get-achievement-of-the-week.html) \| [example](examples/event/getachievementoftheweek/getachievementoftheweek.go)|

<h3>Ticket</h3>

|Function|Description|Links|
|-|-|-|
|`GetTicketByID()`|Gets ticket metadata information about a single achievement ticket, targeted by its ticket ID.|[docs](https://api-docs.retroachievements.org/v1/get-ticket-data/get-ticket-by-id.html) \| [example](examples/ticket/getticketbyid/getticketbyid.go)|
|`GetMostTicketedGames()`|Gets the games on the site with the highest count of opened achievement tickets.|[docs](https://api-docs.retroachievements.org/v1/get-ticket-data/get-most-ticketed-games.html) \| [example](examples/ticket/getmostticketedgames/getmostticketedgames.go)|
|`GetMostRecentTickets()`|Gets ticket metadata information about the latest opened achievement tickets on RetroAchievements.|[docs](https://api-docs.retroachievements.org/v1/get-ticket-data/get-most-recent-tickets.html) \| [example](examples/ticket/getmostrecenttickets/getmostrecenttickets.go)|
|`GetGameTicketStats()`|Gets ticket stats for a game, targeted by that game's unique ID.|[docs](https://api-docs.retroachievements.org/v1/get-ticket-data/get-game-ticket-stats.html) \| [example](examples/ticket/getgameticketstats/getgameticketstats.go)|
|`GetDeveloperTicketStats()`|Gets ticket stats for a developer, targeted by that developer's site username.|[docs](https://api-docs.retroachievements.org/v1/get-ticket-data/get-developer-ticket-stats.html) \| [example](examples/ticket/getdeveloperticketstats/getdeveloperticketstats.go)|
|`GetAchievementTicketStats()`|Gets ticket stats for an achievement, targeted by that achievement's unique ID.|[docs](https://api-docs.retroachievements.org/v1/get-ticket-data/get-achievement-ticket-stats.html) \| [example](examples/ticket/getachievementticketstats/getachievementticketstats.go)|

## Documentation
This library uses [doc2go](https://abhinav.github.io/doc2go/) to generate local static docs for the package. you will first need to install the package and then run the make command to serve up the site

```sh
go install go.abhg.dev/doc2go@latest
make docs
```