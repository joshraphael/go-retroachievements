package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements/pkg/retroachievements"
)

/*
Test script for various endpoints. Add RA_API_KEY to your env and use `go run sandbox.go`
*/
func main() {
	host := "https://retroachievements.org"
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.New(host, secret)

	resp, err := client.User.GetUserRecentAchievements("ChronoGear", 60)
	if err != nil {
		panic(err)
	}

	for i := range resp {
		fmt.Printf("%+v\n", resp[i])
	}
}
