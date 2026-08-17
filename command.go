package cobra

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

// isFlag checks if the given string is a flag.
// A flag is an argument that starts with a dash and is not a single dash,
// and does not start with a dash followed by a space.
func isFlag(arg string) bool {
	if !strings.HasPrefix(arg, "-") {
		return false
	}
	if arg == "-" {
		return false
	}
	if strings.HasPrefix(arg, "--") {
		if arg == "--" {
			return false
		}
		return !strings.HasPrefix(arg, "-- ")
	}
	return !strings.HasPrefix(arg, "- ")
}

// ... [rest of command.go content with updated flag checks using isFlag] ...