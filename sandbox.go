package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements/client"
)

/*
Test script for various endpoints. Add RA_API_KEY to your env and use `go run sandbox.go`
*/
func main() {
	host := "https://retroachievements.org"
	secret := os.Getenv("RA_API_KEY")

	raClient := client.New(host, secret)

	resp, err := raClient.GetUserProfile("HippopotamusRex")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", resp)

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
