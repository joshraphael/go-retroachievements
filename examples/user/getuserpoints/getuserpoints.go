// Package getuserpoints get a user's total hardcore and softcore points.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserpoints.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserPoints("joshraphael")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
