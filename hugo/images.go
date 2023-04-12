package hugo

import (
	"regexp"
)

var (
	imgPattern    = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	figurePattern = regexp.MustCompile(`{{<\s*figure\s*src=["']([^"']+)["'].*>}}`)
)

// ExtractImagesFromContent extracts image references from a markdown content
func ExtractImagesFromContent(content string) []string {
	var images []string

	// Extract images from standard markdown image syntax
	matches := imgPattern.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		images = append(images, match[2])
	}

	// Extract images from figure shortcode syntax
	figureMatches := figurePattern.FindAllStringSubmatch(content, -1)
	for _, match := range figureMatches {
		images = append(images, match[1])
	}

	return images
}
