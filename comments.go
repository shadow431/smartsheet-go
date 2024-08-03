package smartsheet

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

//Comment Structs
type Comment struct {
	Attachments   []Attachment `json:"attachments,omitempty"`
	Created_At    string       `json:"createdAt,omitempty"`
	Created_By    User         `json:"createdBy,omitempty"`
	Modified_At   string       `json:"modifiedAt,omitempty"`
	Discussion_Id int          `json:"discussionId,omitempty"`
	Id            int          `json:"id"`
	Text          string       `json:"text,omitempty"`
}

type comment_request struct {
	Id            int
	Discussion_Id int
	Sheet_Id      int
	Row_Id        int
}

type comment_response struct {
	Smartsheet_Response
	Result Comment `json:"result,omitempty"`
}

//Request Modifiers
func Comments(options ...func(*comment_request)) *comment_request {
	d := &comment_request{}
	for _, option := range options {
		option(d)
	}
	return d
}

func Comments_With_Id(id int) func(*comment_request) {
	return func(d *comment_request) {
		d.Id = id
	}
}

func Comments_With_Discussion_Id(id int) func(*comment_request) {
	return func(d *comment_request) {
		d.Discussion_Id = id
	}
}

func Comments_With_Sheet_Id(sheetId int) func(*comment_request) {
	return func(d *comment_request) {
		d.Sheet_Id = sheetId
	}
}

//Request Modifiers
func (c *SmartsheetClient) GetComment(sheetId int, commentId int) (Comment, error) {
	log.Info().Msg("getting comment")
	comment_request := Comments(Comments_With_Id(commentId), Comments_With_Sheet_Id(sheetId))

	return c.getComment(comment_request)
}

func (c *SmartsheetClient) CreateComment(sheetId int, disccusion_id int, text string) (Comment, error) {
	log.Info().Msg("creating comment")
	url := fmt.Sprintf("%s/sheets/%d/discussions/%d/comments", c.BaseURL, sheetId, disccusion_id)
	data := fmt.Sprintf(`{"text":"%s"}`, text)

	body, err := c.Post_Call(url, data)
	if err != nil {
		return Comment{}, err
	}
	var comment Comment
	var smar_body comment_response
	err = json.Unmarshal(body, &smar_body)
	if err != nil {
		return comment, err
	}
	if smar_body.ErrorCode != 0 {
		return comment, fmt.Errorf("error creating comment - %d: %s - %s", smar_body.ErrorCode, smar_body.RefId, smar_body.Message)
	}
	//var discussion_id int
	return smar_body.Result, nil
}

//Request Functions
func (c *SmartsheetClient) getComment(request *comment_request) (Comment, error) {
	log.Info().Msg("getting comment")
	url := fmt.Sprintf("%s/sheets/%d/comments/%d", c.BaseURL, request.Sheet_Id, request.Id)

	var comment Comment
	body, err := c.Get_Call(url)
	if err != nil {
		return comment, err
	}

	log.Debug().Msgf("body: %s", body)
	err = json.Unmarshal(body, &comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
