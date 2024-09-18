package smartsheet

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Column struct {
	ID      int      `json:"id,omitempty"`
	INDEX   int      `json:"index,omitempty"`
	TITLE   string   `json:"title,omitempty"`
	TYPE    string   `json:"type,omitempty"`
	OPTIONS []string `json:"options,omitempty"`
}

type column_response struct {
	Smartsheet_Response
	Result Column `json:"result,omitempty"`
}

type column_get struct {
	Page_number int      `json:"pageNumber,omitempty"`
	Page_size   int      `json:"pageSize,omitempty"`
	Total_pages int      `json:"totalPages,omitempty"`
	Total_count int      `json:"totalCount,omitempty"`
	Columns     []Column `json:"data,omitempty"`
}

//Request Functions
func (c *SmartsheetClient) GetColumns(sheetID int) ([]Column, error) {
	log.Info().Msgf("getting sheet columns with sheet ID: %d", sheetID)
	url := fmt.Sprintf("%s/sheets/%d/columns", c.BaseURL, sheetID)
	body, err := c.Get_Call(url)

	var column_get column_get
	//log.Debug().Msgf("body: %s", body)
	err = json.Unmarshal(body, &column_get)
	if err != nil {
		return nil, err
	}

	return column_get.Columns, nil
}

//update column
func (c *SmartsheetClient) UpdateColumn(sheetID int, columnID int, data string) (Column, error) {
	log.Info().Msgf("updating column with sheet ID: %d and column ID: %d", sheetID, columnID)
	url := fmt.Sprintf("%s/sheets/%d/columns/%d", c.BaseURL, sheetID, columnID)
	body, err := c.Put_Call(url, data)

	var column column_response
	err = json.Unmarshal(body, &column)
	if err != nil {
		return Column{}, err
	}

	return column.Result, nil
}
