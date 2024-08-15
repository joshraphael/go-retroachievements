// Package getuserprofile provides an example for getting a users profile
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getuserprofile.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserProfile("jamiras")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
