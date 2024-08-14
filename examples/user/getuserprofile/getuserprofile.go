// Package getuserprofile provides an example for getting a users profile
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements/client"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getuserprofile.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	raClient := client.DefaultClient(secret)

	resp, err := raClient.GetUserProfile("jamiras")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
