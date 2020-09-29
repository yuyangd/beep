package beep

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://api.newrelic.com/v2"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

type DMarker struct {
	Deployment Deploy `json:"deployment"`
}

type Deploy struct {
	User        string `json:"user"`
	Description string `json:"description"`
	Changelog   string `json:"changelog"`
	Revision    string `json:"revision"`
	Timestamp   string `json:"timestamp,omitempty"`
}

// SetDMarker creates a deployment marker after each deployment
func (s *Client) SetDMarker(appID string, dMarker *DMarker) error {
	endpoint := fmt.Sprintf("%s/applications/%s/deployments.json", baseURL, appID)
	j, err := json.Marshal(&dMarker)
	if err != nil {
		return err
	}
	log.Printf("Deployment descriptions, %v", string(j))

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	r, err := s.doRequest(req)
	var deployRep DMarker
	err = json.Unmarshal(r, &deployRep)
	log.Println(fmt.Sprintf("Deployment Mark set at: %v", deployRep.Deployment.Timestamp))
	return err

}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-Api-Key", s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
