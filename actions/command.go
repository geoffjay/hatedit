package actions

import (
	"bytes"
	"encoding/json"
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

// CommandHandler default implementation.
func CommandHandler(c buffalo.Context) error {
	resp := response{
		ResponseType: "in_channel",
		Blocks: []block{
			{
				BlockType: "section",
				Text: section{
					SectionType: "mrkdwn",
					Text:        "*@jefe* hates:",
				},
			},
			{
				BlockType: "section",
				Text: section{
					SectionType: "mrkdwn",
					Text:        "\u2022 foo\n\u2022 bar\n\u2022 baz",
				},
			},
		},
	}
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
