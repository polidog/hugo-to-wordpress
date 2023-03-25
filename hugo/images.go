package hugo

import (
	"regexp"
)

var (
	imgPattern = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
)

// ExtractImagesFromContent extracts image references from a markdown content
func ExtractImagesFromContent(content string) []string {
	var images []string
	matches := imgPattern.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		images = append(images, match[2])
	}

	return images
}
