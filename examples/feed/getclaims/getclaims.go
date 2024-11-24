// Package getclaims provides an example for getting information about all achievement set development claims of a specified kind: completed, dropped, or expired (max: 1000).
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getclaims.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetClaims(models.GetClaimsParameters{
		Kind: &models.GetClaimsParametersKindDropped{},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
