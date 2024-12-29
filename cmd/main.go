package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"gopilot/internal/chat"
	"gopilot/internal/config"
)

func main() {
	// Define flags
	streamFlag := flag.Bool("stream", false, "Stream the response")
	sFlag := flag.Bool("s", false, "Stream the response (shorthand)")
	newFlag := flag.Bool("new", false, "Start new chat history")
	nFlag := flag.Bool("n", false, "Start new chat history (shorthand)")
	oneShotFlag := flag.Bool("one-shot", false, "Disable chat history")
	oFlag := flag.Bool("o", false, "Disable chat history (shorthand)")
	withContextFlag := flag.String("with-context", "", "Additional context files")
	wFlag := flag.String("w", "", "Additional context files (shorthand)")
	configFlag := flag.String("config", "", "Configuration file path")
	cFlag := flag.String("c", "", "Configuration file path (shorthand)")

	flag.Parse()

	// Validate args
	if len(flag.Args()) < 1 {
		fmt.Println("Error: Please provide a prompt")
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.Load(*configFlag)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Create chat session
	session := chat.NewSession(cfg)

	// Handle prompt
	prompt := flag.Arg(0)
	var input interface{}

	// Try to parse as JSON
	if err := json.Unmarshal([]byte(prompt), &input); err != nil {
		// Not JSON, treat as plain text
		input = prompt
	}

	// Handle context files if provided
	contextFiles := *withContextFlag
	if *wFlag != "" {
		contextFiles = *wFlag
	}
	if contextFiles != "" {
		files := strings.Split(contextFiles, ",")
		for _, file := range files {
			content, err := os.ReadFile(file)
			if err != nil {
				fmt.Printf("Error reading context file %s: %v\n", file, err)
				continue
			}
			session.AddContext(string(content))
		}
	}

	// Configure session options
	opts := chat.Options{
		Stream:  *streamFlag || *sFlag,
		NewChat: *newFlag || *nFlag,
		OneShot: *oneShotFlag || *oFlag,
	}

	// Get response
	response, err := session.Send(input, opts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(response)
}
