package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mitch292/gimmeplan/utils"
)

// RequestBody is used to marshal our text to json
type RequestBody struct {
	Text string `json:"text"`
}

// Send will submit a given content message to a given webhook url
func Send(webhookURL string, content []byte) {
	slackBody, err := json.Marshal(RequestBody{Text: utils.RemoveRefreshData(string(content))})
	if err != nil {
		log.Fatalf("There was a problem creating the json body: %s\n", err)
	}

	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		log.Fatalf("There was a problem creating the slack webhook request: %s\n", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("There was a problem submitting the slack webhook request %s\n", err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		log.Fatalf("We did not get an okay response back from slack: %s\n", err)
	}

}
