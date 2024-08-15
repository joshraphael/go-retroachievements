// Package getachievementsearnedonday provides an example for a users achievements on a specific day
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getachievementsearnedonday.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	if err != nil {
		panic(err)
	}

	resp, err := client.GetAchievementsEarnedOnDay("jamiras", now)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
