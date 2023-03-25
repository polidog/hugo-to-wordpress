package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *Client) getOrCreateCategory(name string) (int, error) {
	var categories []Category
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wp-json/wp/v2/categories?slug=%s", c.URL, slug), nil)
	if err != nil {
		return 0, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&categories)
	if err != nil {
		return 0, err
	}

	if len(categories) > 0 {
		return categories[0].ID, nil
	}

	newCategory := Category{Name: name, Slug: slug}
	categoryJson, _ := json.Marshal(newCategory)

	req, err = http.NewRequest("POST", fmt.Sprintf("%s/wp-json/wp/v2/categories", c.URL), bytes.NewReader(categoryJson))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if err := checkResponse(res); err != nil {
		return 0, err
	}

	err = json.NewDecoder(res.Body).Decode(&newCategory)
	if err != nil {
		return 0, err
	}

	return newCategory.ID, nil
}

func (c *Client) getOrCreateTag(name string) (int, error) {
	var tags []Tag
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wp-json/wp/v2/tags?slug=%s", c.URL, slug), nil)
	if err != nil {
		return 0, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&tags)
	if err != nil {
		return 0, err
	}

	if len(tags) > 0 {
		return tags[0].ID, nil
	}

	newTag := Tag{Name: name, Slug: slug}
	tagJson, _ := json.Marshal(newTag)
	req, err = http.NewRequest("POST", fmt.Sprintf("%s/wp-json/wp/v2/tags", c.URL), bytes.NewReader(tagJson))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&newTag)
	if err := checkResponse(res); err != nil {
		return 0, err
	}

	return newTag.ID, nil
}
