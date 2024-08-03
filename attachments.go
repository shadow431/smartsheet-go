package smartsheet

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
