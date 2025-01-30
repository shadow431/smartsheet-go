package smartsheet

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Sheet struct
type Sheet struct {
	ID        int      `json:"id"`
	OWNDER_ID int      `json:"ownerId,omitempty"`
	Columns   []Column `json:"columns,omitempty"`
	Rows      []Row    `json:"rows,omitempty"`
}

type sheet_request struct {
	Name         string
	folder_id    int
	workspace_id int
	fromId       int
}

// GetSheet returns a sheet by ID
func (c *SmartsheetClient) GetSheet(sheetID int) (Sheet, error) {
	log.Info().Msgf("getting sheet with ID: %d", sheetID)
	url := fmt.Sprintf("%s/sheets/%d", c.BaseURL, sheetID)
	body, err := c.Get_Call(url)

	var sheet Sheet
	err = json.Unmarshal(body, &sheet)
	if err != nil {
		return Sheet{
			ID: sheetID,
		}, err
	}

	return sheet, nil
}

func (c *SmartsheetClient) Create_Sheet(req *sheet_request, sheet_name string, folder_id string) (Sheet, error) {
	log.Info().Msgf("creating sheet with name: %s", sheet_name)
	url := fmt.Sprintf("%s", c.BaseURL)
	if folder != "" {
		url := fmt.Sprintf("%s/folders/%d", url, req.folder_id)
	}
	url := fmt.Sprintf("%s/sheets", url)

	body, err := c.Post_Call(url, req)

	var sheet Sheet
	err = json.Unmarshal(body, &sheet)
	if err != nil {
		return Sheet{}, err
	}

	return sheet, nil
}
