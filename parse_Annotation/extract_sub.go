package parse

import (
	"fmt"
	"strings"
)

func ExtractSubs(value string, sep string) (string, string, error) {
	// parse the value to see if it's a valid subvalue format
	sp := strings.SplitN(value, sep, 2)
	if len(sp) == 1 {
		return "", "", fmt.Errorf("invalid subvalue format for %s (no %s found)", value, sep)
	}

	subkey := strings.TrimSpace(sp[0])
	subvalue := strings.TrimSpace(sp[1])

	return subkey, subvalue, nil
}
