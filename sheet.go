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
	Name        string
	columns     []Column
	FolderId    int
	WorkspaceId int
}

type sheet_response struct {
	Smartsheet_Response
	Result Sheet `json:"result,omitempty"`
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

func Sheets(options ...func(*sheet_request)) *sheet_request {
	d := &sheet_request{}
	for _, option := range options {
		option(d)
	}
	return d
}

func Sheet_In_Folder(folderId int) func(*sheet_request) {
	return func(d *sheet_request) {
		d.FolderId = folderId
	}
}

func Sheet_In_Workspace(workspaceId int) func(*sheet_request) {
	return func(d *sheet_request) {
		d.WorkspaceId = workspaceId
	}
}
func Sheet_Columns(columns []Column) func(*sheet_request) {
	return func(d *sheet_request) {
		d.columns = columns
	}
}

func Sheet_Name(name string) func(*sheet_request) {
	return func(d *sheet_request) {
		d.Name = name
	}
}
func (c *SmartsheetClient) CreateSheet(name string, columns []Column) (Sheet, error) {
	log.Info().Msg("createing Sheet")
	sheet_request := Sheets(Sheet_Columns(columns), Sheet_Name(name))
	return c.PostSheet(sheet_request)
}

func (c *SmartsheetClient) CreateSheetInFolder(name string, columns []Column, folder_id int) (Sheet, error) {
	log.Info().Msg("createing Sheet in Folder")
	sheet_request := Sheets(Sheet_Columns(columns), Sheet_In_Folder(folder_id), Sheet_Name(name))
	return c.PostSheet(sheet_request)
}

func (c *SmartsheetClient) CreateSheetInWorkspace(name string, columns []Column, workspace_id int) (Sheet, error) {
	log.Info().Msg("createing Sheet in Workspace")
	sheet_request := Sheets(Sheet_Columns(columns), Sheet_In_Workspace(workspace_id), Sheet_Name(name))
	return c.PostSheet(sheet_request)
}

func (c *SmartsheetClient) PostSheet(req *sheet_request) (Sheet, error) {
	log.Info().Msgf("creating sheet with name: %s", req.Name)
	url := fmt.Sprintf("%s", c.BaseURL)
	if req.WorkspaceId != 0 {
		url = fmt.Sprintf("%s/workspaces/%d", url, req.WorkspaceId)
	}
	if req.FolderId != 0 {
		url = fmt.Sprintf("%s/folders/%d", url, req.FolderId)
	}
	url = fmt.Sprintf("%s/sheets", url)

	columns, err := json.Marshal(req.columns)

	data := fmt.Sprintf(`{"name":"%s","columns":%s}`, req.Name, columns)

	body, err := c.Post_Call(url, data)

	var sheet sheet_response
	err = json.Unmarshal(body, &sheet)
	if err != nil {
		return Sheet{}, err
	}

	return sheet.Result, nil
}
