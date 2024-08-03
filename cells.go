package smartsheet

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
	VIRTUAL_COLUMN_ID   int          `json:"virtualColumnId,omitempty"`
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
