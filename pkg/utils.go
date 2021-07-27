package pkg

import (
	"regexp"
)

const (
	TERMCOLOR_RESETALL               = "\033[0m"
	TERMCOLOR_BOLD                   = "\033[1m"
	TERMCOLOR_DIM                    = "\033[2m"
	TERMCOLOR_ITALIC                 = "\033[3m"
	TERMCOLOR_UNDERLINED             = "\033[4m"
	TERMCOLOR_BLINK                  = "\033[5m"
	TERMCOLOR_REVERSE                = "\033[7m"
	TERMCOLOR_HIDDEN                 = "\033[8m"
	TERMCOLOR_RESETBOLD              = "\033[21m"
	TERMCOLOR_RESETDIM               = "\033[22m"
	TERMCOLOR_RESETITALIC            = "\033[23m"
	TERMCOLOR_RESETUNDERLINED        = "\033[24m"
	TERMCOLOR_RESETBLINK             = "\033[25m"
	TERMCOLOR_RESETREVERSE           = "\033[27m"
	TERMCOLOR_RESETHIDDEN            = "\033[28m"
	TERMCOLOR_DEFAULT                = "\033[39m"
	TERMCOLOR_BLACK                  = "\033[30m"
	TERMCOLOR_RED                    = "\033[31m"
	TERMCOLOR_GREEN                  = "\033[32m"
	TERMCOLOR_YELLOW                 = "\033[33m"
	TERMCOLOR_BLUE                   = "\033[34m"
	TERMCOLOR_MAGENTA                = "\033[35m"
	TERMCOLOR_CYAN                   = "\033[36m"
	TERMCOLOR_LIGHTGRAY              = "\033[37m"
	TERMCOLOR_DARKGRAY               = "\033[90m"
	TERMCOLOR_LIGHTRED               = "\033[91m"
	TERMCOLOR_LIGHTGREEN             = "\033[92m"
	TERMCOLOR_LIGHTYELLOW            = "\033[93m"
	TERMCOLOR_LIGHBLUE               = "\033[94m"
	TERMCOLOR_LIGHTMAGENTA           = "\033[95m"
	TERMCOLOR_LIGHTCYAN              = "\033[96m"
	TERMCOLOR_WHITE                  = "\033[97m"
	TERMCOLOR_BACKGROUNDDEFAULT      = "\033[49m"
	TERMCOLOR_BACKGROUNDBLACK        = "\033[40m"
	TERMCOLOR_BACKGROUNDRED          = "\033[41m"
	TERMCOLOR_BACKGROUNDGREEN        = "\033[42m"
	TERMCOLOR_BACKGROUNDYELLOW       = "\033[43m"
	TERMCOLOR_BACKGROUNDNLUE         = "\033[44m"
	TERMCOLOR_BACKGROUNDMAGENTA      = "\033[45m"
	TERMCOLOR_BACKGROUNDCYAN         = "\033[46m"
	TERMCOLOR_BACKGROUNDLIGHTGRAY    = "\033[47m"
	TERMCOLOR_BACKGROUNDDARKGRAY     = "\033[100m"
	TERMCOLOR_BACKGROUNDLIGHTRED     = "\033[101m"
	TERMCOLOR_BACKGROUNDLIGHTGREEN   = "\033[102m"
	TERMCOLOR_BACKGROUNDLIGHTYELLOW  = "\033[103m"
	TERMCOLOR_BACKGROUNDLIGHTBLUE    = "\033[104m"
	TERMCOLOR_BACKGROUNDLIGHTMAGENTA = "\033[105m"
	TERMCOLOR_BACKGROUNDLIGHTCYAN    = "\033[106m"
	TERMCOLOR_BACKGROUNDWHITE        = "\033[107m"
)

func BoldText(s string) string {
	return TERMCOLOR_BOLD + s + TERMCOLOR_RESETBOLD
}

