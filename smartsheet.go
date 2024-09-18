//Client that will be used by other functions to interact with the Smartsheet API
package smartsheet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

// SmartsheetClient struct
type SmartsheetClient struct {
	AccessToken string

	// BaseURL is the base URL for the Smartsheet API
	BaseURL string
}

/*
//Opting for a string instead of a struct for now
type Includes struct {
	Attachments bool
	Comments    bool
}
*/

type Smartsheet_Response struct {
	Version    int    `json:"version,omitempty"`
	Message    string `json:"message,omitempty"`
	ResultCode int    `json:"resultCode,omitempty"`
	ErrorCode  int    `json:"errorCode,omitempty"`
	RefId      string `json:"refId,omitempty"`
}

// NewClient creates a new SmartsheetClient
func NewClient() *SmartsheetClient {
	log.Info().Msg("creating new client")
	return &SmartsheetClient{
		AccessToken: os.Getenv("SMARTSHEET_ACCESS_TOKEN"),
		BaseURL:     "https://api.smartsheet.com/2.0",
	}
}

func (c *SmartsheetClient) Call(url string, method string, data string) ([]byte, error) {
	log.Info().Msgf("calling %s with method %s", url, method)
	var req *http.Request
	var err error
	if data == "" {
		log.Debug().Msg("no data provided")
		req, err = http.NewRequest(method, url, nil)
	} else {
		log.Debug().Msgf("data: %s", data)
		jsonBody := []byte(data)
		bodyReader := bytes.NewReader(jsonBody)
		req, err = http.NewRequest(method, url, bodyReader)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	log.Info().Msgf("sending request %s", req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Info().Msgf("response: %s", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error calling %s: %s", url, body)
	}
	if err != nil {
		return nil, err
	}

	return body, nil

}

func (c *SmartsheetClient) Get_Call(url string) ([]byte, error) {
	log.Debug().Msg("Making Get Smartsheet Call")
	body, err := c.Call(url, "GET", "")
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *SmartsheetClient) Post_Call(url string, data string) ([]byte, error) {
	log.Debug().Msg("Making POst Smartsheet Call")
	body, err := c.Call(url, "POST", data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *SmartsheetClient) Put_Call(url string, data string) ([]byte, error) {
	log.Debug().Msg("Making Put Smartsheet Call")
	body, err := c.Call(url, "PUT", data)
	if err != nil {
		return nil, err
	}

	return body, nil
}
