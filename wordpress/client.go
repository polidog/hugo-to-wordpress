package wordpress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL      string
	Username string
	Password string
}

func NewClient(url, username, password string) *Client {
	return &Client{
		URL:      url,
		Username: username,
		Password: password,
	}
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}

	// Read response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if len(bodyBytes) == 0 {
		return fmt.Errorf("empty response from server when getting tag ID for '%s'", resp.Request.URL)
	}

	// Unmarshal JSON response
	var jsonResp map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &jsonResp); err != nil {
		return fmt.Errorf("failed to parse response JSON: %v", err)
	}

	// Extract error message
	message := "unknown error"
	if msg, ok := jsonResp["message"]; ok {
		message = fmt.Sprintf("%v", msg)
	}

	return fmt.Errorf("failed to create post with status code %d: %s", resp.StatusCode, message)
}
