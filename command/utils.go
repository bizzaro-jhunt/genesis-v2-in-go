package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/starkandwayne/goutils/ansi"
)

var debug bool

func BoolString(tf bool) string {
	if tf {
		return "Y"
	}
	return "N"
}

func PrettyJSON(raw string) string {
	tmpBuf := bytes.Buffer{}
	err := json.Indent(&tmpBuf, []byte(raw), "", "  ")
	if err != nil {
		DEBUG("json.Indent failed with %s", err)
		return raw
	}
	return tmpBuf.String()
}

func DEBUG(format string, args ...interface{}) {
	if debug {
		content := ansi.Sprintf(format, args...)
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			lines[i] = "DEBUG> " + line
		}
		content = strings.Join(lines, "\n")
		fmt.Fprintf(os.Stderr, "%s\n", content)
	}
}
