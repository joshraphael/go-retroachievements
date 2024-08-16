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
Click the function names to open their complete docs on the docs site.

<h3>User</h3>

* [`GetUserProfile()`](https://api-docs.retroachievements.org/v1/get-user-profile.html) - Get a user's basic profile information.
* [`GetUserRecentAchievements()`](https://api-docs.retroachievements.org/v1/get-user-recent-achievements.html) - Get a list of achievements recently earned by the user.
* [`GetAchievementsEarnedBetween()`](https://api-docs.retroachievements.org/v1/get-achievements-earned-between.html) - Get a list of achievements earned by a user between two dates.
* [`GetAchievementsEarnedOnDay()`](https://api-docs.retroachievements.org/v1/get-achievements-earned-on-day.html) - Get a list of achievements earned by a user on a given date.

<h3>Game</h3>

* [`GetGame()`](https://api-docs.retroachievements.org/v1/get-game.html) - Get basic metadata about a game.
* [`GetGameExtended()`](https://api-docs.retroachievements.org/v1/get-game-extended.html) - Get extended metadata about a game.