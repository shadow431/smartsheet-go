package smartsheet

type Comment struct {
	Attachments   []Attachment `json:"attachments,omitempty"`
	Created_At    string       `json:"createdAt,omitempty"`
	Created_By    User         `json:"createdBy,omitempty"`
	Modified_At   string       `json:"modifiedAt,omitempty"`
	Discussion_Id int          `json:"discussionId,omitempty"`
	Id            int          `json:"id"`
	Text          string       `json:"text,omitempty"`
}
