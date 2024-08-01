//Client that will be used by other functions to interact with the Smartsheet API
package smartsheet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// SmartsheetClient struct
type SmartsheetClient struct {
	AccessToken string

	// BaseURL is the base URL for the Smartsheet API
	BaseURL string
}

// Sheet struct
type Sheet struct {
	ID int `json:"id"`
}

// NewClient creates a new SmartsheetClient
func NewClient() *SmartsheetClient {
	return &SmartsheetClient{
		AccessToken: os.Getenv("SMARTSHEET_ACCESS_TOKEN"),
		BaseURL:     "https://api.smartsheet.com/2.0",
	}
}

// GetSheet returns a sheet by ID
func (c *SmartsheetClient) GetSheet(sheetID int) (Sheet, error) {
	url := fmt.Sprintf("%s/sheets/%d", c.BaseURL, sheetID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}

	var sheet Sheet
	err = json.Unmarshal(body, &sheet)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}

	return sheet, nil
}
