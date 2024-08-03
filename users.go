package smartsheet

type User struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
