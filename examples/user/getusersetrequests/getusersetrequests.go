// Package getusersetrequests provides an example for a user's list of set requests
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusersetrequests.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	all := true
	resp, err := client.GetUserSetRequests(models.GetUserSetRequestsParameters{
		Username: "jamiras",
		All:      &all,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
