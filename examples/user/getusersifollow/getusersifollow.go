// Package getusersifollow provides an example for the caller's "Following" users list.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusersifollow.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	count := 1
	offset := 1
	resp, err := client.GetUsersIFollow(models.GetUsersIFollowParameters{
		Count:  &count,
		Offset: &offset,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
