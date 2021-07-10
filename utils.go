package pkg

import "regexp"

const (
	ESCAPE_CODE_UNDERLINE_START = `\033[4m`
	ESCAPE_CODE_END             = `\033[0m`
)

func Html2BashText(s string) string {
	// remove <p> </p>
	s = regexp.MustCompile(`<p>|</p>`).ReplaceAllString(s, "")

	// handle <strong> and </strong>
	//s = regexp.MustCompile(`<strong>`).ReplaceAllString(s, ESCAPE_CODE_UNDERLINE_START)
	//s = regexp.MustCompile(`</strong>`).ReplaceAllString(s, ESCAPE_CODE_END)

	s = regexp.MustCompile(`<u>`).ReplaceAllString(s, ESCAPE_CODE_UNDERLINE_START)
	s = regexp.MustCompile(`</u>`).ReplaceAllString(s, ESCAPE_CODE_END)
	return s
}
