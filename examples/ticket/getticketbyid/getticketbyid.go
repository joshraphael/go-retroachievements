// Package getticketbyid provides an example for getting ticket metadata information about a single achievement ticket, targeted by its ticket ID.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getticketbyid.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetTicketByID(models.GetTicketByIDParameters{
		TicketID: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
