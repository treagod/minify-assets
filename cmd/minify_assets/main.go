package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

func minifyFile(m *minify.M, filePath string, mimeType string) error {
	// Read the original file
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Minify the content
	minified, err := m.Bytes(mimeType, input)
	if err != nil {
		return fmt.Errorf("failed to minify file %s: %w", filePath, err)
	}

	// Write the minified content back to the file
	err = ioutil.WriteFile(filePath, minified, 0644)
	if err != nil {
		return fmt.Errorf("failed to write minified file %s: %w", filePath, err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the root directory as an argument")
	}

	rootDir := os.Args[1]

	// Initialize the minifier
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a CSS or JS file
		if !info.IsDir() {
			switch filepath.Ext(path) {
			case ".css":
				err := minifyFile(m, path, "text/css")
				if err != nil {
					log.Printf("Failed to minify CSS file %s: %v", path, err)
				}
			case ".js":
				err := minifyFile(m, path, "application/javascript")
				if err != nil {
					log.Printf("Failed to minify JS file %s: %v", path, err)
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %s: %v", rootDir, err)
	}
}
