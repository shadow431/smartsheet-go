package smartsheet

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

//Structs for discussions
type discussions struct {
	Discussions []Discussion `json:"data,omitempty"`
	Page_Number int          `json:"pageNumber,omitempty"`
	Page_Size   int          `json:"pageSize,omitempty"`
	Total_Pages int          `json:"totalPages,omitempty"`
	Total_Count int          `json:"totalCount,omitempty"`
}

type Discussion struct {
	Id                  int          `json:"id"`
	Access_Level        string       `json:"accessLevel,omitempty"`
	Comments            []Comment    `json:"comments,omitempty"`
	Comment_Attachments []Attachment `json:"commentAttachments,omitempty"`
	Comment_Count       int          `json:"commentCount,omitempty"`
	Created_By          User         `json:"createdBy,omitempty"`
	Last_comment_At     string       `json:"lastCommentAt,omitempty"`
	Last_Commented_User User         `json:"lastCommentedUser,omitempty"`
	Parent_Id           int          `json:"parentId,omitempty"`
	Parent_Type         string       `json:"parentType,omitempty"`
	Read_Only           bool         `json:"readOnly,omitempty"`
	Title               string       `json:"title,omitempty"`
}

type discussion_request struct {
	Id       int
	Sheet_Id int
	Row_Id   int
	Includes []string
}

type discussion_response struct {
	Smartsheet_Response
	Result Discussion `json:"result,omitempty"`
}

//Request Options

func Discussions(options ...func(*discussion_request)) *discussion_request {
	d := &discussion_request{}
	for _, option := range options {
		option(d)
	}
	return d
}

func Discussions_With_Id(id int) func(*discussion_request) {
	return func(d *discussion_request) {
		d.Id = id
	}
}

func Discussions_With_Row_Id(rowId int) func(*discussion_request) {
	return func(d *discussion_request) {
		d.Row_Id = rowId
	}
}

func Discussions_With_Sheet_Id(sheetId int) func(*discussion_request) {
	return func(d *discussion_request) {
		d.Sheet_Id = sheetId
	}
}

func Discussions_With_Includes(includes ...string) func(*discussion_request) {
	return func(d *discussion_request) {
		d.Includes = includes
	}
}

//Request Modifiers

func (c *SmartsheetClient) ListSheetDiscussions(sheetId int, includes ...string) (discussions, error) {
	log.Info().Msg("listing Sheet discussions")
	discussion_request := Discussions(Discussions_With_Sheet_Id(sheetId), Discussions_With_Includes(includes...))

	return c.getDiscussions(discussion_request)
}

func (c *SmartsheetClient) ListRowDiscussions(sheetId int, rowId int, includes ...string) (discussions, error) {
	log.Info().Msg("listing Row discussions")
	discussion_request := Discussions(Discussions_With_Sheet_Id(sheetId), Discussions_With_Row_Id(rowId), Discussions_With_Includes(includes...))

	return c.getDiscussions(discussion_request)
}

func (c *SmartsheetClient) GetDiscussion(sheetId int, discussionId int, includes ...string) (discussions, error) {
	log.Info().Msg("getting discussion by ID")
	discussion_request := Discussions(Discussions_With_Id(discussionId), Discussions_With_Sheet_Id(sheetId))

	return c.getDiscussions(discussion_request)
}

func (c *SmartsheetClient) CreatSheetDiscussion(sheetId int, comment string) (Discussion, error) {
	log.Info().Msg("getting discussions")
	discussion_request := Discussions(Discussions_With_Sheet_Id(sheetId))

	return c.PostDiscussion(discussion_request, comment)
}

func (c *SmartsheetClient) CreatRowDiscussion(sheetId int, rowId int, comment string) (Discussion, error) {
	log.Info().Msg("getting discussions")
	discussion_request := Discussions(Discussions_With_Sheet_Id(sheetId), Discussions_With_Row_Id(rowId))

	return c.PostDiscussion(discussion_request, comment)
}

//Request Functions

func (c *SmartsheetClient) getDiscussions(discussion *discussion_request) (discussions, error) {
	log.Info().Msg("getting discussions")
	url := fmt.Sprintf("%s/sheets/%d", c.BaseURL, discussion.Sheet_Id)
	if discussion.Row_Id != 0 {
		url = fmt.Sprintf("%s/rows/%d", url, discussion.Row_Id)
	}

	url = fmt.Sprintf("%s/discussions", url)
	if discussion.Includes != nil {
		url = fmt.Sprintf("%s?include=%s", url, discussion.Includes[0])
	}

	var discussions discussions
	body, err := c.Get_Call(url)
	if err != nil {
		return discussions, err
	}

	log.Debug().Msgf("body: %s", body)
	err = json.Unmarshal(body, &discussions)
	if err != nil {
		return discussions, err
	}

	return discussions, nil
}

func (c *SmartsheetClient) PostDiscussion(req *discussion_request, comment string) (Discussion, error) {
	log.Info().Msg("getting discussions")
	url := fmt.Sprintf("%s/sheets/%d", c.BaseURL, req.Sheet_Id)
	if req.Row_Id != 0 {
		url = fmt.Sprintf("%s/rows/%d", url, req.Row_Id)
	}
	url = fmt.Sprintf("%s/discussions", url)
	data := fmt.Sprintf(`{"comment": {"text":"%s"}}`, comment)
	log.Debug().Msgf("data: %s", data)
	var discussion Discussion
	body, err := c.Post_Call(url, string(data))
	if err != nil {
		return discussion, err
	}

	var smar_body discussion_response
	err = json.Unmarshal(body, &smar_body)
	if err != nil {
		log.Debug().Msgf("Error unmarshalling post_call body: %s", body)
		return discussion, err
	}
	log.Debug().Msgf("Result: %s", smar_body.Result)
	if smar_body.ErrorCode != 0 {
		return discussion, fmt.Errorf("error creating discussion - %d: %s - %s", smar_body.ErrorCode, smar_body.RefId, smar_body.Message)
	}

	return smar_body.Result, nil
}
