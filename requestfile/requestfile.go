package requestfile

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var firstListRegex = regexp.MustCompile(`(GET|POST|PUT|DELETE|PATCH|OPTIONS)\ ?([a-zA-Z0-9/?=&%.]+)\ ?(HTTPS?/\d.?\d?|\d)`)
var headerBlockRegex = regexp.MustCompile(`([a-zA-Z-\d,\ ]+):\ ?([a-zA-Z-,\ ;*:=+()//“”""‘’<>.\d]+)`)
var bodyRegex = regexp.MustCompile(`\n([a-zA-Z\d<>=/&{}""\ ]+)$`)

func FromReadCloser(fallbackDomain string, rc *os.File) (*http.Request, error) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(rc)

	if err != nil {
		return nil, err
	}

	in := buf.String()

	res := firstListRegex.FindAllStringSubmatch(in, -1)

	if len(res) != 1 || len(res[0]) != 4 {
		return nil, ErrUnableToParseRequest("Mangled first line of request")
	}

	method := res[0][1]
	path := res[0][2]
	headers := parseHeaders(in)
	host := resolveHost(headers, fallbackDomain)

	schema := parseSchema(res[0][3])

	r, err := http.NewRequest(method, fmt.Sprintf("%s://%s%s", schema, host, path), parseBody(in))
	r.Header = parseHeaders(in)

	return r, err
}

// parseHeaders received the WHOLE request as a string, and returns http.Headers populated with all headers
func parseHeaders(input string) http.Header {

	res := headerBlockRegex.FindAllStringSubmatch(input, -1)

	headers := make(http.Header, len(res))
	for _, keyVal := range res {
		headers[keyVal[1]] = append(headers[keyVal[1]], keyVal[2])
	}

	return headers
}

func parseSchema(input string) schema {
	if len(input) >= 5 && input[:5] == "HTTPS" {
		return schema{
			isSecure: true,
		}
	}
	return schema{}
}

type schema struct {
	isSecure bool
}

func (s schema) String() string {
	if s.isSecure {
		return "https"
	}
	return "http"
}

func parseBody(input string) io.Reader {

	sub := bodyRegex.FindStringSubmatch(input)

	if len(sub) == 0 || sub[1] == "" {
		return http.NoBody
	}

	return strings.NewReader(sub[1])
}

func resolveHost(h http.Header, fallbackDomain string) string {
	v, ok := h["Host"]
	if !ok {
		return fallbackDomain
	}
	return v[0]
}
