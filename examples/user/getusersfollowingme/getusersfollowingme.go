// Package getusersfollowingme provides an example for the caller's "Followers" users list.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusersfollowingme.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	count := 1
	offset := 1
	resp, err := client.GetUsersFollowingMe(models.GetUsersFollowingMeParameters{
		Count:  &count,
		Offset: &offset,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
