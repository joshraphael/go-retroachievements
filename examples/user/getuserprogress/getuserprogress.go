// Package getuserprogress provides an example for getting a users game progress
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserprogress.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserProgress("jamiras", []int{1, 16247})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
