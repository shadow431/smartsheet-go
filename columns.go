package smartsheet

type Column struct {
	ID    int    `json:"id,omitempty"`
	INDEX int    `json:"index,omitempty"`
	TITLE string `json:"title,omitempty"`
	TYPE  string `json:"type,omitempty"`
}
