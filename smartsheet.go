//Client that will be used by other functions to interact with the Smartsheet API
package smartsheet

import (
	"os"

	"github.com/rs/zerolog/log"
)

// SmartsheetClient struct
type SmartsheetClient struct {
	AccessToken string

	// BaseURL is the base URL for the Smartsheet API
	BaseURL string
}

// NewClient creates a new SmartsheetClient
func NewClient() *SmartsheetClient {
	log.Info().Msg("creating new client")
	return &SmartsheetClient{
		AccessToken: os.Getenv("SMARTSHEET_ACCESS_TOKEN"),
		BaseURL:     "https://api.smartsheet.com/2.0",
	}
}

type Hyperlink struct {
	REPORT_ID int    `json:"reportId"`
	SHEET_ID  int    `json:"sheetId"`
	URL       string `json:"url"`
	SIGGHT_ID int    `json:"sightId"`
}

type Image struct {
	ALT_TEXT string `json:"altText"`
	HEIGHT   int    `json:"height"`
	ID       int    `json:"id"`
	WIDTH    int    `json:"width"`
}

type CELL_LINK struct {
	COLUMN_ID  int    `json:"columnId"`
	ROW_ID     int    `json:"rowId"`
	SHEET_ID   int    `json:"sheetId"`
	SHEET_NAME string `json:"sheetName"`
	STATUS     string `json:"status"`
}

type OBJECT_VALUE struct {
	OBJECT_TYPE string `json:"objectType"`
}

type Cell struct {
	COLUMN_ID           int          `json:"columnId"`
	COLUMN_TYPE         string       `json:"columnType"`
	CONDITIONAL_FORMAT  string       `json:"conditionalFormat"`
	DISPLAY_VALUE       string       `json:"displayValue"`
	FORMAT              string       `json:"format"`
	FORMULA             string       `json:"formula"`
	HYPERLINK           Hyperlink    `json:"hyperlink"`
	IMAGE               Image        `json:"image"`
	LINK_IN_FROM_CELL   CELL_LINK    `json:"linkInFromCell"`
	LINK_OUT_TO_CELL    []CELL_LINK  `json:"linkOutToCell"`
	OBJECT_VALUE        OBJECT_VALUE `json:"objectValue"`
	OVERRIDE_VALIDATION bool         `json:"overrideValidation"`
	STRICT              bool         `json:"strict"`
	VALUE               string       `json:"value"`
}

type Column struct {
	ID    int    `json:"id"`
	INDEX int    `json:"index"`
	TITLE string `json:"title"`
	TYPE  string `json:"type"`
}

type Attachment struct {
	ATTACHMENT_TYPE string `json:"attachmentType"`
	NAME            string `json:"name"`
	URL             string `json:"url"`
	ID              int    `json:"id"`
	PARENT_ID       int    `json:"parentId"`
	PARENT_TYPE     string `json:"parentType"`
	MIME_TYPE       string `json:"mimeType"`
	SIZE            int    `json:"size"`
}

type Row struct {
	ID          int          `json:"id"`
	SHEET_ID    int          `json:"sheetId"`
	ATTACHMENTS []Attachment `json:"attachments"`
	COLUMNS     []Column     `json:"columns"`
	CELLS       []Cell       `json:"cells"`
	SIBLING_ID  int          `json:"siblingId"`
	EXPADED     bool         `json:"expanded"`
	ROW_NUMBER  int          `json:"rowNumber"`
	PERMALINK   string       `json:"permalink"`
}
