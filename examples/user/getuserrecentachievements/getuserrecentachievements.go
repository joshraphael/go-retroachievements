// Package getuserrecentachievements provides an example for a users achievements in the last X minutes
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getuserrecentachievements.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserRecentAchievements("jamiras", 1000)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
