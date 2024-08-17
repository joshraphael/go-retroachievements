// Package getachievementsearnedbetween provides an example for a users achievements between two timestamps
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getachievementsearnedbetween.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	if err != nil {
		panic(err)
	}
	later := now.Add(10 * time.Minute)

	resp, err := client.GetAchievementsEarnedBetween("jamiras", now, later)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
