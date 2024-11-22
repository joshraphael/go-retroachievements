// Package getconsoleids provides an example for getting the complete list of all system ID and name pairs on the site.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getconsoleids.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetConsoleIDs(models.GetConsoleIDsParameters{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
