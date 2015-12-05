package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	port = flag.String("port", "8080", "")

	// BuildVersion sets commit hash of git
	BuildVersion string

	// BuildDate sets date of built datetime
	BuildDate string
)

var base64encodeHandler = func(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("v")
	data := []byte(v)
	io.WriteString(w, base64.StdEncoding.EncodeToString(data))
}

var base64decodeHandler = func(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("v")
	data, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, string(data))
}

func main() {
	flag.Parse()
	http.HandleFunc("/encode", base64encodeHandler)
	http.HandleFunc("/decode", base64decodeHandler)
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("Version: %s\n", BuildVersion))
		io.WriteString(w, fmt.Sprintf("   Date: %s\n", BuildDate))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
