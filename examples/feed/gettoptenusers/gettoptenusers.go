// Package gettoptenusers provides an example for getting the current top ten users, ranked by hardcore points, on the site.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run gettoptenusers.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetTopTenUsers(models.GetTopTenUsersParameters{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
