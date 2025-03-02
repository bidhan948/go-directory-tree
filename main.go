package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	htmlOutput = flag.Bool("html", false, "Output the directory tree to an HTML file instead of the console")
	maxDepth   = flag.Int("maxdepth", -1, "Limit the directory traversal to a maximum depth (-1 means no limit)")
)

func main() {
	flag.Parse()

	// Check if a directory path is provided.
	if flag.NArg() < 1 {
		fmt.Println("Usage: rec_traverse [options] <directory>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	rootDir := flag.Arg(0)

	// Choose output destination.
	var output *os.File
	var err error
	if *htmlOutput {
		outputFile := "directory_tree.html"
		output, err = os.Create(outputFile)
		if err != nil {
			log.Fatalf("Error creating HTML file: %v", err)
		}
		defer output.Close()
		// Write HTML header.
		fmt.Fprintln(output, "<!DOCTYPE html>")
		fmt.Fprintln(output, "<html>")
		fmt.Fprintln(output, "<head>")
		fmt.Fprintln(output, "<meta charset=\"utf-8\">")
		fmt.Fprintln(output, "<title>Directory Tree</title>")
		fmt.Fprintln(output, "<style>body { font-family: monospace; }</style>")
		fmt.Fprintln(output, "</head>")
		fmt.Fprintln(output, "<body>")
		fmt.Fprintf(output, "<h1>Directory Tree for %s</h1>\n", rootDir)
		fmt.Fprintln(output, "<pre>")
	} else {
		output = os.Stdout
	}

	// Print the root directory.
	fmt.Fprintln(output, rootDir)

	// Walk the directory recursively.
	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error accessing %s: %v", path, err)
			return err
		}

		// Get relative path from the root directory.
		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		// Skip the root directory itself.
		if relPath == "." {
			return nil
		}

		// Check the maximum depth if specified.
		if *maxDepth >= 0 {
			level := len(strings.Split(relPath, string(os.PathSeparator))) - 1
			if level > *maxDepth {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// Calculate the indentation level.
		level := len(strings.Split(relPath, string(os.PathSeparator))) - 1
		indent := strings.Repeat("    ", level)
		prefix := indent + "└── "

		// Print directory or file details.
		if d.IsDir() {
			fmt.Fprintf(output, "%s[DIR] %s\n", prefix, d.Name())
		} else {
			info, err := d.Info()
			if err != nil {
				log.Printf("Unable to get info for %s: %v", path, err)
				return err
			}
			fmt.Fprintf(output, "%s%s (%d bytes)\n", prefix, d.Name(), info.Size())
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the directory: %v", err)
	}

	// If HTML output, close HTML tags.
	if *htmlOutput {
		fmt.Fprintln(output, "</pre>")
		fmt.Fprintln(output, "</body>")
		fmt.Fprintln(output, "</html>")
		fmt.Println("HTML file generated: directory_tree.html")
	}
}
