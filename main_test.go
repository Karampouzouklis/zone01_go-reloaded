package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	// Create temporary directory for test files
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		args        []string
		inputText   string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid arguments",
			args:        []string{"input.txt", "output.txt"},
			inputText:   "Hello, world!",
			expectError: false,
		},
		{
			name:        "missing arguments",
			args:        []string{},
			expectError: true,
			errorMsg:    "usage:",
		},
		{
			name:        "too many arguments",
			args:        []string{"input.txt", "output.txt", "extra.txt"},
			expectError: true,
			errorMsg:    "usage:",
		},
		{
			name:        "input file not found",
			args:        []string{"nonexistent.txt", "output.txt"},
			expectError: true,
			errorMsg:    "failed to read input file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test files in temp directory
			if tt.inputText != "" {
				inputPath := filepath.Join(tempDir, tt.args[0])
				err := os.WriteFile(inputPath, []byte(tt.inputText), 0644)
				if err != nil {
					t.Fatalf("Failed to create test input file: %v", err)
				}
				// Update args to use temp directory paths
				tt.args[0] = inputPath
				tt.args[1] = filepath.Join(tempDir, tt.args[1])
			}

			// Run the function
			err := run(tt.args)

			// Check error expectations
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				// Verify output file was created and has correct content
				if len(tt.args) == 2 {
					outputContent, err := os.ReadFile(tt.args[1])
					if err != nil {
						t.Errorf("Failed to read output file: %v", err)
					} else if string(outputContent) != tt.inputText {
						t.Errorf("Expected output '%s', got '%s'", tt.inputText, string(outputContent))
					}
				}
			}
		})
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    s[:len(substr)] == substr || 
		    s[len(s)-len(substr):] == substr ||
		    findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}