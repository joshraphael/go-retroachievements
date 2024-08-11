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
	host := "https://retroachievements.org/API"
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.New(host, secret)

	profile, err := client.User.GetUserProfile("MaxMilyin")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", profile)
}
