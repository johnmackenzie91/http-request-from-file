package requestfile_test

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/johnmackenzie91/http-request-file/requestfile"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	testCases := []struct {
		filePath       string
		expectedOutput *http.Request
		expectedErr    error
	}{
		{
			filePath: "../test_data/delete.txt",
			expectedOutput: func() *http.Request {
				r, err := http.NewRequest("DELETE", "http://www.example.com/route?name=Joe%20Bloggs/", http.NoBody)
				if err != nil {
					panic(err)
				}
				r.Header = http.Header{
					"Host":                             []string{"www.example.com"},
					"Access-Control-Allow-Credentials": []string{"true"},
				}
				return r
			}(),
		},
		{
			filePath: "../test_data/get.txt",
			expectedOutput: func() *http.Request {
				r, err := http.NewRequest("GET", "http://www.example.com/route?name=Joe%20Bloggs/", http.NoBody)
				if err != nil {
					panic(err)
				}
				r.Header = http.Header{
					"Host": []string{"www.example.com"},
				}

				return r
			}(),
		},
		{
			filePath: "../test_data/post.txt",
			expectedOutput: func() *http.Request {
				r, err := http.NewRequest(
					"POST",
					"http://example.com/route.php",
					strings.NewReader("name1=value1&name2=value2"),
				)
				if err != nil {
					panic(err)
				}
				r.Header = http.Header{
					"Host": []string{"example.com"},
				}
				return r
			}(),
		},
		{
			filePath: "../test_data/put.txt",
			expectedOutput: func() *http.Request {
				r, err := http.NewRequest(
					"PUT",
					"https://example.com/route.html",
					strings.NewReader("<p>New File</p>"),
				)
				if err != nil {
					panic(err)
				}
				r.Header = http.Header{
					"Host":           []string{"example.com"},
					"Content-type":   []string{"text/html"},
					"Content-length": []string{"16"},
				}
				return r
			}(),
		},
		{
			filePath: "../test_data/delete-no_host_header.txt",
			expectedOutput: func() *http.Request {
				r, err := http.NewRequest(
					"DELETE",
					"http://example.com/route?name=Joe%20Bloggs/",
					http.NoBody,
				)
				if err != nil {
					panic(err)
				}
				r.Header = http.Header{
					"Access-Control-Allow-Credentials": []string{"true"},
				}
				return r
			}(),
		},
		{
			filePath: "../test_data/get-websocket-request.txt",
			expectedOutput: func() *http.Request {
				r, err := http.NewRequest(
					"GET",
					"http://server.example.com/chat",
					http.NoBody,
				)
				if err != nil {
					panic(err)
				}
				r.Header = http.Header{
					"Host": []string{"server.example.com"},
					"Upgrade": []string{"websocket"},
					"Connection": []string{"Upgrade"},
					"Sec-WebSocket-Key": []string{"x3JJHMbDL1EzLkh9GBhXDw=="},
					"Sec-WebSocket-Protocol": []string{"chat, superchat"},
					"Sec-WebSocket-Version": []string{"13"},
					"Origin": []string{"http://example.com"},
				}
				return r
			}(),
		},
		{
			filePath:    "../test_data/rubbish.txt",
			expectedErr: requestfile.ErrUnableToParseRequest("Mangled first line of request"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.filePath, func(t *testing.T) {
			// arrange
			f, err := os.Open(tc.filePath)
			if err != nil {
				t.Fatal(err)
			}

			// act
			out, err := requestfile.FromReadCloser("example.com", f)

			// assert
			assert.Equal(t, tc.expectedErr, err)

			if err == nil {
				assert.Equal(t, tc.expectedOutput.Method, out.Method, "method not as expected")
				assert.Equal(t, tc.expectedOutput.URL.Scheme, out.URL.Scheme, "scheme not as expected")
				assert.Equal(t, tc.expectedOutput.Host, out.URL.Host, "host not as expected")
				assert.Equal(t, tc.expectedOutput.URL.String(), out.URL.String(), "url not as expected")
				assert.Equal(t, tc.expectedOutput.Body, out.Body, "body not as expected")
				assert.Equal(t, tc.expectedOutput.Header, out.Header, "headers not as expected")
			}
		})
	}
}
