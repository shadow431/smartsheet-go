package main

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"gitlab.allcomputergeek.net/libs/smartsheet-go"
)

func someFunction() {
	log.Println("Loading .env file")
	err := godotenv.Load("../.env")
	// handle error if needed
	if err != nil {
		log.Println("Error loading .env file")
	}

}

func TestNewClient(t *testing.T) {
	someFunction()
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Ensure the client is not nil
	if client == nil {
		t.Errorf("Client is nil")
	}

	// Ensure the AccessToken is set
	if client.AccessToken == "" {
		t.Errorf("AccessToken is not set")
	}

	// Ensure the BaseURL is set
	if client.BaseURL == "" {
		t.Errorf("BaseURL is not set")
	}
}

func TestGetSheet(t *testing.T) {
	someFunction()
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Get a sheet by ID
	sheet, err := client.GetSheet(1234567890)
	if err != nil {
		t.Errorf("Error getting sheet: %v", err)
	}

	log.Printf("Sheet ID: %d\n", sheet.ID)
}
