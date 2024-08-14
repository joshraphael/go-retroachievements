// Package getuserrecentachievements provides an example for a users achievements in the last X minutes
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements/client"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getuserrecentachievements.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	raClient := client.DefaultClient(secret)

	resp, err := raClient.GetUserRecentAchievements("jamiras", 1000)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
