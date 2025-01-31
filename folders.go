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
	WorkspaceId  int
}

func Folders(options ...func(*folder_request)) *folder_request {
	d := &folder_request{}
	for _, option := range options {
		option(d)
	}
	return d
}

func Folder(name string) func(*folder_request) {
	return func(d *folder_request) {
		d.Name = name
	}
}
func Folder_In_Workspace(workspaceId int) func(*folder_request) {
	return func(d *folder_request) {
		d.WorkspaceId = workspaceId
	}
}
func Folder_In_ParentFolder(parentFolder int) func(*folder_request) {
	return func(d *folder_request) {
		d.ParentFolder = parentFolder
	}
}

func (c *SmartsheetClient) Create_Folder(name string) (folder, error) {
	req := Folders(Folder(name))
	f, err := c.PostFolder(req)
	if err != nil {
		log.Fatal().Msgf("error creating folder: %s", err)
	}
	return f, nil
}

func (c *SmartsheetClient) Create_Folder_In_Workspace(name string, workspaceId int) (folder, error) {
	req := Folders(Folder(name), Folder_In_Workspace(workspaceId))
	f, err := c.PostFolder(req)
	if err != nil {
		log.Fatal().Msgf("error creating folder: %s", err)
	}
	return f, nil
}
func (c *SmartsheetClient) Create_Folder_In_ParentFolder(name string, parentFolder int) (folder, error) {
	req := Folders(Folder(name), Folder_In_ParentFolder(parentFolder))
	f, err := c.PostFolder(req)
	if err != nil {
		log.Fatal().Msgf("error creating folder: %s", err)
	}
	return f, nil
}

// Create_Folder creates a folder
func (c *SmartsheetClient) PostFolder(req *folder_request) (folder, error) {
	log.Info().Msgf("creating folder with name: %s", req.Name)
	url := fmt.Sprintf("%s/folders", c.BaseURL)
	data := fmt.Sprintf(`{"name":"%s", "parentId": %d}`, req.Name, req.ParentFolder)
	body, err := c.Post_Call(url, data)

	var f folder
	err = json.Unmarshal(body, &f)
	if err != nil {
		return folder{}, err
	}

	return f, nil
}
