package smartsheet

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type folder struct {
	ID        int    `json:"id"`
	favorite  bool   `json:"favorite,omitempty"`
	folders   []int  `json:"folders,omitempty"`
	name      string `json:"name,omitempty"`
	permalink string `json:"permalink,omitempty"`
	reports   []int  `json:"reports,omitempty"`
	sheets    []int  `json:"sheets,omitempty"`
	sights    []int  `json:"sights,omitempty"`
	templates bool   `json:"template,omitempty"`
}

type folder_request struct {
	Name         string
	ParentFolder int
}

// Create_Folder creates a folder
func (c *SmartsheetClient) Create_Folder(req *folder_request) (folder, error) {
	log.Info().Msgf("creating folder with name: %s", req.Name)
	url := fmt.Sprintf("%s/folders", c.BaseURL)

	body, err := c.Post_Call(url, req)

	var f folder
	err = json.Unmarshal(body, &f)
	if err != nil {
		return folder{}, err
	}

	return f, nil
}
