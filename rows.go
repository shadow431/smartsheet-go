package smartsheet

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
