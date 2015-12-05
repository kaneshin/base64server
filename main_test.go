package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCandidates = []struct {
	raw     string
	encoded string
}{
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},
	{"hello world", "aGVsbG8gd29ybGQ="},
}

func testSetupBase64encodeServer() *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(base64encodeHandler))
}

func testSetupBase64decodeServer() *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(base64decodeHandler))
}

func TestBase64encodeServer(t *testing.T) {
	assert := assert.New(t)

	server := testSetupBase64encodeServer()
	server.Start()

	for _, c := range testCandidates {
		if u, err := url.Parse(server.URL + "?v=" + c.raw); assert.NoError(err) {
			u.RawQuery = u.Query().Encode()
			if req, err := http.NewRequest("GET", u.String(), nil); assert.NoError(err) {
				if resp, err := http.DefaultClient.Do(req); assert.NoError(err) {
					defer resp.Body.Close()
					if b, err := ioutil.ReadAll(resp.Body); assert.NoError(err) {
						assert.Equal(c.encoded, string(b))
					}
				}
			}
		}
	}
}

func TestBase64decodeServer(t *testing.T) {
	assert := assert.New(t)

	server := testSetupBase64decodeServer()
	server.Start()

	for _, c := range testCandidates {
		if u, err := url.Parse(server.URL + "?v=" + c.encoded); assert.NoError(err) {
			u.RawQuery = u.Query().Encode()
			if req, err := http.NewRequest("GET", u.String(), nil); assert.NoError(err) {
				if resp, err := http.DefaultClient.Do(req); assert.NoError(err) {
					defer resp.Body.Close()
					if b, err := ioutil.ReadAll(resp.Body); assert.NoError(err) {
						assert.Equal(c.raw, string(b))
					}
				}
			}
		}
	}
}
