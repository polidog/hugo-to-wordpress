package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) uploadImage(imgPath string) (int, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(imgPath))
	if err != nil {
		return 0, err
	}
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/wp-json/wp/v2/media", c.URL), body)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to upload image: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return int(result["id"].(float64)), nil
}

func (c *Client) getImageURLByID(id int) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wp-json/wp/v2/media/%d", c.URL, id), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get image URL, status code: %d", resp.StatusCode)
	}

	var mediaData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&mediaData); err != nil {
		return "", err
	}

	imgURL, ok := mediaData["source_url"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract image URL from media data")
	}

	return imgURL, nil
}
