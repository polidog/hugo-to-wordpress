package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/polidog/hugo-to-wordpress/hugo"
)

func (c *Client) MigratePosts(posts []*hugo.Post) error {
	for _, post := range posts {
		fromHugo := *post
		if err := c.createPost(&fromHugo); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) createPost(post *hugo.Post) error {
	// Upload featured image
	var featuredMediaID int
	if post.Eyecatch != "" {
		imgID, err := c.uploadImage(post.Eyecatch)
		if err != nil {
			return err
		}
		featuredMediaID = imgID
	}

	contentImages := hugo.ExtractImagesFromContent(post.Content)
	imageURLMap := make(map[string]string)
	for _, imgPath := range contentImages {
		imgID, err := c.uploadImage(imgPath)
		if err != nil {
			return err
		}
		newImgURL, err := c.getImageURLByID(imgID)
		if err != nil {
			return err
		}
		imageURLMap[imgPath] = newImgURL
	}

	// Replace old image URLs with new ones
	updatedContent := post.Content
	for oldURL, newURL := range imageURLMap {
		updatedContent = strings.ReplaceAll(updatedContent, oldURL, newURL)
	}

	// Convert category names to IDs
	categoryIDs := make([]int, len(post.Categories))
	for i, categoryName := range post.Categories {
		id, err := c.getOrCreateCategory(categoryName)
		if err != nil {
			return fmt.Errorf("failed to get category ID for '%s': %v", categoryName, err)
		}
		categoryIDs[i] = id
	}

	// Convert tag names to IDs
	tagIDs := make([]int, len(post.Tags))
	for i, tagName := range post.Tags {
		id, err := c.getOrCreateTag(tagName)
		if err != nil {
			return fmt.Errorf("failed to get tag ID for '%s': %v", tagName, err)
		}
		tagIDs[i] = id
	}

	// Prepare post data
	date, err := time.Parse("2006-01-02T15:04:05", post.Date)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	// Prepare post data
	postData := map[string]interface{}{
		"title":          post.Title,
		"content":        post.Content,
		"categories":     categoryIDs,
		"tags":           tagIDs,
		"featured_media": featuredMediaID,
		"status":         "publish",
		"date":           date.Format(time.RFC3339),
	}

	// Create post
	jsonData, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/wp-json/wp/v2/posts", c.URL), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
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

	return nil
}
