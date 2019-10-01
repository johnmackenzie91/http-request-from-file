# HTTP Request File

[![Build Status](https://travis-ci.org/johnmackenzie91/http-request-from-file.svg?branch=master)](https://travis-ci.org/johnmackenzie91/http-request-from-file)
[![Coverage Status](https://coveralls.io/repos/github/johnmackenzie91/http-request-from-file/badge.svg?branch=master)](https://coveralls.io/github/johnmackenzie91/http-request-from-file?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmackenzie91/http-request-from-file)](https://goreportcard.com/report/github.com/johnmackenzie91/http-request-from-file)
[![](https://godoc.org/github.com/johnmackenzie91/http-request-from-file/paginator?status.svg)](https://godoc.org/github.com/johnmackenzie91/http-request-from-file/paginator)

A package that parses http requests and outputs *http.Request.
This could come in useful for unit testing http.HandlerFunc.

## Installing

Install http-request-file via the "go get" command:

```cgo
go get github.com/johnmackenzie91/http-request-from-file
```

## Getting Started

```cgo

func main() {
	// open a file stream
	f, _ := os.Open(tc.filePath)	
	// translate it into a *http.Request
	req, err := requestfile.FromReadCloser("example.com", f)
	
	fmt.Println(req.URL)
	fmt.Println(req.Method)
}
```

## Social Media

* [Twitter](https://twitter.com/JohnnMackk)
* [LinkedIn](https://www.linkedin.com/in/john-mackenzie-web-developer/)