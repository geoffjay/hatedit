package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/geoffjay/hatedit/models"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

type section struct {
	SectionType string `json:"type"`
	Text        string `json:"text"`
}

type block struct {
	BlockType string  `json:"type"`
	Text      section `json:"text"`
}

type response struct {
	ResponseType string  `json:"response_type"`
	Blocks       []block `json:"blocks"`
}

func generateResponse(users models.Users) response {
	var blocks []block
	for _, user := range users {
		text := fmt.Sprintf("*%s* hates:", user.Name)
		blocks = append(blocks, block{
			BlockType: "section",
			Text: section{
				SectionType: "mrkdwn",
				Text:        text,
			},
		})
		list := ""
		var items []models.Item
		stmt := fmt.Sprintf("list_id = '%s'", user.List.ID)
		query := models.DB.Where(stmt)
		_ = query.Eager().All(&items)
		fmt.Printf("\n\n%+v\n\n", items)
		for _, item := range items {
			list = fmt.Sprintf("%s\u2022 %s\n", list, item.Text)
		}
		blocks = append(blocks, block{
			BlockType: "section",
			Text: section{
				SectionType: "mrkdwn",
				Text:        list,
			},
		})
	}
	r := response{
		ResponseType: "in_channel",
		Blocks:       blocks,
	}
	return r
}

// CommandHandler default implementation.
func CommandHandler(c buffalo.Context) error {
	users := models.Users{}
	if err := models.DB.Eager().All(&users); err != nil {
		return err
	}
	resp := generateResponse(users)
	b, err := json.Marshal(resp)
	if err != nil {
		_ = c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "something happened while handling the command"}))
	}

	body := bytes.NewBuffer(b)
	responseUrl := c.Params().Get("response_url")
	slackResp, err := http.Post(responseUrl, "application/json", body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer slackResp.Body.Close()

	return nil
}
