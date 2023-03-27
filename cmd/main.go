package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/polidog/hugo-to-wordpress/config"
	"github.com/polidog/hugo-to-wordpress/hugo"
	"github.com/polidog/hugo-to-wordpress/wordpress"
)

func main() {
	// Parse command line arguments
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	// Check if the configuration file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalf("Configuration file not found: %s", configFile)
	}

	// Load config
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Parse Hugo posts
	posts, err := hugo.ParseHugoPosts(cfg.Hugo.ContentPath)
	if err != nil {
		log.Fatalf("Failed to parse Hugo posts: %v", err)
	}

	// Migrate posts to WordPress
	wpClient := wordpress.NewClient(cfg.WordPress.URL, cfg.WordPress.Username, cfg.WordPress.Password, cfg.ConcurrentWorkers)
	if err := wpClient.MigratePosts(posts); err != nil {
		log.Fatalf("Failed to migrate posts: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
