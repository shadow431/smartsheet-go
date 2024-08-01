package smartsheet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Sheet struct
type Sheet struct {
	ID        int `json:"id"`
	OWNDER_ID int `json:"ownerId,omitempty"`
}

// GetSheet returns a sheet by ID
func (c *SmartsheetClient) GetSheet(sheetID int) (Sheet, error) {
	log.Info().Msgf("getting sheet with ID: %d", sheetID)
	url := fmt.Sprintf("%s/sheets/%d", c.BaseURL, sheetID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	log.Info().Msgf("sending request %s", req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return Sheet{ID: sheetID}, fmt.Errorf("error getting sheet: %s", body)
	}
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
