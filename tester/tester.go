package tester

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// Media types
const (
	Version         = "1.0"
	UserAgent       = "User-Agent"
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	ApplicationForm = "application/x-www-form-urlencoded"
)

// HTTPResponse is basic HTTP response type
type HTTPResponse *httptest.ResponseRecorder

// HTTPRequest is basic HTTP request type
type HTTPRequest *http.Request

// ResponseFunc response handling func type
type ResponseFunc func(HTTPResponse, HTTPRequest)

// H is HTTP Header Type
type H map[string]string

// D is HTTP Data Type
type D map[string]interface{}

// RequestConfig provide user input request structure
type RequestConfig struct {
	Method  string
	Path    string
	Body    string
	Headers H
	Cookies H
	Debug   bool
}

// New supply initial structure
func New() *RequestConfig {

	return &RequestConfig{}
}

// SetDebug supply enable debug mode.
func (rc *RequestConfig) SetDebug(enable bool) *RequestConfig {
	rc.Debug = enable

	return rc
}

// GET is request method.
func (rc *RequestConfig) GET(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "GET"

	return rc
}

// POST is request method.
func (rc *RequestConfig) POST(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "POST"

	return rc
}

// PUT is request method.
func (rc *RequestConfig) PUT(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "PUT"

	return rc
}

// DELETE is request method.
func (rc *RequestConfig) DELETE(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "DELETE"

	return rc
}

// PATCH is request method.
func (rc *RequestConfig) PATCH(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "PATCH"

	return rc
}

// HEAD is request method.
func (rc *RequestConfig) HEAD(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "HEAD"

	return rc
}

// OPTIONS is request method.
func (rc *RequestConfig) OPTIONS(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "OPTIONS"

	return rc
}

// SetHeader supply http header what you defined.
func (rc *RequestConfig) SetHeader(headers H) *RequestConfig {
	if len(headers) > 0 {
		rc.Headers = headers
	}

	return rc
}

// SetJSON supply JSON body.
func (rc *RequestConfig) SetJSON(body D) *RequestConfig {
	if b, err := json.Marshal(body); err == nil {
		rc.Body = string(b)
	}

	return rc
}

// SetForm supply form body.
func (rc *RequestConfig) SetForm(body H) *RequestConfig {
	f := make(url.Values)

	for k, v := range body {
		f.Set(k, v)
	}

	rc.Body = f.Encode()

	return rc
}

// SetQuery supply query string.
func (rc *RequestConfig) SetQuery(query H) *RequestConfig {
	f := make(url.Values)

	for k, v := range query {
		f.Set(k, v)
	}

	if strings.Contains(rc.Path, "?") {
		rc.Path = rc.Path + "&" + f.Encode()
	} else {
		rc.Path = rc.Path + "?" + f.Encode()
	}

	return rc
}

// SetBody supply raw body.
func (rc *RequestConfig) SetBody(body string) *RequestConfig {
	if len(body) > 0 {
		rc.Body = body
	}

	return rc
}

// SetCookie supply cookies what you defined.
func (rc *RequestConfig) SetCookie(cookies H) *RequestConfig {
	if len(cookies) > 0 {
		rc.Cookies = cookies
	}

	return rc
}

func (rc *RequestConfig) initTest() (*http.Request, *httptest.ResponseRecorder) {
	qs := ""
	if strings.Contains(rc.Path, "?") {
		ss := strings.Split(rc.Path, "?")
		rc.Path = ss[0]
		qs = ss[1]
	}

	body := bytes.NewBufferString(rc.Body)

	req, _ := http.NewRequest(rc.Method, rc.Path, body)

	if len(qs) > 0 {
		req.URL.RawQuery = qs
	}

	// Auto add user agent
	req.Header.Set(UserAgent, "cuxs-tester/"+Version)

	if rc.Method == "POST" || rc.Method == "PUT" {
		if strings.HasPrefix(rc.Body, "{") {
			req.Header.Set(ContentType, ApplicationJSON)
		} else {
			req.Header.Set(ContentType, ApplicationForm)
		}
	}

	if len(rc.Headers) > 0 {
		for k, v := range rc.Headers {
			req.Header.Set(k, v)
		}
	}

	if len(rc.Cookies) > 0 {
		for k, v := range rc.Cookies {
			req.AddCookie(&http.Cookie{Name: k, Value: v})
		}
	}

	if rc.Debug {
		log.Printf("Request Method: %s", rc.Method)
		log.Printf("Request Path: %s", rc.Path)
		log.Printf("Request Body: %s", rc.Body)
		log.Printf("Request Headers: %s", rc.Headers)
		log.Printf("Request Cookies: %s", rc.Cookies)
		log.Printf("Request Header: %s", req.Header)
	}

	w := httptest.NewRecorder()

	return req, w
}

// Run execute http request
func (rc *RequestConfig) Run(r http.Handler, response ResponseFunc) {

	req, w := rc.initTest()
	r.ServeHTTP(w, req)

	response(w, req)
}
