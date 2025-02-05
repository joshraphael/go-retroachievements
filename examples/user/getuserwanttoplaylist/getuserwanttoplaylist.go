// Package getuserwanttoplaylist provides an example for a users "Want to Play Games" list
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserwanttoplaylist.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserWantToPlayList(models.GetUserWantToPlayListParameters{
		Username: "spoony",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
