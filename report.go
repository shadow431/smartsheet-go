package smartsheet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Report struct {
	ID          int          `json:"id"`
	OWNDER_ID   int          `json:"ownerId,omitempty"`
	ATTACHMENTS []Attachment `json:"attachments,omitempty"`
	NAME        string       `json:"name,omitempty"`
	OWNER       string       `json:"owner,omitempty"`
	PERMALINK   string       `json:"permalink,omitempty"`
	ROWS        []Row        `json:"rows,omitempty"`
}

// GetSheet returns a report by ID
func (c *SmartsheetClient) GetReport(reportID int) (Report, error) {
	log.Info().Msgf("getting report with ID: %d", reportID)
	url := fmt.Sprintf("%s/reports/%d", c.BaseURL, reportID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Report{
			ID: reportID,
		}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	log.Info().Msgf("sending request %s", req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Report{
			ID: reportID,
		}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return Report{ID: reportID}, fmt.Errorf("error getting report: %s", body)
	}
	if err != nil {
		return Report{
			ID: reportID,
		}, err
	}

	var report Report
	err = json.Unmarshal(body, &report)
	if err != nil {
		return Report{
			ID: reportID,
		}, err
	}

	return report, nil
}
