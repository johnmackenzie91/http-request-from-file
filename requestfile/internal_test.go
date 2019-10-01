package requestfile

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseHeaders(t *testing.T) {
	testCases := []struct {
		name           string
		inputFile      string
		expectedOutput http.Header
	}{
		{

			name:      "a lot of headers",
			inputFile: `../test_data/get-lots_of_headers.txt`,
			expectedOutput: http.Header{
				"Access-Control-Allow-Credentials": []string{"true"},
				"Access-Control-Allow-Headers":     []string{"X-PINGOTHER"},
				"Access-Control-Allow-Methods":     []string{"PUT, DELETE, XMODIFY"},
				"Access-Control-Allow-Origin":      []string{"http://example.org"},
				"Access-Control-Expose-Headers":    []string{"X-My-Custom-Header, X-Another-Custom-Header"},
				"Access-Control-Max-Age":           []string{"2520"},
				"Accept-Ranges":                    []string{"bytes"},
				"Age":                              []string{"12"},
				"Allow":                            []string{"GET, HEAD, POST, OPTIONS"},
				"Alternate-Protocol":               []string{"443:npn-spdy/2,443:npn-spdy/2"},
				"Cache-Control":                    []string{"private, no-cache, must-revalidate"},
				"Client-Date":                      []string{"Tue, 27 Jan 2009 18:17:30 GMT"},
				"Client-Peer":                      []string{"123.123.123.123:80"},
				"Client-Response-Num":              []string{"1"},
				"Connection":                       []string{"Keep-Alive"},
				"Content-Disposition":              []string{"attachment; filename=”example.exe”"},
				"Content-Encoding":                 []string{"gzip"},
				"Content-Language":                 []string{"en"},
				"Content-Length":                   []string{"1329"},
				"Content-Location":                 []string{"/index.htm"},
				"Content-MD5":                      []string{"Q2hlY2sgSW50ZWdyaXR5IQ=="},
				"Content-Range":                    []string{"bytes 21010-47021/47022"},
				"Content-Security-Policy, X-Content-Security-Policy, X-WebKit-CSP": []string{"default-src ‘self’"},
				"Content-Type":                      []string{"text/html"},
				"Date":                              []string{"Fri, 22 Jan 2010 04:00:00 GMT"},
				"ETag":                              []string{"“737060cd8c284d8af7ad3082f209582d”"},
				"Expires":                           []string{"Mon, 26 Jul 1997 05:00:00 GMT"},
				"HTTP":                              []string{"/1.1 401 Unauthorized"},
				"Keep-Alive":                        []string{"timeout=3, max=87"},
				"Last-Modified":                     []string{"Tue, 15 Nov 1994 12:45:26 +0000"},
				"Link":                              []string{"<http://www.example.com/>; rel=”cononical”"},
				"Location":                          []string{"http://www.example.com/"},
				"P3P":                               []string{"policyref=”http://www.example.com/w3c/p3p.xml”, CP=”NOI DSP COR ADMa OUR NOR STA”"},
				"Pragma":                            []string{"no-cache"},
				"Proxy-Authenticate":                []string{"Basic"},
				"Proxy-Connection":                  []string{"Keep-Alive"},
				"Refresh":                           []string{"5; url=http://www.example.com/"},
				"Retry-After":                       []string{"120"},
				"Server":                            []string{"Apache"},
				"Set-Cookie":                        []string{"test=1; domain=example.com; path=/; expires=Tue, 01-Oct-2013 19:16:48 GMT"},
				"Status":                            []string{"200 OK"},
				"Strict-Transport-Security":         []string{"max-age=16070400; includeSubDomains"},
				"Timing-Allow-Origin":               []string{"www.example.com"},
				"Trailer":                           []string{"Max-Forwards"},
				"Transfer-Encoding":                 []string{"chunked"},
				"Upgrade":                           []string{"HTTP/2.0, SHTTP/1.3, IRC/6.9, RTA/x11"},
				"Vary":                              []string{"*"},
				"Via":                               []string{"1.0 fred, 1.1 example.com (Apache/1.1)"},
				"WWW-Authenticate":                  []string{"Basic"},
				"X-Aspnet-Version":                  []string{"2.0.50727"},
				"X-Content-Type-Options":            []string{"nosniff"},
				"X-Frame-Options":                   []string{"deny"},
				"X-Permitted-Cross-Domain-Policies": []string{"master-only"},
				"X-Pingback":                        []string{"http://www.example.com/pingback/xmlrpc"},
				"X-Powered-By":                      []string{"PHP/5.4.0"},
				"X-Robots-Tag":                      []string{"noindex,nofollow"},
				"X-UA-Compatible":                   []string{"Chome=1"},
				"X-XSS-Protection":                  []string{"1; mode=block"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			content, err := ioutil.ReadFile(tc.inputFile)
			if err != nil {
				panic(err)
			}
			// act
			out := parseHeaders(string(content))
			// assert
			assert.Equal(t, tc.expectedOutput, out)
		})
	}
}

func Test_parseSchema(t *testing.T) {
	testCases := []struct {
		input       string
		expectedOut schema
	}{
		{
			input:       "HTTP/1.0",
			expectedOut: schema{isSecure: false},
		},
		{
			input:       "HTTPS/1.1",
			expectedOut: schema{isSecure: true},
		},
		{
			input:       "1",
			expectedOut: schema{isSecure: false},
		},
		{
			input:       "HTTP/2",
			expectedOut: schema{isSecure: false},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// arrange
			// act
			out := parseSchema(tc.input)
			// assert
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}
