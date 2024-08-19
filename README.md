# <p align="center">![ra_gopher](assets/ra_gopher_small.png "Retro Achievements Gopher")<br>go-retroachievements</p>

<p align="center">
    <i>A Golang library that lets you get achievement, user, game data and more from RetroAchievements.</i>
</p>

<p align="center">
    <a href="https://api-docs.retroachievements.org/getting-started.html"><strong>Documentation: Get Started</strong></a>
</p>

<br>
<hr />

[![pipeline](https://github.com/joshraphael/go-retroachievements/actions/workflows/ci.yaml/badge.svg)](https://github.com/joshraphael/go-retroachievements/actions)

## Installation
Use go get to install the latest version of the library.
```sh
go get github.com/joshraphael/go-retroachievements@latest
```
Then include it in your application:
```go
import "github.com/joshraphael/go-retroachievements"
```

## Usage

Construct a new Retro Achievement client using your personal web API key

```go
client := retroachievements.NewClient("<your web API key>")
```

you can now use the client to call any of the available endpoints, for example:

```go
profile, err := client.GetUserProfile("jamiras")
```

Check out the [examples](examples/) directory for how to call each endpoint, as well as our GoDocs (TBD)

## API
For convenience, the API docs and examples can be found in the tables below

<h3>User</h3>

|Function|Description|Links|
|-|-|-|
|`GetUserProfile(string)`|Get a user's basic profile information.|[docs](https://api-docs.retroachievements.org/v1/get-user-profile.html) \| [example](examples/user/getuserprofile/getuserprofile.go)|
|`GetUserRecentAchievements(string,int)`|Get a list of achievements recently earned by the user.|[docs](https://api-docs.retroachievements.org/v1/get-user-recent-achievements.html) \| [example](examples/user/getuserrecentachievements/getuserrecentachievements.go)|
|`GetAchievementsEarnedBetween(string,Time,Time)`|Get a list of achievements earned by a user between two dates.|[docs](https://api-docs.retroachievements.org/v1/get-achievements-earned-between.html) \| [example](examples/user/getachievementsearnedbetween/getachievementsearnedbetween.go)|
|`GetAchievementsEarnedOnDay(string,Time)`|Get a list of achievements earned by a user on a given date.|[docs](https://api-docs.retroachievements.org/v1/get-achievements-earned-on-day.html) \| [example](examples/user/getachievementsearnedonday/getachievementsearnedonday.go)|
|`GetGameInfoAndUserProgress(string,int,bool)`|Get metadata about a game as well as a user's progress on that game.|[docs](https://api-docs.retroachievements.org/v1/get-game-info-and-user-progress.html) \| [example](examples/user/getgameinfoanduserprogress/getgameinfoanduserprogress.go)|
|`GetUserCompletionProgress(string)`|Get metadata about all the user's played games and any awards associated with them.|[docs](https://api-docs.retroachievements.org/v1/get-user-completion-progress.html) \| [example](examples/user/getusercompletionprogress/getusercompletionprogress.go)|
|`GetUserAwards(string)`|Get a list of a user's site awards/badges.|[docs](https://api-docs.retroachievements.org/v1/get-user-awards.html) \| [example](examples/user/getuserawards/getuserawards.go)|

<h3>Game</h3>

|Function|Description|Links|
|-|-|-|
|`GetGame(int)`|Get basic metadata about a game.|[docs](https://api-docs.retroachievements.org/v1/get-game.html) \| [example](examples/game/getgame/getgame.go)|
|`GetGameExtended(int)`|Get extended metadata about a game.|[docs](https://api-docs.retroachievements.org/v1/get-game-extended.html) \| [example](examples/game/getgameextended/getgameextended.go)|