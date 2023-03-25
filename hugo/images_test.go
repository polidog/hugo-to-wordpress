package hugo_test

import (
	"testing"

	"github.com/polidog/hugo-to-wordpress/hugo"
)

func TestExtractImagesFromContent(t *testing.T) {
	content := `![Alt text 1](/path/to/image1.jpg)
This is some text.
![Alt text 2](/path/to/image2.jpg)`

	expectedImages := []string{"/path/to/image1.jpg", "/path/to/image2.jpg"}

	images := hugo.ExtractImagesFromContent(content)

	if len(images) != len(expectedImages) {
		t.Fatalf("Expected %d images, got %d", len(expectedImages), len(images))
	}

	for i, img := range images {
		if img != expectedImages[i] {
			t.Errorf("Expected image '%s', got '%s'", expectedImages[i], img)
		}
	}
}
