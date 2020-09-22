package badge

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasttemplate"
)

var badgeSVG = trim(`
<svg
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink" width="{{rectWidth}}" height="20" role="img" aria-label="{{title}}: {{value}}">
	<title>{{title}}: {{value}}</title>
	<linearGradient id="s" x2="0" y2="100%">
		<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<clipPath id="r">
		<rect width="100" height="20" rx="3" fill="#fff"/>
	</clipPath>
	<g clip-path="url(#r)">
		<rect width="69" height="20" fill="#555"/>
		<rect x="69" width="31" height="20" fill="#97ca00"/>
		<rect width="100" height="20" fill="url(#s)"/>
	</g>
	<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
		<text aria-hidden="true" x="355" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{titleTextLength}}">{{title}}</text>
		<text x="355" y="140" transform="scale(.1)" fill="#fff" textLength="590">{{title}}</text>
		<text aria-hidden="true" x="835" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{valueTextLength}}">{{value}}</text>
		<text x="835" y="140" transform="scale(.1)" fill="#fff" textLength="{{valueTextLength}}">{{value}}</text>
	</g>
</svg>
`)

const (
	rectWidth       = "100"
	titleTextLength = "590"
)

var (
	valueTextLength = "130"
	tmpl            = fasttemplate.New(badgeSVG, "{{", "}}")
)

// Generate will create a github badge with the provided metadata
func Generate(title, value, color string) (badge []byte) {
	var valueTextLength string

	if valueLen := len(value); valueLen <= 2 {
		valueTextLength = "130"
	} else {
		valueTextLength = strconv.Itoa(len(value) * 70)
	}
	buf := bytebufferpool.Get()
	_, _ = tmpl.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "rectWidth":
			return buf.WriteString(rectWidth)
		case "title":
			return buf.WriteString(title)
		case "value":
			return buf.WriteString(value)
		case "color":
			return buf.WriteString(color)
		case "titleTextLength":
			return buf.WriteString(titleTextLength)
		case "valueTextLength":
			return buf.WriteString(valueTextLength)
		}
		return 0, nil
	})
	badge = buf.Bytes()
	bytebufferpool.Put(buf)
	return
}

// Remove all tabs, spaces and new lines outside the tags
func trim(str string) string {
	trimmed := strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(str, " "))
	trimmed = strings.Replace(trimmed, " <", "<", -1)
	trimmed = strings.Replace(trimmed, "> ", ">", -1)
	return trimmed
}
