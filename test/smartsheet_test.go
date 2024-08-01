package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
	"gitlab.allcomputergeek.net/libs/smartsheet-go"
)

func someFunction() {
	log.Info().Msg("Loading .env file")
	err := godotenv.Load("../.env")
	// handle error if needed
	if err != nil {
		log.Info().Msg("Error loading .env file")
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
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Get a sheet by ID
	sheet, err := client.GetSheet(sheet_id)
	if err != nil {
		t.Errorf("Error getting sheet: %v", err)
	}

	log.Info().Msgf("Ownder-ID: %d\n", sheet.OWNDER_ID)
}

func TestGetReport(t *testing.T) {
	someFunction()
	report_id, err := strconv.Atoi(os.Getenv("REPORT_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Get a report by ID
	report, err := client.GetReport(report_id)
	if err != nil {
		t.Errorf("Error getting report: %v", err)
	}

	log.Info().Msgf("Ownder-ID: %d\n", report.OWNDER_ID)
}
