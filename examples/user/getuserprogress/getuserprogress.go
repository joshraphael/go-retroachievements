// Package getuserprogress provides an example for getting a users game progress
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserprogress.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserProgress(models.GetUserProgressParameters{
		Username: "jamiras",
		GameIDs:  []int{1, 16247},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
