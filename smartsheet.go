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
	REPORT_ID int    `json:"reportId,omitempty"`
	SHEET_ID  int    `json:"sheetId,omitempty"`
	URL       string `json:"url,omitempty"`
	SIGGHT_ID int    `json:"sightId,omitempty"`
}

type Image struct {
	ALT_TEXT string `json:"altText,omitempty"`
	HEIGHT   int    `json:"height,	omitempty"`
	ID       string `json:"id,omitempty"`
	WIDTH    int    `json:"width,omitempty"`
}

type CELL_LINK struct {
	COLUMN_ID  int    `json:"columnId,omitempty"`
	ROW_ID     int    `json:"rowId,omitempty"`
	SHEET_ID   int    `json:"sheetId,omitempty"`
	SHEET_NAME string `json:"sheetName,omitempty"`
	STATUS     string `json:"status,omitempty"`
}

type OBJECT_VALUE struct {
	OBJECT_TYPE string `json:"objectType,omitempty"`
}

type Cell struct {
	COLUMN_ID           int          `json:"columnId,omitempty"`
	COLUMN_TYPE         string       `json:"columnType,omitempty"`
	CONDITIONAL_FORMAT  string       `json:"conditionalFormat,omitempty"`
	DISPLAY_VALUE       string       `json:"displayValue,omitempty"`
	FORMAT              string       `json:"format,omitempty"`
	FORMULA             string       `json:"formula,omitempty"`
	HYPERLINK           Hyperlink    `json:"hyperlink,omitempty"`
	IMAGE               Image        `json:"image,omitempty"`
	LINK_IN_FROM_CELL   CELL_LINK    `json:"linkInFromCell,omitempty"`
	LINK_OUT_TO_CELL    []CELL_LINK  `json:"linkOutToCell,omitempty"`
	OBJECT_VALUE        OBJECT_VALUE `json:"objectValue,omitempty"`
	OVERRIDE_VALIDATION bool         `json:"overrideValidation,omitempty"`
	STRICT              bool         `json:"strict,omitempty"`
	VALUE               interface{}  `json:"value,omitempty"`
}

type Column struct {
	ID    int    `json:"id,omitempty"`
	INDEX int    `json:"index,omitempty"`
	TITLE string `json:"title,omitempty"`
	TYPE  string `json:"type,omitempty"`
}

type Attachment struct {
	ATTACHMENT_TYPE string `json:"attachmentType,omitempty"`
	NAME            string `json:"name,omitempty"`
	URL             string `json:"url,omitempty"`
	ID              int    `json:"id,omitempty"`
	PARENT_ID       int    `json:"parentId,omitempty"`
	PARENT_TYPE     string `json:"parentType,omitempty"`
	MIME_TYPE       string `json:"mimeType,omitempty"`
	SIZE            int    `json:"size,omitempty"`
}

type Row struct {
	ID          int          `json:"id,omitempty"`
	SHEET_ID    int          `json:"sheetId,omitempty"`
	ATTACHMENTS []Attachment `json:"attachments,omitempty"`
	COLUMNS     []Column     `json:"columns,omitempty"`
	CELLS       []Cell       `json:"cells,omitempty"`
	SIBLING_ID  int          `json:"siblingId,omitempty"`
	EXPADED     bool         `json:"expanded,omitempty"`
	ROW_NUMBER  int          `json:"rowNumber,omitempty"`
	PERMALINK   string       `json:"permalink,omitempty"`
}
