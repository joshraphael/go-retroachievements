// Package getachievementsearnedbetween provides an example for a users achievements between two timestamps
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joshraphael/go-retroachievements/client"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getachievementsearnedbetween.go`
*/
func main() {
	host := "https://retroachievements.org"
	secret := os.Getenv("RA_API_KEY")

	raClient := client.New(host, secret)

	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	if err != nil {
		panic(err)
	}
	later := now.Add(10 * time.Minute)

	resp, err := raClient.GetAchievementsEarnedBetween("jamiras", now, later)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
