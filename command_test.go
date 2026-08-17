package cobra

import (
	"testing"
)

func TestQuotedArgumentWithHyphenSpace(t *testing.T) {
	c := &Command{
		Use: "test",
		Run: func(cmd *Command, args []string) {
			if len(args) != 1 || args[0] != "- http" {
				t.Errorf("Expected argument '- http', got %v", args)
			}
		},
	}

	c.SetArgs([]string{"- http"})
	err := c.Execute()
	if err != nil {
		tt.Fatalf("Unexpected error: %v", err)
	}
}

// ... [rest of command_test.go content] ...