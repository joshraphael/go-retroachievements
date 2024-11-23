// Package getcomments provides an example for getting comments of a specified kind: game, achievement, or user.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getcomments.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetComments(models.GetCommentsParameters{
		Type: models.GetCommentsUser{
			Username: "123",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
