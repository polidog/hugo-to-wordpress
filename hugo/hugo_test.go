package hugo_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/polidog/hugo-to-wordpress/hugo"
)

const samplePost = `
---
title: "Sample Post"
categories:
- Category1
- Category2
tags:
- Tag1
- Tag2
eyecatch: "/path/to/image.jpg"
---

# Sample Post

This is a sample post content.
`

func TestParseHugoPosts(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "hugo-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := ioutil.WriteFile(filepath.Join(tmpDir, "sample.md"), []byte(samplePost), 0644); err != nil {
		t.Fatalf("Failed to write sample post: %v", err)
	}

	posts, err := hugo.ParseHugoPosts(tmpDir)
	if err != nil {
		t.Fatalf("Error parsing Hugo posts: %v", err)
	}

	if len(posts) != 1 {
		t.Fatalf("Expected 1 post, got %d", len(posts))
	}

	post := posts[0]

	if post.Title != "Sample Post" {
		t.Errorf("Expected title 'Sample Post', got '%s'", post.Title)
	}

	expectedCategories := []string{"Category1", "Category2"}
	for i, category := range post.Categories {
		if category != expectedCategories[i] {
			t.Errorf("Expected category '%s', got '%s'", expectedCategories[i], category)
		}
	}

	expectedTags := []string{"Tag1", "Tag2"}
	for i, tag := range post.Tags {
		if tag != expectedTags[i] {
			t.Errorf("Expected tag '%s', got '%s'", expectedTags[i], tag)
		}
	}

	if post.Eyecatch != "/path/to/image.jpg" {
		t.Errorf("Expected eyecatch '/path/to/image.jpg', got '%s'", post.Eyecatch)
	}

	expectedContent := "# Sample Post\n\nThis is a sample post content."
	if post.Content != expectedContent {
		t.Errorf("Expected content '%s', got '%s'", expectedContent, post.Content)
	}
}
