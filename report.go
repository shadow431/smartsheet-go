package smartsheet

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Report struct {
	ID          int             `json:"id"`
	OWNDER_ID   int             `json:"ownerId,omitempty"`
	ATTACHMENTS []Attachment    `json:"attachments,omitempty"`
	NAME        string          `json:"name,omitempty"`
	OWNER       string          `json:"owner,omitempty"`
	PERMALINK   string          `json:"permalink,omitempty"`
	ROWS        []Row           `json:"rows,omitempty"`
	COLUMNS     []Report_Column `json:"columns,omitempty"`
}

type Report_Column struct {
	ID    int    `json:"virtualId"`
	INDEX int    `json:"index,omitempty"`
	TITLE string `json:"title,omitempty"`
	TYPE  string `json:"type,omitempty"`
}

// GetSheet returns a report by ID
func (c *SmartsheetClient) GetReport(reportID int) (Report, error) {
	log.Info().Msgf("getting report with ID: %d", reportID)
	url := fmt.Sprintf("%s/reports/%d", c.BaseURL, reportID)
	body, err := c.Get_Call(url)

	var report Report
	log.Debug().Msgf("body: %s", body)
	err = json.Unmarshal(body, &report)
	if err != nil {
		return Report{
			ID: reportID,
		}, err
	}

	return report, nil
}
