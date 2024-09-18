package smartsheet

type Callback struct {
	Nonce         string  `json:"nonce"`
	Timestamp     string  `json:"timestamp"`
	WebhookId     int     `json:"webhookId"`
	Scope         string  `json:"scope"`
	ScopeObjectId int     `json:"scopeObjectId"`
	Events        []Event `json:"events"`
}

type Event struct {
	ObjectType string `json:"objectType"`
	EventType  string `json:"eventType"`
	Id         int    `json:"id,omitempty"`
	UserId     int    `json:"userId,omitempty"`
	Timestamp  string `json:"timestamp"`
	RowId      int    `json:"rowId,omitempty"`
	ColumnId   int    `json:"columnId,omitempty"`
}
