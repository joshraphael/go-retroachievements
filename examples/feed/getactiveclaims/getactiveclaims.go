// Package getactiveclaims provides an example for getting information about all active set claims (max: 1000).
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getactiveclaims.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetActiveClaims(models.GetActiveClaimsParameters{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
