// Package getuserawards provides an example for getting a users awards
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserawards.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserAwards("jamiras")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
