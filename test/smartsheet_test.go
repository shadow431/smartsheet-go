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

	log.Info().Msgf("Columns: %v\n", sheet.Columns)
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

func TestListSheetDiscussions(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// List sheet discussions
	discussions, err := client.ListSheetDiscussions(sheet_id)
	if err != nil {
		t.Errorf("Error listing sheet discussions: %v", err)
	}

	log.Info().Msgf("Discussions: %v\n", discussions)
}

func TestListRowDiscussions(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	row_id, err := strconv.Atoi(os.Getenv("ROW_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// List row discussions
	discussions, err := client.ListRowDiscussions(sheet_id, row_id)
	if err != nil {
		t.Errorf("Error listing row discussions: %v", err)
	}

	log.Info().Msgf("Discussions: %v\n", discussions)
}

func TestListRowDiscussionsWithInclude(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	row_id, err := strconv.Atoi(os.Getenv("ROW_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// List row discussions
	discussions, err := client.ListRowDiscussions(sheet_id, row_id, "comments")
	if err != nil {
		t.Errorf("Error listing row discussions: %v", err)
	}

	log.Info().Msgf("Discussions: %v\n", discussions)
}

func TestGetComment(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	comment_id, err := strconv.Atoi(os.Getenv("COMMENT_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Get a comment by ID
	comment, err := client.GetComment(sheet_id, comment_id)
	if err != nil {
		t.Errorf("Error getting comment: %v", err)
	}

	log.Info().Msgf("Comment: %v\n", comment)
}

func TestCreateRowDiscussion(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	row_id, err := strconv.Atoi(os.Getenv("ROW_ID"))
	comment := "This is a test Discussion"
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Create a discussion
	discussions, err := client.CreatRowDiscussion(sheet_id, row_id, comment)
	if err != nil {
		t.Errorf("Error creating discussion: %v", err)
	}

	log.Info().Msgf("Discussions: %v\n", discussions)
}

func TestCreateComment(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	discussion_id, err := strconv.Atoi(os.Getenv("DISCUSSION_ID"))
	text := "This is a test Comment"
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Create a comment
	comment, err := client.CreateComment(sheet_id, discussion_id, text)
	if err != nil {
		t.Errorf("Error creating comment: %v", err)
	}

	log.Info().Msgf("Comment: %v\n", comment)
}

func TestGetColumns(t *testing.T) {
	someFunction()
	sheet_id, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	// Create a new SmartsheetClient
	client := smartsheet.NewClient()

	// Get columns
	columns, err := client.GetColumns(sheet_id)
	if err != nil {
		t.Errorf("Error getting columns: %v", err)
	}

	t.Logf("Columns: %v\n", columns)
}
