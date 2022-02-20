package str2color

import (
	"fmt"

	"github.com/michalslomczynski/str2color/hash"
)

const (
	// ANSI RGB pallete <18,230>
	min = 18
	max = 230
	// Random value <0,1>
	c = 0.61
)

// Returns 256 byte colored string with ANSI escape code.
func Format(s string) string {
	color := min + hash.HashString(s, c, max)
	return fmt.Sprintf("\033[38;5;%vm%s", color, s)
}
