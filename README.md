# <p align="center">![ra_gopher](assets/ra_gopher_small.png "Retro Achievements Gopher")<br>go-retroachievements</p>

<p align="center">
    <i>A Golang library that lets you get achievement, user, and game data from RetroAchievements.</i>
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