func ItalicText(s string) string {
	return TERMCOLOR_ITALIC + s + TERMCOLOR_RESETITALIC
}

func UnderlineText(s string) string {
	return TERMCOLOR_UNDERLINED + s + TERMCOLOR_RESETUNDERLINED
}

func Html2BashText(s string) string {
	// remove <p> </p>
	s = regexp.MustCompile(`<p>|</p>|<pre>|</pre>`).ReplaceAllString(s, "")

	// bold text
	s = regexp.MustCompile(`<strong>`).ReplaceAllString(s, TERMCOLOR_BOLD)
	s = regexp.MustCompile(`</strong>`).ReplaceAllString(s, TERMCOLOR_RESETALL)

	// underline text
	s = regexp.MustCompile(`<u>`).ReplaceAllString(s, TERMCOLOR_UNDERLINED)
	s = regexp.MustCompile(`</u>`).ReplaceAllString(s, TERMCOLOR_RESETUNDERLINED)

	// <em>
	s = regexp.MustCompile(`<em>`).ReplaceAllString(s, TERMCOLOR_ITALIC)
	s = regexp.MustCompile(`</em>`).ReplaceAllString(s, TERMCOLOR_RESETITALIC)

	// make <code> italic
	s = regexp.MustCompile(`<code>`).ReplaceAllString(s, TERMCOLOR_ITALIC)
	s = regexp.MustCompile(`</code>`).ReplaceAllString(s, TERMCOLOR_RESETITALIC)

	// Word hightlight
	s = regexp.MustCompile(`(Example.*:)`).ReplaceAllString(s, TERMCOLOR_RED+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(Constraints.*:)`).ReplaceAllString(s, TERMCOLOR_RED+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(Input)`).ReplaceAllString(s, TERMCOLOR_GREEN+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(Output)`).ReplaceAllString(s, TERMCOLOR_GREEN+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(Explanation)`).ReplaceAllString(s, TERMCOLOR_GREEN+`$1`+TERMCOLOR_RESETALL)

	s = regexp.MustCompile(`(示例.*[:：])`).ReplaceAllString(s, TERMCOLOR_RED+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(输入\s*[:：])`).ReplaceAllString(s, TERMCOLOR_GREEN+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(输出\s*[:：])`).ReplaceAllString(s, TERMCOLOR_GREEN+`$1`+TERMCOLOR_RESETALL)
	s = regexp.MustCompile(`(提示\s*[:：])`).ReplaceAllString(s, TERMCOLOR_GREEN+`$1`+TERMCOLOR_RESETALL)

	// make <code> italic
	s = regexp.MustCompile(`<sup>`).ReplaceAllString(s, "^{")
	s = regexp.MustCompile(`<sub>`).ReplaceAllString(s, "_{")
	s = regexp.MustCompile(`</sup>|</sub>`).ReplaceAllString(s, "}")

	// HTML escape symbols, just some of them.
	s = regexp.MustCompile(`&nbsp;`).ReplaceAllString(s, " ")
	s = regexp.MustCompile(`&lt;`).ReplaceAllString(s, "<")
	s = regexp.MustCompile(`&gt;`).ReplaceAllString(s, ">")
	s = regexp.MustCompile(`&quot;`).ReplaceAllString(s, "\"")
	s = regexp.MustCompile(`&amp;`).ReplaceAllString(s, "&")
	s = regexp.MustCompile(`&apos;|&#39;`).ReplaceAllString(s, "'")

	// list
	s = regexp.MustCompile(`<ul>|</ul>|<ol>|</ol>`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`[\t]*<li>`).ReplaceAllString(s, "    · ")
	s = regexp.MustCompile(`[\t]*</li>`).ReplaceAllString(s, "")

	// Too many empty line
	s = regexp.MustCompile(`[\n]{2,}`).ReplaceAllString(s, "\n\n")
	s = regexp.MustCompile(`[\r]{2,}`).ReplaceAllString(s, "\r\r")

	return s
}
