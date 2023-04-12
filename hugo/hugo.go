package hugo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Post struct {
	Title      string   `yaml:"title"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
	Tags       []string `yaml:"tags"`
	Eyecatch   string   `yaml:"eyecatch"`
	Content    string
	ParsedDate time.Time
}

var (
	frontMatterRe = regexp.MustCompile(`(?s)(?m)^---\n(.+?)\n---\n(.+)$`)
)

func ParseHugoPosts(contentPath string) ([]*Post, error) {
	var posts []*Post

	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			post, err := parsePost(path)
			if err != nil {
				return err
			}

			posts = append(posts, post)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func parsePost(path string) (*Post, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	matches := frontMatterRe.FindSubmatch(content)
	if len(matches) != 3 {
		return nil, fmt.Errorf("failed to parse front matter and content from file: %s", path)
	}

	frontMatter := matches[1]
	postContent := strings.TrimSpace(string(matches[2]))

	var post Post
	if err := yaml.Unmarshal(frontMatter, &post); err != nil {
		return nil, err
	}

	post.Content = postContent

	dateStr := post.Date
	if dateStr != "" {
		parsedDate, err := Parse(dateStr)
		if err != nil {
			return nil, fmt.Errorf("EEEE failed to parse date: %v", err)
		}
		post.ParsedDate = parsedDate
	}

	return &post, nil
}
