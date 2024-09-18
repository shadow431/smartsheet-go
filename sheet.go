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
